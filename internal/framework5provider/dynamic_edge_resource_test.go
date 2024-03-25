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
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// TODO: The computed dynamic value type should be allowed to change, bug will be fixed with:
// - https://github.com/hashicorp/terraform-plugin-framework/issues/969
func TestDynamicEdge_computed_type_changes(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_dynamic_edge" "test" {
					required_dynamic = "value1"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_dynamic_edge.test", tfjsonpath.New("required_dynamic"), knownvalue.StringExact("value1")),
					// Created as a boolean
					statecheck.ExpectKnownValue("framework_dynamic_edge.test", tfjsonpath.New("computed_dynamic_type_changes"), knownvalue.Bool(true)),
				},
			},
			{
				Config: `resource "framework_dynamic_edge" "test" {
					required_dynamic = "new value"
				}`,
				// ConfigStateChecks: []statecheck.StateCheck{
				// 	statecheck.ExpectKnownValue("framework_dynamic_edge.test", tfjsonpath.New("required_dynamic"), knownvalue.StringExact("new value")),
				// 	// After update, it's a number!
				// 	statecheck.ExpectKnownValue("framework_dynamic_edge.test", tfjsonpath.New("computed_dynamic_type_changes"), knownvalue.Int64Exact(200)),
				// },
				ExpectError: regexp.MustCompile(`unexpected new value: .computed_dynamic_type_changes: wrong final value type:\nstring required.`),
			},
		},
	})
}
