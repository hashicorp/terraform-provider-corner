// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// A bug in Terraform v0.12.x results in dynamic null values incorrectly causing a diff during plan.
// These tests will always fail on v0.12.x due to consistently having a non-empty plan after apply.
// The bug originated from go-cty and was resolved in the Terraform v0.13.0 release.
//
// References:
//   - https://github.com/zclconf/go-cty/pull/55
//   - https://github.com/hashicorp/terraform/pull/25216
//
// As this bug is in diff detection, the plan renderer and machine readable plan will not "display"
// any differences, like below:
//
//	An execution plan has been generated and is shown below.
//	Resource actions are indicated with the following symbols:
//	  ~ update in-place
//
//	Terraform will perform the following actions:
//
//	  # framework_dynamic_schema.test will be updated in-place
//	  ~ resource "framework_dynamic_schema" "test" {}
//
//	Plan: 0 to add, 1 to change, 0 to destroy.
func TestDynamicSchemaResource_null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// See comment at top of test file related to dynamic null bug in Terraform v0.12.x
			tfversion.SkipBelow(tfversion.Version0_13_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_schema" "test" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestDynamicSchemaResource_DynamicAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// See comment at top of test file related to dynamic null bug in Terraform v0.12.x
			tfversion.SkipBelow(tfversion.Version0_13_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_schema" "test" {
					dynamic_attribute = "value1"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.StringExact("value1")),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					dynamic_attribute = tolist(["value1", "value2"])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
								knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					dynamic_attribute = {
						"attribute_one": "value1",
						"attribute_two": false,
						"attribute_three": 1234.5,
						"attribute_four": [true, 1234.5],
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"attribute_one":   knownvalue.StringExact("value1"),
								"attribute_two":   knownvalue.Bool(false),
								"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
								"attribute_four": knownvalue.TupleExact(
									[]knownvalue.Check{
										knownvalue.Bool(true),
										knownvalue.NumberExact(big.NewFloat(1234.5)),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestDynamicSchemaResource_ObjectAttributeWithDynamic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// See comment at top of test file related to dynamic null bug in Terraform v0.12.x
			tfversion.SkipBelow(tfversion.Version0_13_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = tolist(["value1", "value2"])
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ListExact(
									[]knownvalue.Check{
										knownvalue.StringExact("value1"),
										knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"attribute_one":   knownvalue.StringExact("value1"),
										"attribute_two":   knownvalue.Bool(false),
										"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"attribute_four": knownvalue.TupleExact(
											[]knownvalue.Check{
												knownvalue.Bool(true),
												knownvalue.NumberExact(big.NewFloat(1234.5)),
											},
										),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestDynamicSchemaResource_SingleNestedBlockWithDynamic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// See comment at top of test file related to dynamic null bug in Terraform v0.12.x
			tfversion.SkipBelow(tfversion.Version0_13_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = tolist(["value1", "value2"])
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ListExact(
									[]knownvalue.Check{
										knownvalue.StringExact("value1"),
										knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
				},
			},
			{
				Config: `resource "framework_dynamic_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_dynamic_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"attribute_one":   knownvalue.StringExact("value1"),
										"attribute_two":   knownvalue.Bool(false),
										"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"attribute_four": knownvalue.TupleExact(
											[]knownvalue.Check{
												knownvalue.Bool(true),
												knownvalue.NumberExact(big.NewFloat(1234.5)),
											},
										),
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
