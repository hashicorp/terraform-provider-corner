// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccResourceDeferredAction(t *testing.T) resource.TestCase {
	return resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Apply: resource.ApplyOptions{AllowDeferral: true},
			Plan:  resource.PlanOptions{AllowDeferral: true},
		},
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test that the resource CustomizeDiff logic is skipped during deferral
			// when Plan Modification behavior is disabled.
			{
				Config: `provider "corner" {
					deferral = true
				}

				resource "corner_deferred_action" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 200 # invalid age value
				}`,
				// Expect a passing test with an invalid age value
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("corner_deferred_action.foo", plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
			},
			{
				Config: `resource "corner_deferred_action" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 50
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
			},
		},
	}
}

func testAccResourceDeferredActionPlanModification(t *testing.T) resource.TestCase {
	return resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Apply: resource.ApplyOptions{AllowDeferral: true},
			Plan:  resource.PlanOptions{AllowDeferral: true},
		},
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test that the resource CustomizeDiff logic correctly runs during deferral
			// when Plan Modification behavior is enabled.
			{
				Config: `provider "corner" {
					deferral = true
				}

				resource "corner_deferred_action_plan_modification" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 200 # invalid age value
				}`,
				// Expect a validation error with an invalid age value
				ExpectError: regexp.MustCompile("Error: age value must be less than 100"),
			},
			{
				Config: `resource "corner_deferred_action_plan_modification" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 50
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
			},
		},
	}
}
