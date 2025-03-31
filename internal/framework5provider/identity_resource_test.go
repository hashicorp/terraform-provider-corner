// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestIdentityResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Latest alpha version that this JSON data is available in
		// https://github.com/hashicorp/terraform/releases/tag/v1.12.0-alpha20250319
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.12.0-alpha20250319"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "test" {
					name = "john"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.test", map[string]knownvalue.Check{
						"id":   knownvalue.StringExact("id-123"),
						"name": knownvalue.StringExact("my name is john"),
					}),
				},
			},
			{
				Config: `resource "framework_identity" "test" {
					name = "jerry"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("framework_identity.test", plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.test", map[string]knownvalue.Check{
						"id":   knownvalue.StringExact("id-123"),
						"name": knownvalue.StringExact("my name is john"), // doesn't get updated, since identity should not change.
					}),
				},
			},
		},
	})
}
