// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestRefinementConsumerResource_conflicting_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						conflicting_bool_one = framework_refinement_producer.test.bool_with_not_null
						conflicting_bool_two = framework_refinement_producer.test.bool_with_not_null
					}
				`,
				// Even though "framework_refinement_producer.test.bool_with_not_null" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Error: Invalid Attribute Combination`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_most_string_length_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_most_string_length = framework_refinement_producer.test.string_with_prefix
					}
				`,
				// Even though "framework_refinement_producer.test.string_with_prefix" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_most_string_length string length must be at most 8`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_most_int64_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_most_int64 = framework_refinement_producer.test.int64_with_bounds
					}
				`,
				// Even though "framework_refinement_producer.test.int64_with_bounds" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_most_int64 value must be at most 9`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_least_float64_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_least_float64 = framework_refinement_producer.test.float64_with_bounds
					}
				`,
				// Even though "framework_refinement_producer.test.float64_with_bounds" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_least_float64 value must be at least 20.235000`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_least_list_size_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_least_list_size = framework_refinement_producer.test.list_with_length_bounds
					}
				`,
				// Even though "framework_refinement_producer.test.list_with_length_bounds" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_least_list_size list must contain at least 6 elements`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_most_list_size_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_most_list_size = framework_refinement_producer.test.list_with_length_bounds
					}
				`,
				// Even though "framework_refinement_producer.test.list_with_length_bounds" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_most_list_size list must contain at most 1 elements`),
			},
		},
	})
}

func TestRefinementConsumerResource_at_least_set_size_validation_error(t *testing.T) {
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
				Config: `
					resource "framework_refinement_producer" "test" {}

					resource "framework_refinement_consumer" "test" {
						at_least_set_size = framework_refinement_producer.test.list_with_length_bounds
					}
				`,
				// Even though "framework_refinement_producer.test.list_with_length_bounds" is unknown, there is
				// enough information to fail validation early.
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Attribute at_least_set_size set must contain at least 6 elements`),
			},
		},
	})
}
