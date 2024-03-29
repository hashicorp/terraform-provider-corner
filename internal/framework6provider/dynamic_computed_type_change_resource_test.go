// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/969
// This test confirms that dynamic computed attributes are marked as unknown, both value AND type.
// Dynamic computed attributes can change the type after plan in this scenario.
func TestDynamicComputedTypeChange(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_computed_type_change" "test" {
					required_dynamic = "value1"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("computed_dynamic_type_changes")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("required_dynamic"), knownvalue.StringExact("value1")),
					// Created as a boolean
					statecheck.ExpectKnownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("computed_dynamic_type_changes"), knownvalue.Bool(true)),
				},
			},
			{
				Config: `resource "framework_dynamic_computed_type_change" "test" {
					required_dynamic = "new value"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("computed_dynamic_type_changes")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("required_dynamic"), knownvalue.StringExact("new value")),
					// After update, it's a number!
					statecheck.ExpectKnownValue("framework_dynamic_computed_type_change.test", tfjsonpath.New("computed_dynamic_type_changes"), knownvalue.Int64Exact(200)),
				},
			},
		},
	})
}
