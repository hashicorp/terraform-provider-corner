// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestObjectFunction_known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::object({
    					"attr1" = "value1",
    					"attr2" = 123,
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"attr1": knownvalue.StringExact("value1"),
							"attr2": knownvalue.Int64Exact(123),
						},
					)),
				},
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/955
func TestObjectFunction_Known_AttributeRequired_Error(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::object({
    					"attr1" = "value1"
					})
				}`,
				// This error should always remain with the existing definition
				// as provider developers may be reliant and desire this
				// Terraform behavior. If new framework functionality is added
				// to support optional object attributes, it should be tested
				// separately.
				ExpectError: regexp.MustCompile(`attribute "attr2" is required`),
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/955
func TestObjectFunction_Known_AttributeRequired_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// AllowNullValue being disabled should not affect this
				// configuration being valid. That setting only refers to the
				// object itself.
				Config: `
				output "test" {
					value = provider::framework::object({
    					"attr1" = "value1"
    					"attr2" = null
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"attr1": knownvalue.StringExact("value1"),
							"attr2": knownvalue.Null(),
						},
					)),
				},
			},
		},
	})
}

func TestObjectFunction_null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::object(null)
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestObjectFunction_unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = provider::framework::object({
    					"attr1" = "value1",
    					"attr2" = 123,
					})
				}

				output "test" {
					value = terraform_data.test.output
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownOutputValue("test"),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"attr1": knownvalue.StringExact("value1"),
							"attr2": knownvalue.Int64Exact(123),
						},
					)),
				},
			},
		},
	})
}
