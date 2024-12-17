// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself, to verify the config data is passed to the resource Create function.
// The state check assertions cannot be used for testing this data as write-only data should never be persisted in state.

func TestWriteOnlyResource_CustomType(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_custom_ipv6 = "::"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
				},
			},
		},
	})
}

func TestWriteOnlyResource_Primitive(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_string = "fakepassword"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
				},
			},
		},
	})
}

func TestWriteOnlyResource_Collection(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_set = ["fake", "password"]
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
				},
			},
		},
	})
}

func TestWriteOnlyResource_NestedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
					writeonly_nested_object = {
						writeonly_int64 = 1234
						writeonly_bool = true

						writeonly_nested_list = [
							{
								writeonly_string = "fakepassword1"
							},
							{
								writeonly_string = "fakepassword2"
							}
						]
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("nested_block_list"), knownvalue.ListSizeExact(0)),
				},
			},
		},
	})
}

func TestWriteOnlyResource_NestedBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly" "test" {
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
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
						plancheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
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
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_custom_ipv6"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_set"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_writeonly.test", tfjsonpath.New("writeonly_nested_object"), knownvalue.Null()),
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
		},
	})
}
