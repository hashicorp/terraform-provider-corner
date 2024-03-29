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
	"github.com/hashicorp/terraform-provider-corner/internal/cornertesting"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectWithDynamicFunction_known_primitive(t *testing.T) {
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
					value = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = "value1",
    					"dynamic_attr2" = true,
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"dynamic_attr1": cty.String,
								"dynamic_attr2": cty.Bool,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"dynamic_attr1": knownvalue.StringExact("value1"),
							"dynamic_attr2": knownvalue.Bool(true),
						},
					)),
				},
			},
		},
	})
}

func TestObjectWithDynamicFunction_known_collection(t *testing.T) {
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
					value = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = tolist([1, 2, 3, 4]),
    					"dynamic_attr2" = tomap({
							"key1": "hello",
							"key2": true,
						}),
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"dynamic_attr1": cty.List(cty.Number),
								"dynamic_attr2": cty.Map(cty.String),
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"dynamic_attr1": knownvalue.ListExact(
								[]knownvalue.Check{
									knownvalue.Int64Exact(1),
									knownvalue.Int64Exact(2),
									knownvalue.Int64Exact(3),
									knownvalue.Int64Exact(4),
								},
							),
							"dynamic_attr2": knownvalue.MapExact(
								map[string]knownvalue.Check{
									"key1": knownvalue.StringExact("hello"),
									"key2": knownvalue.StringExact("true"),
								},
							),
						},
					)),
				},
			},
		},
	})
}

func TestObjectWithDynamicFunction_known_structural(t *testing.T) {
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
					value = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = [1, 2, 3, 4],
    					"dynamic_attr2" = {
							"attr1": "hello",
							"attr2": true,
						},
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"dynamic_attr1": cty.Tuple([]cty.Type{cty.Number, cty.Number, cty.Number, cty.Number}),
								"dynamic_attr2": cty.Object(
									map[string]cty.Type{
										"attr1": cty.String,
										"attr2": cty.Bool,
									},
								),
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"dynamic_attr1": knownvalue.TupleExact(
								[]knownvalue.Check{
									knownvalue.Int64Exact(1),
									knownvalue.Int64Exact(2),
									knownvalue.Int64Exact(3),
									knownvalue.Int64Exact(4),
								},
							),
							"dynamic_attr2": knownvalue.ObjectExact(
								map[string]knownvalue.Check{
									"attr1": knownvalue.StringExact("hello"),
									"attr2": knownvalue.Bool(true),
								},
							),
						},
					)),
				},
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/955
func TestObjectWithDynamicFunction_Known_AttributeRequired_Error(t *testing.T) {
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
					value = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = "value1"
					})
				}`,
				// This error should always remain with the existing definition
				// as provider developers may be reliant and desire this
				// Terraform behavior. If new framework functionality is added
				// to support optional object attributes, it should be tested
				// separately.
				ExpectError: regexp.MustCompile(`attribute "dynamic_attr2" is\srequired`),
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/955
func TestObjectWithDynamicFunction_Known_AttributeRequired_Null(t *testing.T) {
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
					value = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = "value1"
    					"dynamic_attr2" = null
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"dynamic_attr1": cty.String,
								// No type has been determined yet
								"dynamic_attr2": cty.DynamicPseudoType,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"dynamic_attr1": knownvalue.StringExact("value1"),
							"dynamic_attr2": knownvalue.Null(),
						},
					)),
				},
			},
		},
	})
}

func TestObjectWithDynamicFunction_null(t *testing.T) {
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
					value = provider::framework::object_with_dynamic(null)
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestObjectWithDynamicFunction_unknown(t *testing.T) {
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
					input = provider::framework::object_with_dynamic({
    					"dynamic_attr1" = "value1",
    					"dynamic_attr2" = 123,
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
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"dynamic_attr1": cty.String,
								"dynamic_attr2": cty.Number,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"dynamic_attr1": knownvalue.StringExact("value1"),
							"dynamic_attr2": knownvalue.Int64Exact(123),
						},
					)),
				},
			},
		},
	})
}
