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

func TestAccResourceWriteOnlyDataCheck_success(t *testing.T) {
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
				Config: `resource "corner_v6_writeonly_datacheck" "test" {
					writeonly_attr = "hello world!"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_v6_writeonly_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_v6_writeonly_datacheck.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_v6_writeonly_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
				},
			},
			{
				ResourceName:      "corner_v6_writeonly_datacheck.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_plan_error(t *testing.T) {
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
				Config: `resource "corner_v6_writeonly_datacheck_planerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid plan`),
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_apply_error(t *testing.T) {
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
				Config: `resource "corner_v6_writeonly_datacheck_applyerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid object`),
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_read_error(t *testing.T) {
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
				Config: `resource "corner_v6_writeonly_datacheck_readerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid object`),
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_import_error(t *testing.T) {
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
				Config: `resource "corner_v6_writeonly_datacheck_importerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_v6_writeonly_datacheck_importerror.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_v6_writeonly_datacheck_importerror.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_v6_writeonly_datacheck_importerror.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
				},
			},
			{
				ResourceName:      "corner_v6_writeonly_datacheck_importerror.test",
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(`Error: Import returned a non-null value for a write-only attribute`),
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_upgraderesource_error(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
					//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
					"corner": func() (tfprotov6.ProviderServer, error) {
						return Server(true), nil
					},
				},
				Config: `resource "corner_v6_writeonly_datacheck_upgraderesourceerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ExpectError: regexp.MustCompile(`Error: Invalid resource state upgrade`),
			},
			{
				ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
					//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
					"corner": func() (tfprotov6.ProviderServer, error) {
						// Back to the original config and turn off upgradeResourceDataError to avoid a destroy clean-up error.
						return Server(false), nil
					},
				},
				Config: `resource "corner_v6_writeonly_datacheck_upgraderesourceerror" "test" {
					writeonly_attr = "hello world!"
				}`,
			},
		},
	})
}

func TestAccResourceWriteOnlyDataCheck_moveresource_error(t *testing.T) {
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
				Config: `resource "terraform_data" "test" {
					input = "hello world!"
				}`,
			},
			{
				Config: `resource "corner_v6_writeonly_datacheck_moveresourceerror" "test" {
					writeonly_attr = "hello world!"
				}
					
				moved {
					from = terraform_data.test
					to   = corner_v6_writeonly_datacheck_moveresourceerror.test
				}`,
				ExpectError: regexp.MustCompile(`Error: Provider returned invalid value`),
			},
			// Back to the original config to avoid a destroy clean-up error.
			{
				Config: `resource "terraform_data" "test" {
					input = "hello world!"
				}`,
			},
		},
	})
}
