// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccResourceDeferredAction(t *testing.T) resource.TestCase {
	return resource.TestCase{
		// Deferred action support is only available in 1.9 alpha binaries (or dev builds)
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		// Add the -allow-deferral flag
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Apply: resource.ApplyOptions{AllowDeferral: true},
			Plan:  resource.PlanOptions{AllowDeferral: true},
		},
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configWithDeferredResource,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("corner_deferred_action.foo", plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
				ExpectNonEmptyPlan: true,
			},
			{
				Config: configWithDeferredResource,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_deferred_action.foo", tfjsonpath.New("name"), knownvalue.StringExact("Ford Prefect")),
				},
			},
		},
	}
}

const configWithDeferredResource = `
provider "corner" {
	api_key = random_string.name.result
}

resource "random_string" "name" {
	length = 5
}

resource "corner_deferred_action" "foo" {
 email = "ford@prefect.co"
 name = "Ford Prefect"
 age = 200
}
`

const configWithDeferredResourcePlanModification = `
provider "corner" {
	api_key = random_string.name.result
}

resource "random_string" "name" {
	length = 5
}

resource "corner_deferred_action_plan_modification" "foo" {
 email = "ford@prefect.co"
 name = "Ford Prefect"
 age = 200
}
`
