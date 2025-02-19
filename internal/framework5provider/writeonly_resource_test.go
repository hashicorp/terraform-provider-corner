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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create function.
func TestWriteOnlyResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_custom_ipv6 = "::"
					writeonly_string = "fakepassword"
					writeonly_list = ["fake", "password"]
					nested_block_list {
						string_attr = "hello"
						writeonly_string = "fakepassword1"

						double_nested_object {
							bool_attr = true
							writeonly_bool = false
						}
					}
					nested_block_list {
						string_attr = "world"
						writeonly_string = "fakepassword2"

						double_nested_object {
							bool_attr = false
							writeonly_bool = true
						}
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_writeonly.test", tfjsonpath.New("computed_attr")),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_list"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":      knownvalue.StringExact("hello"),
								"writeonly_string": knownvalue.Null(),
								"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"bool_attr":      knownvalue.Bool(true),
									"writeonly_bool": knownvalue.Null(),
								}),
							}),
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":      knownvalue.StringExact("world"),
								"writeonly_string": knownvalue.Null(),
								"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"bool_attr":      knownvalue.Bool(false),
									"writeonly_bool": knownvalue.Null(),
								}),
							}),
						})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("computed_attr"), knownvalue.StringExact("computed_val")),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_list"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":      knownvalue.StringExact("hello"),
							"writeonly_string": knownvalue.Null(),
							"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"bool_attr":      knownvalue.Bool(true),
								"writeonly_bool": knownvalue.Null(),
							}),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":      knownvalue.StringExact("world"),
							"writeonly_string": knownvalue.Null(),
							"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"bool_attr":      knownvalue.Bool(false),
								"writeonly_bool": knownvalue.Null(),
							}),
						}),
					})),
				},
			},
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_custom_ipv6 = "::"
					writeonly_string = "fakepassword"
					writeonly_list = ["fake", "password"]
					nested_block_list {
						string_attr = "world"
						writeonly_string = "fakepassword1"

						double_nested_object {
							bool_attr = true
							writeonly_bool = false
						}
					}
					nested_block_list {
						string_attr = "hello"
						writeonly_string = "fakepassword2"

						double_nested_object {
							bool_attr = false
							writeonly_bool = true
						}
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_writeonly.test", tfjsonpath.New("computed_attr")),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_list"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":      knownvalue.StringExact("world"),
								"writeonly_string": knownvalue.Null(),
								"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"bool_attr":      knownvalue.Bool(true),
									"writeonly_bool": knownvalue.Null(),
								}),
							}),
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":      knownvalue.StringExact("hello"),
								"writeonly_string": knownvalue.Null(),
								"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"bool_attr":      knownvalue.Bool(false),
									"writeonly_bool": knownvalue.Null(),
								}),
							}),
						})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("computed_attr"), knownvalue.StringExact("computed_val")),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_list"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":      knownvalue.StringExact("world"),
							"writeonly_string": knownvalue.Null(),
							"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"bool_attr":      knownvalue.Bool(true),
								"writeonly_bool": knownvalue.Null(),
							}),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":      knownvalue.StringExact("hello"),
							"writeonly_string": knownvalue.Null(),
							"double_nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"bool_attr":      knownvalue.Bool(false),
								"writeonly_bool": knownvalue.Null(),
							}),
						}),
					})),
				},
			},
		},
	})
}

func TestWriteOnlyResource_OldTerraformVersion_Error(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Run on all Terraform versions that don't support write-only attributes
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipAbove(tfversion.Version1_10_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_custom_ipv6 = "::"
					nested_block_list {
						string_attr = "hello"
						double_nested_object {
							bool_attr = true
						}
					}
					nested_block_list {
						string_attr = "world"
						double_nested_object {
							bool_attr = false
						}
					}
				}`,
				ExpectError: regexp.MustCompile(`WriteOnly Attribute Not Allowed`),
			},
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_list = ["fake", "password"]
					nested_block_list {
						string_attr = "hello"
						double_nested_object {
							bool_attr = true
						}
					}
					nested_block_list {
						string_attr = "world"
						double_nested_object {
							bool_attr = false
						}
					}
				}`,
				ExpectError: regexp.MustCompile(`WriteOnly Attribute Not Allowed`),
			},
			{
				Config: `resource "framework_writeonly" "test" {
					nested_block_list {
						string_attr = "hello"
						writeonly_string = "fakepassword1"
						double_nested_object {
							bool_attr = true
						}
					}
					nested_block_list {
						string_attr = "world"
						writeonly_string = "fakepassword2"
						double_nested_object {
							bool_attr = false
						}
					}
				}`,
				ExpectError: regexp.MustCompile(`WriteOnly Attribute Not Allowed`),
			},
			{
				Config: `resource "framework_writeonly" "test" {
					nested_block_list {
						string_attr = "hello"
						double_nested_object {
							bool_attr = true
							writeonly_bool = false
						}
					}
					nested_block_list {
						string_attr = "world"
						double_nested_object {
							bool_attr = false
							writeonly_bool = true
						}
					}
				}`,
				ExpectError: regexp.MustCompile(`WriteOnly Attribute Not Allowed`),
			},
		},
	})
}

func TestWriteOnlyResource_NoWriteOnlyValuesSet(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Since there are no write-only values set (despite the schema defining them), this test
		// should pass on all Terraform versions.
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					nested_block_list {
						string_attr = "hello"
						double_nested_object {
							bool_attr = true
						}
					}
					nested_block_list {
						string_attr = "world"
						double_nested_object {
							bool_attr = false
						}
					}
				}`,
			},
		},
	})
}
