// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck" "test" {
					writeonly_attr = "hello world!"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_datacheck.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_datacheck.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
				},
			},
			{
				ResourceName:      "corner_writeonly_datacheck.test",
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck_planerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				// TODO: This error message will likely be changed to be more specific before 1.11 GA
				ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck_applyerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				// TODO: This error message will likely be changed to be more specific before 1.11 GA
				ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck_readerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				// TODO: This error message will likely be changed to be more specific before 1.11 GA
				ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck_importerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly_datacheck_importerror.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
						plancheck.ExpectResourceAction("corner_writeonly_datacheck_importerror.test", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_datacheck_importerror.test", tfjsonpath.New("writeonly_attr"), knownvalue.Null()),
				},
			},
			{
				ResourceName:      "corner_writeonly_datacheck_importerror.test",
				ImportState:       true,
				ImportStateVerify: true,
				// TODO: This error message will likely be changed to be more specific before 1.11 GA
				ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_datacheck_upgraderesourceerror" "test" {
					writeonly_attr = "hello world!"
				}`,
				// TODO: This test is currently bugged because UpgradeResourceState TF core is not returning errors for non-null W/O values.
				// This should be uncommented when fixed: https://hashicorp.slack.com/archives/C071HC4JJCC/p1736289662267609
				// ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
				ExpectError: regexp.MustCompile(`After applying this test step, the non-refresh plan was not empty.`),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "terraform_data" "test" {
					input = "hello world!"
				}`,
			},
			{
				Config: `resource "corner_writeonly_datacheck_moveresourceerror" "test" {
					writeonly_attr = "hello world!"
				}
					
				moved {
					from = terraform_data.test
					to   = corner_writeonly_datacheck_moveresourceerror.test
				}`,
				// TODO: This error message will likely be changed to be more specific before 1.11 GA
				ExpectError: regexp.MustCompile(`Error: Write-only attribute set`),
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
