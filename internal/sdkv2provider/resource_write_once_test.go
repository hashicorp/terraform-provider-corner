// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create function.
func TestWriteOnceResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "1"
					writeonce_string = "fakepassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
				},
			},
			{
				// Now that the resource is created, we can remove the attribute with no planned changes
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "1"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionNoop),
					},
				},
			},
			{
				// Write-only attributes cannot participate and plan as they will always be null in prior/proposed new state
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "1"
					writeonce_string = "this value cannot prompt a change on it's own"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionNoop),
					},
				},
			},
			{
				// trigger_attr will prompt the replace action here
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "2"
					writeonce_string = "fakepassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionReplace),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestWriteOnceResource_error_on_create(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// This error message should occur on all Terraform versions.
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "1"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionCreate),
					},
				},
				ExpectError: regexp.MustCompile(`"writeonce_string" is required when creating the resource, but no definition was found.`),
			},
		},
	})
}

func TestWriteOnceResource_error_on_replace(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "1"
					writeonce_string = "fakepassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "corner_writeonce" "test" {
					trigger_attr = "2"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonce.test", tfjsonpath.New("writeonce_string"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonce.test", plancheck.ResourceActionReplace),
					},
				},
				ExpectError: regexp.MustCompile(`"writeonce_string" is required when creating the resource, but no definition was found.`),
			},
		},
	})
}
