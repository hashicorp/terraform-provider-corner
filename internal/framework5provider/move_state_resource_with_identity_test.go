// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

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
					to   = framework_move_state_with_identity.new
				}
				resource "framework_move_state_with_identity" "new" {}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// The previous framework_identity.old identity should be moved to this new location, split into the new location identity and state.
					statecheck.ExpectIdentity("framework_move_state_with_identity.new", map[string]knownvalue.Check{
						"id": knownvalue.StringExact("id-123"),
					}),
					statecheck.ExpectKnownValue("framework_move_state_with_identity.new", tfjsonpath.New("moved_random_string"), knownvalue.StringExact("tom")),
				},
			},
		},
	})
}
