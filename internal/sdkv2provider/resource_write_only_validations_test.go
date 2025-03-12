// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create function.
//
// This test flips between using a normal configured attribute (old_password_attr) and a write-only attribute
// with a configured version attribute (password_version, writeonly_password).
func TestWriteOnlyValidationsResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// TODO: Swap version check with below line once terraform-plugin-testing@v1.12.0 is released
			// tfversion.SkipBelow(tfversion.Version1_11_0),
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      `resource "corner_writeonly_validations" "test" {}`,
				ExpectError: regexp.MustCompile(`Invalid combination of arguments`),
			},
			{
				// TODO: The testing framework can't verify warning diagnostics currently, although one would be returned here
				// to indicate that the preferred attribute is "writeonly_password". This should only appear when the client supports write-only attributes.
				// https://github.com/hashicorp/terraform-plugin-testing/issues/69
				Config: `resource "corner_writeonly_validations" "test" {
					old_password_attr = "oldpassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_validations.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.StringExact("oldpassword")),
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "corner_writeonly_validations" "test" {
					password_version = "v1"
					old_password_attr = "oldpassword"
				}`,
				ExpectError: regexp.MustCompile(`"old_password_attr": conflicts with password_version`),
			},
			{
				Config: `resource "corner_writeonly_validations" "test" {
					writeonly_password = "newpassword"
				}`,
				ExpectError: regexp.MustCompile(`Missing required argument`), // password_version is needed to handle triggering the replacements
			},
			{
				// Replaces resource with new write-only attribute + version
				Config: `resource "corner_writeonly_validations" "test" {
					password_version = "v1"
					writeonly_password = "newpassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				// No-op, as password_version did not change
				Config: `resource "corner_writeonly_validations" "test" {
					password_version = "v1"
					writeonly_password = "won't trigger an update on it's own"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_validations.test", plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "corner_writeonly_validations" "test" {
					password_version = "v2"
				}`,
				ExpectError: regexp.MustCompile(`Missing required argument`), // writeonly_password is needed to set new password
			},
			{
				// Triggers replace with new password_version
				Config: `resource "corner_writeonly_validations" "test" {
					password_version = "v2"
					writeonly_password = "newpassword2"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
			{
				// Switching back to normal configured attribute which is stored in state
				Config: `resource "corner_writeonly_validations" "test" {
					old_password_attr = "oldpassword2"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_validations.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("old_password_attr"), knownvalue.StringExact("oldpassword2")),
					statecheck.ExpectKnownValue("corner_writeonly_validations.test", tfjsonpath.New("writeonly_password"), knownvalue.Null()),
				},
			},
		},
	})
}
