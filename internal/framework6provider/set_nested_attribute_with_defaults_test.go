// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
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
//	Error: Duplicate Set Element
//
//	with framework_set_nested_attribute_with_defaults.test,
//	on terraform_plugin_test.tf line 12, in resource "framework_set_nested_attribute_with_defaults" "test":
//	12: 					set = [
//	13: 						{ value = "one" },
//	14: 						{ value = "two" },
//	15: 						{ value = "three" },
//	16: 					]
//
//	This attribute contains duplicate values of:
//	tftypes.Object["default_value":tftypes.String,
//	"value":tftypes.String]<"default_value":tftypes.String<"this is a default">,
//	"value":tftypes.String<"zero">>
//
// Once this bug is fixed, the ExpectError regex in this test should be removed and the plan check should be switched to a state check.
//
// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/783
func TestSetNestedAttributeWithDefaults(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_set_nested_attribute_with_defaults" "test" {
					set = [
						{ value = "one" },
						{ value = "two" },
						{ value = "three" },
					]
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(
							"framework_set_nested_attribute_with_defaults.test",
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
