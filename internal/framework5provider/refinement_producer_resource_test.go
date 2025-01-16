// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
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

func TestRefinementProducerResource_basic_pre_1_3(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// This test runs against earlier Terraform versions to ensure provider-returned refinements
		// don't cause unexpected issues (core should ignore them)
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Terraform 1.2 and older treat the entire output value as unknown
			tfversion.SkipAbove(tfversion.Version1_2_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
					resource "framework_refinement_producer" "test" {}

					output "test_out" {
						value = framework_refinement_producer.test
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownOutputValue("test_out"),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test_out", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"bool_with_not_null":  knownvalue.Bool(true),
							"int64_with_bounds":   knownvalue.Int64Exact(15),
							"float64_with_bounds": knownvalue.Float64Exact(12.102),
							"list_with_length_bounds": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("hello"),
								knownvalue.StringExact("there"),
								knownvalue.StringExact("world!"),
							}),
							"string_with_prefix": knownvalue.StringExact("prefix://hello-world"),
						},
					)),
				},
			},
		},
	})
}

func TestRefinementProducerResource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// This test runs against earlier Terraform versions to ensure provider-returned refinements don't cause unexpected issues (core should ignore them)
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Terraform 1.3 and above have more fine-grained unknown output values to assert during plan
			tfversion.SkipBelow(tfversion.Version1_3_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
					resource "framework_refinement_producer" "test" {}

					output "test_out" {
						value = framework_refinement_producer.test
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownOutputValueAtPath("test_out", tfjsonpath.New("bool_with_not_null")),
						plancheck.ExpectUnknownOutputValueAtPath("test_out", tfjsonpath.New("int64_with_bounds")),
						plancheck.ExpectUnknownOutputValueAtPath("test_out", tfjsonpath.New("float64_with_bounds")),
						plancheck.ExpectUnknownOutputValueAtPath("test_out", tfjsonpath.New("list_with_length_bounds")),
						plancheck.ExpectUnknownOutputValueAtPath("test_out", tfjsonpath.New("string_with_prefix")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test_out", knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"bool_with_not_null":  knownvalue.Bool(true),
							"int64_with_bounds":   knownvalue.Int64Exact(15),
							"float64_with_bounds": knownvalue.Float64Exact(12.102),
							"list_with_length_bounds": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("hello"),
								knownvalue.StringExact("there"),
								knownvalue.StringExact("world!"),
							}),
							"string_with_prefix": knownvalue.StringExact("prefix://hello-world"),
						},
					)),
				},
			},
		},
	})
}

func TestRefinementProducerResource_notnull(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support in the provider, this config would return an error like:
				//
				//		Error: Invalid count argument
				//
				//		  on terraform_plugin_test.tf line 15, in resource "terraform_data" "test_out":
				//		  15: 						count = framework_refinement_producer.test.bool_with_not_null != null ? 1 : 0
				//
				//		The "count" value depends on resource attributes that cannot be determined
				//		until apply, so Terraform cannot predict how many instances will be created.
				//		To work around this, use the -target argument to first apply only the
				//		resources that the count depends on.
				//
				Config: `
					resource "framework_refinement_producer" "test" {}
				
					resource "terraform_data" "test_out" {
						count = framework_refinement_producer.test.bool_with_not_null != null ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_refinement_producer.test", tfjsonpath.New("bool_with_not_null")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_refinement_producer.test", tfjsonpath.New("bool_with_not_null"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestRefinementProducerResource_stringprefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support in the provider, this config would return an error like:
				//
				//		Error: Invalid count argument
				//
				//		  on terraform_plugin_test.tf line 15, in resource "terraform_data" "test_out":
				//		  15: 						count = startswith(framework_refinement_producer.test.string_with_prefix, "prefix://") ? 1 : 0
				//
				//		The "count" value depends on resource attributes that cannot be determined
				//		until apply, so Terraform cannot predict how many instances will be created.
				//		To work around this, use the -target argument to first apply only the
				//		resources that the count depends on.
				//
				Config: `
					resource "framework_refinement_producer" "test" {}
				
					resource "terraform_data" "test_out" {
						count = startswith(framework_refinement_producer.test.string_with_prefix, "prefix://") ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_refinement_producer.test", tfjsonpath.New("string_with_prefix")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_refinement_producer.test", tfjsonpath.New("string_with_prefix"), knownvalue.StringExact("prefix://hello-world")),
				},
			},
		},
	})
}

func TestRefinementProducerResource_int64_bounds(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support in the provider, this config would return an error like:
				//
				//		Error: Invalid count argument
				//
				//		  on terraform_plugin_test.tf line 15, in resource "terraform_data" "test_out":
				//		  15: 						count = framework_refinement_producer.test.int64_with_bounds > 9 ? 1 : 0
				//
				//		The "count" value depends on resource attributes that cannot be determined
				//		until apply, so Terraform cannot predict how many instances will be created.
				//		To work around this, use the -target argument to first apply only the
				//		resources that the count depends on.
				//
				Config: `
					resource "framework_refinement_producer" "test" {}
				
					resource "terraform_data" "test_out" {
						count = framework_refinement_producer.test.int64_with_bounds > 9 ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_refinement_producer.test", tfjsonpath.New("int64_with_bounds")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_refinement_producer.test", tfjsonpath.New("int64_with_bounds"), knownvalue.Int64Exact(15)),
				},
			},
		},
	})
}
func TestRefinementProducerResource_float64_bounds(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support in the provider, this config would return an error like:
				//
				//		Error: Invalid count argument
				//
				//		  on terraform_plugin_test.tf line 15, in resource "terraform_data" "test_out":
				//		  15: 						count = framework_refinement_producer.test.float64_with_bounds > 10.233 ? 1 : 0
				//
				//		The "count" value depends on resource attributes that cannot be determined
				//		until apply, so Terraform cannot predict how many instances will be created.
				//		To work around this, use the -target argument to first apply only the
				//		resources that the count depends on.
				//
				Config: `
					resource "framework_refinement_producer" "test" {}
				
					resource "terraform_data" "test_out" {
						count = framework_refinement_producer.test.float64_with_bounds > 10.233 ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_refinement_producer.test", tfjsonpath.New("float64_with_bounds")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_refinement_producer.test", tfjsonpath.New("float64_with_bounds"), knownvalue.Float64Exact(12.102)),
				},
			},
		},
	})
}

func TestRefinementProducerResource_list_length_bounds(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support in the provider, this config would return an error like:
				//
				//		Error: Invalid count argument

				//		  on terraform_plugin_test.tf line 15, in resource "terraform_data" "test_out":
				//		  15: 						count = length(framework_refinement_producer.test.list_with_length_bounds) > 1 ? 1 : 0

				//		The "count" value depends on resource attributes that cannot be determined
				//		until apply, so Terraform cannot predict how many instances will be created.
				//		To work around this, use the -target argument to first apply only the
				//		resources that the count depends on.
				//
				Config: `
					resource "framework_refinement_producer" "test" {}
				
					resource "terraform_data" "test_out" {
						count = length(framework_refinement_producer.test.list_with_length_bounds) > 1 ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_refinement_producer.test", tfjsonpath.New("list_with_length_bounds")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_refinement_producer.test", tfjsonpath.New("list_with_length_bounds"), knownvalue.ListExact(
						[]knownvalue.Check{
							knownvalue.StringExact("hello"),
							knownvalue.StringExact("there"),
							knownvalue.StringExact("world!"),
						},
					)),
				},
			},
		},
	})
}
