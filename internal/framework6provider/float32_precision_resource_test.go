// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// This test verifies that the correct precision value is stored in state
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
func TestSchemaResource_Float32Attribute_Precision(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 1 - 0.99
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "0.010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"),
				),
			},
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 1 - 0.98
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "0.020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006"),
				),
			},
		},
	})
}

// This test verifies that a plan is detected when encountering a more precise value
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
func TestSchemaResource_Float32Attribute_Precision_Plan(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 0.01
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "0.01"),
				),
			},
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 1 - 0.99
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "0.010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"),
				),
			},
		},
	})
}

// This test documents edge case behavior with math.MaxFloat32
// This number is unique in that it has an exact representation in both float32 and float64,
// it is also an integer, which may affect how this value gets encoded/decoded with go-cty.
// https://github.com/hashicorp/terraform-plugin-framework/issues/1017
func TestSchemaResource_Float32Attribute_Precision_MaxFloat32(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// go-cty v1.14.4 uses string msgpack encoding instead of float msgpack encoding for large whole numbers
			// https://github.com/hashicorp/terraform/pull/34756
			// This changes allows the math.MaxFloat32 value to succeed in planning but fail during apply.
			// Terraform v1.9.0 is the first Terraform version to use this updated encoding.
			tfversion.All(
				tfversion.SkipBelow(tfversion.Version0_15_0),
				tfversion.SkipAbove(tfversion.Version1_8_0),
			),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			// Error when planned with full precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528859811704183484516925440
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*planned value\s{0,10}cty\.NumberIntVal\(3\.4028234663852886e\+38\) does not match config value\s{0,10}cty\.NumberIntVal\(3\.4028234663852885981170418348451692544e\+38\)`),
			},
			// Error when planned with full precision (scientific notation)
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+38
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*planned value\s{0,10}cty\.NumberIntVal\(3\.4028234663852886e\+38\) does not match config value\s{0,10}cty\.NumberIntVal\(3\.4028234663852885981170418348451692544e\+38\)`),
			},
			// No error when planned with 53-bit precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528860000000000000000000000
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
			// Semantic equality with 53-bit precision scientific notation representation
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.4028234663852886e+38
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
			// No plan is detected when 53-bit precision value is replaced with higher precision value
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+38
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
		},
	})
}

// This test documents edge case behavior with math.MaxFloat32 with Terraform v1.9.0 and above
// This number is unique in that it has an exact representation in both float32 and float64,
// it is also an integer, which may affect how this value gets encoded/decoded with go-cty.
// https://github.com/hashicorp/terraform-plugin-framework/issues/1017
//
// go-cty v1.14.4 uses string msgpack encoding instead of float msgpack encoding for large whole numbers
// https://github.com/hashicorp/terraform/pull/34756
// This changes allows the math.MaxFloat32 value to succeed in planning but fail during apply.
// Terraform v1.9.0 is the first Terraform version to use this updated encoding.
func TestSchemaResource_Float32Attribute_Precision_MaxFloat32_TF1_9(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			// Error when planned with full precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528859811704183484516925440
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Provider produced inconsistent result after apply`),
			},
			// Error when planned with full precision (scientific notation)
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+38
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Provider produced inconsistent result after apply`),
			},
			// No error when planned with 53-bit precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528860000000000000000000000
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
			// Semantic equality with 53-bit precision scientific notation representation
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.4028234663852886e+38
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
			// Resource plans an update when 53-bit precision value is replaced with higher precision value but fails during apply
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+38
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("framework_float32_precision.test", plancheck.ResourceActionUpdate),
					},
				},
				ExpectError: regexp.MustCompile(`.*Error: Provider produced inconsistent result after apply`),
			},
		},
	})
}
