// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// This resource tests that semantic equality for elements inside of a set are correctly executed
// Original bug: https://github.com/hashicorp/terraform-plugin-framework/issues/1061
func TestSetSemanticEqualityResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				// The resource Create/Read will return semantically equal data that will cause a diff if returned to Terraform.
				// The semantic equality logic in iptypes.IPv6Address allows this configuration to successfully apply.
				Config: `resource "framework_set_semantic_equality" "test" {
					set_of_ipv6 = [
						"0:0:0:0:0:0:0:0",
						"2001:0DB8:0000:0000:0008:0800:200C:417A",
						"0:0:0:0:0:0:0:101",
						"0:0:0:0:0:FFFF:192.168.255.255",
					]

					set_nested_attribute = [
						{
							ipv6 = "2041:0000:140F:0000:0000:0000:875B:131B"
						},
						{
							ipv6 = "2001:0001:0002:0003:0004:0005:0006:0007"
						},
					]

					set_nested_block {
						ipv6 = "FF01:0:0:0:0:0:0:0"
					}
					set_nested_block {
						ipv6 = "2001:db8::8:800:200c:417a"
					}
					set_nested_block {
						ipv6 = "0:0:0:0:0:FFFF:192.168.255.255"
					}
				}`,
			},
			{
				// Re-ordering the set doesn't produce a diff with semantically equal hardcoded data
				Config: `resource "framework_set_semantic_equality" "test" {
					set_of_ipv6 = [
						"0:0:0:0:0:FFFF:192.168.255.255",
						"0:0:0:0:0:0:0:0",
						"2001:0DB8:0000:0000:0008:0800:200C:417A",
						"0:0:0:0:0:0:0:101",
					]

					set_nested_attribute = [
						{
							ipv6 = "2001:0001:0002:0003:0004:0005:0006:0007"
						},
						{
							ipv6 = "2041:0000:140F:0000:0000:0000:875B:131B"
						},
					]

					set_nested_block {
						ipv6 = "0:0:0:0:0:FFFF:192.168.255.255"
					}
					set_nested_block {
						ipv6 = "FF01:0:0:0:0:0:0:0"
					}
					set_nested_block {
						ipv6 = "2001:db8::8:800:200c:417a"
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("framework_set_semantic_equality.test", plancheck.ResourceActionNoop),
					},
				},
			},
			{
				// User config changes will still produce a diff, but the apply will be successful with semantically equal data
				Config: `resource "framework_set_semantic_equality" "test" {
					set_of_ipv6 = [
						"0:0:0:0:0:FFFF:192.168.255.255",
						"::", # <----------- This update will cause a diff 
						"2001:0DB8:0000:0000:0008:0800:200C:417A",
						"0:0:0:0:0:0:0:101",
					]

					set_nested_attribute = [
						{
							ipv6 = "2001:1:2:3:4:5:6:7" # <----------- This update will cause a diff
						},
						{
							ipv6 = "2041:0000:140F:0000:0000:0000:875B:131B"
						},
					]

					set_nested_block {
						ipv6 = "0:0:0:0:0:FFFF:192.168.255.255"
					}
					set_nested_block {
						ipv6 = "FF01::" # <----------- This update will cause a diff 
					}
					set_nested_block {
						ipv6 = "2001:db8::8:800:200c:417a"
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("framework_set_semantic_equality.test", plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}
