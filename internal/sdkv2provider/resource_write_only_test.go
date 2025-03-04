// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create/Update functions.
func TestWriteOnlyResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// TODO: Swap version check with below line once terraform-plugin-testing@v1.12.0 is released
			// tfversion.SkipBelow(tfversion.Version1_11_0),
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  writeonly_bool = true
				  writeonly_string = "fakepassword"
				  writeonly_int = 1234
				  nested_list_block {
				  	string_attr = "hello!"
				  	writeonly_string = "fakepassword"
					double_nested_list_block {
						string_attr = "hello!"
						writeonly_string = "fakepassword"
					}
				  }
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("hello!")),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("nested_list_block"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":                 knownvalue.StringExact("hello!"),
								"opt_or_computed_string_attr": knownvalue.StringExact("computed value!"),
								"writeonly_string":            knownvalue.Null(),
								"double_nested_list_block": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"string_attr":                 knownvalue.StringExact("hello!"),
										"opt_or_computed_string_attr": knownvalue.StringExact("computed value!"),
										"writeonly_string":            knownvalue.Null(),
									}),
								}),
							}),
						})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("hello!")),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("nested_list_block"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":                 knownvalue.StringExact("hello!"),
							"opt_or_computed_string_attr": knownvalue.StringExact("computed value!"),
							"writeonly_string":            knownvalue.Null(),
							"double_nested_list_block": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"string_attr":                 knownvalue.StringExact("hello!"),
									"opt_or_computed_string_attr": knownvalue.StringExact("computed value!"),
									"writeonly_string":            knownvalue.Null(),
								}),
							}),
						}),
					})),
				},
			},
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "world!"
				  writeonly_bool = true
				  writeonly_string = "fakepassword"
				  writeonly_int = 1234
				  nested_list_block {
				  	string_attr = "world!"
					opt_or_computed_string_attr = "config value!"
				  	writeonly_string = "fakepassword"
					double_nested_list_block {
						string_attr = "world!"
						opt_or_computed_string_attr = "config value!"
						writeonly_string = "fakepassword"
					}
				  }
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("world!")),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("nested_list_block"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_attr":                 knownvalue.StringExact("world!"),
								"opt_or_computed_string_attr": knownvalue.StringExact("config value!"),
								"writeonly_string":            knownvalue.Null(),
								"double_nested_list_block": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"string_attr":                 knownvalue.StringExact("world!"),
										"opt_or_computed_string_attr": knownvalue.StringExact("config value!"),
										"writeonly_string":            knownvalue.Null(),
									}),
								}),
							}),
						})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("world!")),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("nested_list_block"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_attr":                 knownvalue.StringExact("world!"),
							"opt_or_computed_string_attr": knownvalue.StringExact("config value!"),
							"writeonly_string":            knownvalue.Null(),
							"double_nested_list_block": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"string_attr":                 knownvalue.StringExact("world!"),
									"opt_or_computed_string_attr": knownvalue.StringExact("config value!"),
									"writeonly_string":            knownvalue.Null(),
								}),
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
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  writeonly_string = "fakepassword"
				  nested_list_block {
				  	string_attr = "hello!"
					double_nested_list_block {
						string_attr = "hello!"
					}
				  }
				}`,
				ExpectError: regexp.MustCompile(`Write-only Attribute Not Allowed`),
			},
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  nested_list_block {
				  	string_attr = "hello!"
				  	writeonly_string = "fakepassword"
					double_nested_list_block {
						string_attr = "hello!"
					}
				  }
				}`,
				ExpectError: regexp.MustCompile(`Write-only Attribute Not Allowed`),
			},
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  nested_list_block {
				  	string_attr = "hello!"
					double_nested_list_block {
						string_attr = "hello!"
						writeonly_string = "fakepassword"
					}
				  }
				}`,
				ExpectError: regexp.MustCompile(`Write-only Attribute Not Allowed`),
			},
		},
	})
}

func TestWriteOnlyResource_NoWriteOnlyValuesSet(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Since there are no write-only values set (despite the schema defining them), this test
		// should pass on all Terraform versions.
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  nested_list_block {
				  	string_attr = "hello!"
					double_nested_list_block {
						string_attr = "hello!"
					}
				  }
				}`,
			},
		},
	})
}
