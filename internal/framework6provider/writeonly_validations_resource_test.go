// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create function.
func TestWriteOnlyValidationsResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config:      `resource "framework_writeonly_validations" "test" {}`,
				ExpectError: regexp.MustCompile(`Missing Attribute Configuration`),
			},
			{
				// TODO: The testing framework can't verify warning diagnostics currently, although one would be returned here
				// to indicate that the preferred attribute is "writeonly_password". This should only appear when the client supports write-only attributes.
				// https://github.com/hashicorp/terraform-plugin-testing/issues/69
				Config: `resource "framework_writeonly_validations" "test" {
					old_password_attr = "oldpassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("framework_writeonly_validations.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.StringExact("oldpassword")),
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_writeonly_validations" "test" {
					password_version = "v1"
					old_password_attr = "oldpassword"
				}`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			{
				Config: `resource "framework_writeonly_validations" "test" {
					writeonly_password = "newpassword"
				}`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`), // password_version is needed to handle triggering the replacements
			},
			{
				// Replaces resource with new write-only attribute + version
				Config: `resource "framework_writeonly_validations" "test" {
					password_version = "v1"
					writeonly_password = "newpassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("framework_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				// No-op, as password_version did not change
				Config: `resource "framework_writeonly_validations" "test" {
					password_version = "v1"
					writeonly_password = "won't trigger an update on it's own"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("framework_writeonly_validations.test", plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_writeonly_validations" "test" {
					password_version = "v2"
				}`,
				ExpectError: regexp.MustCompile(`Missing Attribute Configuration`), // writeonly_password is needed to set new password
			},
			{
				// Triggers replace with new password_version
				Config: `resource "framework_writeonly_validations" "test" {
					password_version = "v2"
					writeonly_password = "newpassword2"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("framework_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				// Switching back to normal configured attribute which is stored in state
				Config: `resource "framework_writeonly_validations" "test" {
					old_password_attr = "oldpassword2"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("framework_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.StringExact("oldpassword2")),
					statecheck.ExpectKnownValue("framework_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
		},
	})
}
