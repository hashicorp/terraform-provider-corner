// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWriteOnlyReplaceExResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly_replace_ex" "test" {
					string_attr = "hello!"
					writeonly_string = "write-only value"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Create
						plancheck.ExpectResourceAction("framework_writeonly_replace_ex.test", plancheck.ResourceActionCreate),
					},
				},
			},
			{
				Config: `resource "framework_writeonly_replace_ex" "test" {
					string_attr = "hello!"
					writeonly_string = "write-only value"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// No-op, however, requires_replace is populated with a path to "writeonly_string"
						plancheck.ExpectResourceAction("framework_writeonly_replace_ex.test", plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: `resource "framework_writeonly_replace_ex" "test" {
					string_attr = "goodbye!"
					writeonly_string = "write-only value"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Update, however, requires_replace is populated with a path to "writeonly_string"
						plancheck.ExpectResourceAction("framework_writeonly_replace_ex.test", plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}
