// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccResourceRefinements_Nullness(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				// Without refinement support, this config would return an error like:
				//
				// 		Error: Invalid count argument
				//
				// 		on terraform_plugin_test.tf line 23, in resource "random_string" "other_str":
				// 		23: 					count = corner_v6_refinements.foo.str_value != null ? 1 : 0
				//
				//  	 The "count" value depends on resource attributes that cannot be determined
				//  	 until apply, so Terraform cannot predict how many instances will be created.
				//  	 To work around this, use the -target argument to first apply only the
				//  	 resources that the count depends on.
				//
				// This error occurs because the expression populating "str_value" is passing the provider an unknown value
				// with a refinement (the result will definitely not be null, regardless of the value of "random_string.str.id")
				// that is eventually lost by the provider during PlanResourceChange.
				//
				// When the provider implementation preserves the unknown value refinement for "nullness", this configuration can
				// plan/apply successfully and will create 3 resources, including the "random_string.other_str" resource.
				//
				Config: `
				resource "random_string" "str" {
					length = 12
				}

				resource "corner_v6_refinements" "foo" {
					str_value = "This string is ${random_string.str.id}!"
				}

				resource "random_string" "other_str" {
					count = corner_v6_refinements.foo.str_value != null ? 1 : 0
					length = 12
				}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("corner_v6_refinements.foo", tfjsonpath.New("str_value")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_v6_refinements.foo", tfjsonpath.New("str_value"), knownvalue.StringRegexp(regexp.MustCompile(`This string is`))),
					statecheck.ExpectKnownValue("random_string.other_str[0]", tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccResourceRefinements_Prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			// Without refinement support, this config would return an error like:
			//
			// 		Error: Invalid count argument
			//
			// 		on terraform_plugin_test.tf line 25, in resource "random_string" "other_str":
			// 		25: 					count = startswith(corner_v6_refinements.foo.str_value, "known-prefix-") ? 1 : 0
			//
			//  	 The "count" value depends on resource attributes that cannot be determined
			//  	 until apply, so Terraform cannot predict how many instances will be created.
			//  	 To work around this, use the -target argument to first apply only the
			//  	 resources that the count depends on.
			//
			// This error occurs because the expression populating "str_value" is passing the provider an unknown value
			// with a refinement (the result will definitely not be null and will have a prefix of "known-prefix-", regardless
			// of the value of "random_string.str.id") that is eventually lost by the provider during PlanResourceChange.
			//
			// When the provider implementation preserves this unknown value refinement for string prefixing, this configuration can
			// plan/apply successfully and will create 3 resources, including the "random_string.other_str" resource.
			//
			{
				Config: `
				resource "random_string" "str" {
					length = 12
				}

				resource "corner_v6_refinements" "foo" {
					str_value = "known-prefix-${random_string.str.id}!"
				}

				resource "random_string" "other_str" {
					count = startswith(corner_v6_refinements.foo.str_value, "known-prefix-") ? 1 : 0
					length = 12
				}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("corner_v6_refinements.foo", tfjsonpath.New("str_value")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_v6_refinements.foo", tfjsonpath.New("str_value"), knownvalue.StringRegexp(regexp.MustCompile(`known-prefix-`))),
					statecheck.ExpectKnownValue("random_string.other_str[0]", tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}
