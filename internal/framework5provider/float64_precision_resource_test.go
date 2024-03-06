// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// This test verifies that the correct precision value is stored in state
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
func TestSchemaResource_Float64Attribute_Precision(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_float64_precision" "test" {
					float64_attribute = 1 - 0.99
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float64_precision.test", "float64_attribute", "0.010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"),
				),
			},
			{
				Config: `resource "framework_float64_precision" "test" {
					float64_attribute = 1 - 0.98
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float64_precision.test", "float64_attribute", "0.020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006"),
				),
			},
		},
	})
}

// This test verifies that a plan is detected when encountering a more precise value
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
func TestSchemaResource_Float64Attribute_Precision_Plan(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_float64_precision" "test" {
					float64_attribute = 0.01
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float64_precision.test", "float64_attribute", "0.01"),
				),
			},
			{
				Config: `resource "framework_float64_precision" "test" {
					float64_attribute = 1 - 0.99
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_float64_precision.test", "float64_attribute", "0.010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"),
				),
			},
		},
	})
}
