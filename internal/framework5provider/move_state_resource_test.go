// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// This is a smoke test for using the "moved" block to transition state between
// the "random_string" managed resource and the corner provider "framework_move_state"
// managed resource.
//
// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/1039
func TestMoveStateResource(t *testing.T) {
	randomStringSame := statecheck.CompareValue(compare.ValuesSame())

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "hashicorp/random",
			},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "old" {
					length = 12
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					randomStringSame.AddStateValue("random_string.old", tfjsonpath.New("result")),
				},
			},
			{
				Config: `
				moved {
					from = random_string.old
					to   = framework_move_state.new
				}
				resource "framework_move_state" "new" {}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// The previous random_string.result value should be moved to this new location unchanged.
					randomStringSame.AddStateValue("framework_move_state.new", tfjsonpath.New("moved_random_string")),
				},
			},
		},
	})
}

func TestMoveStateResource_identity(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "old" {
					name = "tom"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.old", map[string]knownvalue.Check{
						"id":   knownvalue.StringExact("id-123"),
						"name": knownvalue.StringExact("tom"),
					}),
				},
			},
			{
				Config: `
				moved {
					from = framework_identity.old
					to   = framework_move_state.new
				}
				resource "framework_move_state" "new" {}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// The previous framework_identity.old identity should be moved to this new location, split into the new location identity and state.
					statecheck.ExpectIdentity("framework_move_state.new", map[string]knownvalue.Check{
						"id": knownvalue.StringExact("id-123"),
					}),
					statecheck.ExpectKnownValue("framework_move_state.new", tfjsonpath.New("moved_random_string"), knownvalue.StringExact("tom")),
				},
			},
		},
	})
}
