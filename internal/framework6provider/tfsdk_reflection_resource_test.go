// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestTFSDKReflectionResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_tfsdk_reflection" "test" {
					config_bool = true
					config_dynamic = "hello"
					config_float64 = 2.1
					config_int64 = 201
					config_string = "hello"

					config_list = ["one", "two"]
					config_map = {
						key1 = "val1"
						key2 = "val2"
						key3 = "val3"
					}
					config_map_nested = {
						key1 = {
							nested_string = "val1"
						}
						key2 = {
							nested_string = "val2"
						}
						key3 = {
							nested_string = "val3"
						}
					}

					config_object = {
						nested_string = "hello"
					}
					config_single_nested = {
						nested_string = "hello"
					}
					
					config_set_nested_block {
						nested_string = "one"
					}
					config_set_nested_block {
						nested_string = "two"
					}

					# Blocks themselves are not computed, but the attributes within are. Since it's hardcoded
					# in the resource Create function, this configuration ensures no plan consistency errors
					computed_list_nested_block{}
					computed_list_nested_block{}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_bool"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_dynamic"), knownvalue.StringExact("dynamic string")),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_float64"), knownvalue.Float64Exact(1.2)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_int64"), knownvalue.Int64Exact(100)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_string"), knownvalue.StringExact("computed string")),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_list"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("computed"),
								knownvalue.StringExact("list"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_map"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key6": knownvalue.StringExact("val6"),
								"key7": knownvalue.StringExact("val7"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_map_nested"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key6": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("val6"),
									},
								),
								"key7": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("val7"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_object"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"nested_string": knownvalue.StringExact("computed string"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_single_nested"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"nested_string": knownvalue.StringExact("computed string"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_list_nested_block"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("computed string one"),
									},
								),
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("computed string two"),
									},
								),
							},
						),
					),
				},
			},
			{
				Config: `resource "framework_tfsdk_reflection" "test" {
					config_bool = false
					config_dynamic = "world"
					config_float64 = 3.4
					config_int64 = 301
					config_string = "world"

					config_list = ["three", "four", "five"]
					config_map = {
						key4 = "val4"
						key5 = "val5"
					}
					config_map_nested = {
						key4 = {
							nested_string = "val4"
						}
						key5 = {
							nested_string = "val5"
						}
					}

					config_object = {
						nested_string = "world"
					}
					config_single_nested = {
						nested_string = "world"
					}
					
					config_set_nested_block {
						nested_string = "three"
					}
					config_set_nested_block {
						nested_string = "four"
					}

					# Blocks themselves are not computed, but the attributes within are. Since it's hardcoded
					# in the resource Create function, this configuration ensures no plan consistency errors
					computed_list_nested_block{}
					computed_list_nested_block{}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_bool"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_dynamic"), knownvalue.StringExact("dynamic string")),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_float64"), knownvalue.Float64Exact(1.2)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_int64"), knownvalue.Int64Exact(100)),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_string"), knownvalue.StringExact("computed string")),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_list"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("computed"),
								knownvalue.StringExact("list"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_map"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key6": knownvalue.StringExact("val6"),
								"key7": knownvalue.StringExact("val7"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_map_nested"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key6": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("val6"),
									},
								),
								"key7": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("val7"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_object"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"nested_string": knownvalue.StringExact("computed string"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_single_nested"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"nested_string": knownvalue.StringExact("computed string"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_tfsdk_reflection.test", tfjsonpath.New("computed_list_nested_block"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("computed string one"),
									},
								),
								knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"nested_string": knownvalue.StringExact("computed string two"),
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
