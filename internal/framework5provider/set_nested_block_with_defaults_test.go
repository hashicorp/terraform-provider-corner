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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// This test asserts a bug that has yet to be fixed in plugin framework with defaults being used in an attribute inside of a set.
//
// This specific test will successfully apply with the correct data, then following refresh/plan/apply commands will all raise
// a "duplicate set element" error. The duplicate set elements are being created by the terraform-plugin-framework default logic
// in PlanResourceChange. The framework logic defaults the set elements while traversing it, and since set elements are identified by their
// value, follow-up path lookups for other set elements can't find the correct data, resulting in default values being applied incorrectly
// for the "value" attributes.
//
// Once this bug is fixed, the ExpectError regex in this test should be removed and the plan check should be switched to a state check.
//
// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/783
func TestSetNestedBlockWithDefaults(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_set_nested_block_with_defaults" "test" {
					set {
						value = "one"
					}
					set {
						value = "two"
					}
					set {
						value = "three"
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(
							"framework_set_nested_block_with_defaults.test",
							tfjsonpath.New("set"),
							knownvalue.SetExact(
								[]knownvalue.Check{
									knownvalue.ObjectExact(
										map[string]knownvalue.Check{
											// During plan and after the first apply, this value will be "one"
											// After the first refresh, the bug will cause this value to be defaulted to "zero"
											"value":         knownvalue.StringExact("one"),
											"default_value": knownvalue.StringExact("this is a default"),
										},
									),
									knownvalue.ObjectExact(
										map[string]knownvalue.Check{
											// During plan and after the first apply, this value will be "two"
											// After the first refresh, the bug will cause this value to be defaulted to "zero"
											"value":         knownvalue.StringExact("two"),
											"default_value": knownvalue.StringExact("this is a default"),
										},
									),
									knownvalue.ObjectExact(
										map[string]knownvalue.Check{
											// During plan and after the first apply, this value will be "three"
											// After the first refresh, the bug will cause this value to be defaulted to "zero"
											"value":         knownvalue.StringExact("three"),
											"default_value": knownvalue.StringExact("this is a default"),
										},
									),
								},
							),
						),
					},
				},
				ExpectError: regexp.MustCompile(`Error: Duplicate Set Element`),
			},
		},
	})
}
