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
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// This test asserts a bug that has yet to be fixed in plugin framework with defaults being used in an attribute inside of a set.
//
// This bug can be observed with various different outcomes: producing duplicate set element errors, incorrect diffs during plan,
// consistent diffs with values switching back and forth, etc. Example bug reports:
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/783
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/867
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/1036
//
// They all originate from the same root cause, which is when using `Default` on multiple attributes inside of a set, when one default
// value is applied, the other default values may also be applied due to the set element being modified during traversal. The reason this
// results in differing behavior is because Terraform core can't apply data consistency rules to sets that contain objects, so instead of
// a single consistent error message, we get a bunch of different errors/odd behavior depending on the exact result of the defaulting logic.
//
// This specific test will successfully apply with the correct data, then following refresh/plan/apply commands will all raise
// a "duplicate set element" error. Since the framework logic defaults the set elements while traversing it and set elements are identified by their
// value, follow-up path lookups for other set elements can't find the correct data, resulting in default values being applied incorrectly
// for the "value" attributes.
//
//	Error: Duplicate Set Element
//
//	with framework_set_nested_block_with_defaults.test,
//	on terraform_plugin_test.tf line 11, in resource "framework_set_nested_block_with_defaults" "test":
//	11: resource "framework_set_nested_block_with_defaults" "test" {
//
//	This attribute contains duplicate values of:
//	tftypes.Object["default_value":tftypes.String,
//	"value":tftypes.String]<"default_value":tftypes.String<"this is a default">,
//	"value":tftypes.String<"zero">>
//
// Once this bug is fixed, the ExpectError regex in this test should be removed and the plan check should be switched to a state check.
func TestSetNestedBlockWithDefaults(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// The "Duplicate Set Element" error was introduced in Terraform 1.4
			tfversion.SkipBelow(tfversion.Version1_4_0),
		},
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
