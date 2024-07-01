// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// This test verifies that the correct precision value is stored in state
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
func TestSchemaResource_Float32Attribute_Precision(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
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
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
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
func TestSchemaResource_Float32Attribute_Precision_MaxFloat32(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			// Error when planned with 64-bit precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528859811704183484516925440
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*planned value\s{0,10}cty\.NumberIntVal\(3\.4028234663852886e\+38\) does not match config value\s{0,10}cty\.NumberIntVal\(3\.4028234663852885981170418348451692544e\+38\)`),
			},
			// Error when planned with 64-bit precision (scientific notation)
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+38
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*planned value\s{0,10}cty\.NumberIntVal\(3\.4028234663852886e\+38\) does not match config value\s{0,10}cty\.NumberIntVal\(3\.4028234663852885981170418348451692544e\+38\)`),
			},
			// No error when planned with 32-bit precision
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 340282346638528860000000000000000000000
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528860000000000000000000000"),
				),
			},
			// Semantic equality with 32-bit scientific notation representation
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
			// No plan is detected when 32-bit value is replaced with 64-bit value
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

func TestSchemaResource_Float32Attribute_Precision_Overflow_Underflow(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			// float32 overflow
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 3.40282346638528859811704183484516925440e+39
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Value %!s\(\*big\.Float=3\.402823466e\+39\) cannot be represented as a\s{0,10}32-bit floating point.`),
			},
			// float32 negative overflow
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = -3.40282346638528859811704183484516925440e+39
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Value %!s\(\*big\.Float=-3\.402823466e\+39\) cannot be represented as a\s{0,10}32-bit floating point.`),
			},
			// float32 underflow
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = 1.401298464324817070923729583289916131280e-46
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Value %!s\(\*big\.Float=1\.401298464e-46\) cannot be represented as a\s{0,10}32-bit floating point.`),
			},
			// float32 negative underflow
			{
				Config: `resource "framework_float32_precision" "test" {
					float32_attribute = -1.401298464324817070923729583289916131280e-46
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float32_precision.test", "float32_attribute", "340282346638528859811704183484516925440"),
				),
				ExpectError: regexp.MustCompile(`.*Error: Value %!s\(\*big\.Float=-1\.401298464e-46\) cannot be represented as a\s{0,10}32-bit floating point.`),
			},
		},
	})
}
