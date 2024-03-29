// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"math/big"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-corner/internal/cornertesting"
	"github.com/zclconf/go-cty/cty"
)

func TestDynamicFunction_known_primitive(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::dynamic("test-value")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test", cty.String),
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact("test-value")),
				},
			},
		},
	})
}

func TestDynamicFunction_known_collection(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::dynamic(tolist([true, false]))
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType("test", cty.List(cty.Bool)),
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.Bool(true),
								knownvalue.Bool(false),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicFunction_known_structural(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::dynamic({
						"attr1": "hello",
						"attr2": 1234.5,
						"attr3": true,
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					cornertesting.ExpectOutputType(
						"test",
						cty.Object(
							map[string]cty.Type{
								"attr1": cty.String,
								"attr2": cty.Number,
								"attr3": cty.Bool,
							},
						),
					),
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"attr1": knownvalue.StringExact("hello"),
								"attr2": knownvalue.NumberExact(big.NewFloat(1234.5)),
								"attr3": knownvalue.Bool(true),
							},
						),
					),
				},
			},
		},
	})
}

func TestDynamicFunction_null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::dynamic(null)
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestDynamicFunction_typed_null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::dynamic(var.typed_null)
				}
				
				variable "typed_null" {
					type = string
					default = null
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestDynamicFunction_unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = provider::framework::dynamic(toset(["hello", "world"]))
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
					cornertesting.ExpectOutputType("test", cty.Set(cty.String)),
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.StringExact("hello"),
								knownvalue.StringExact("world"),
							},
						),
					),
				},
			},
		},
	})
}
