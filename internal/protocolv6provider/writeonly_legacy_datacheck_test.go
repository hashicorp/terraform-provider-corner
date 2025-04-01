// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccResourceWriteOnlyLegacyDataCheck_success(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(false), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_v6_writeonly_legacy_datacheck" "test" {
					writeonly_attr = "hello world!"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_v6_writeonly_legacy_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_v6_writeonly_legacy_datacheck.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_v6_writeonly_legacy_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
				},
			},
			{
				ResourceName:      "corner_v6_writeonly_legacy_datacheck.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceWriteOnlyLegacyDataCheck_plan_error(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(false), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_v6_writeonly_legacy_datacheck_planerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid plan`),
			},
		},
	})
}

func TestAccResourceWriteOnlyLegacyDataCheck_apply_error(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(false), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_v6_writeonly_legacy_datacheck_applyerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid object`),
			},
		},
	})
}
