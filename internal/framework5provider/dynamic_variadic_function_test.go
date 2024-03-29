// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"math/big"
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

func TestDynamicVariadicFunction_value_zero(t *testing.T) {
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
					value = provider::framework::dynamic_variadic()
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test", cty.Tuple([]cty.Type{})),
					statecheck.ExpectKnownOutputValue("test", knownvalue.TupleExact([]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestDynamicVariadicFunction_value_one(t *testing.T) {
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
					value = provider::framework::dynamic_variadic("one")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test",
						cty.Tuple(
							[]cty.Type{
								cty.String,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test",
						knownvalue.TupleExact(
							[]knownvalue.Check{
								knownvalue.StringExact("one"),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicVariadicFunction_value_multiple_same_type_primitive(t *testing.T) {
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
					value = provider::framework::dynamic_variadic("one", "two")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test",
						cty.Tuple(
							[]cty.Type{
								cty.String,
								cty.String,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test",
						knownvalue.TupleExact(
							[]knownvalue.Check{
								knownvalue.StringExact("one"),
								knownvalue.StringExact("two"),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicVariadicFunction_value_multiple_same_type_complex(t *testing.T) {
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
					value = provider::framework::dynamic_variadic(
						{a = 1234.5, b = true, c = "hello"},
						{a = 200, b = false, c = "world"},
					)
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test",
						cty.Tuple(
							[]cty.Type{
								cty.Object(map[string]cty.Type{
									"a": cty.Number,
									"b": cty.Bool,
									"c": cty.String,
								}),
								cty.Object(map[string]cty.Type{
									"a": cty.Number,
									"b": cty.Bool,
									"c": cty.String,
								}),
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test",
						knownvalue.TupleExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"a": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"b": knownvalue.Bool(true),
										"c": knownvalue.StringExact("hello"),
									},
								),
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"a": knownvalue.NumberExact(big.NewFloat(200)),
										"b": knownvalue.Bool(false),
										"c": knownvalue.StringExact("world"),
									},
								),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicVariadicFunction_value_multiple_different_type(t *testing.T) {
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
					value = provider::framework::dynamic_variadic(true, "string", 1234.5)
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Tuple(
							[]cty.Type{
								cty.Bool,
								cty.String,
								cty.Number,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test",
						knownvalue.TupleExact(
							[]knownvalue.Check{
								knownvalue.Bool(true),
								knownvalue.StringExact("string"),
								knownvalue.NumberExact(big.NewFloat(1234.5)),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicVariadicFunction_null(t *testing.T) {
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
					value = provider::framework::dynamic_variadic(null)
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestDynamicVariadicFunction_typed_null(t *testing.T) {
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
					value = provider::framework::dynamic_variadic(var.typed_null)
				}
				
				variable "typed_null" {
					type = bool
					default = null
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestDynamicVariadicFunction_unknown(t *testing.T) {
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
					input = provider::framework::dynamic_variadic("test-value")
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
					cornertesting.ExpectOutputType("test",
						cty.Tuple(
							[]cty.Type{
								cty.String,
							},
						),
					),
					statecheck.ExpectKnownOutputValue("test",
						knownvalue.TupleExact(
							[]knownvalue.Check{
								knownvalue.StringExact("test-value"),
							},
						),
					),
				},
			},
		},
	})
}
