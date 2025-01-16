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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// This resource tests Terraform's data consistency rules for refinements. It has an invalid
// promise (string must have prefix of "prefix://") in the plan, which will fail during apply.
func TestInvalidRefinementResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// Unknown value refinements were introduced to Terraform v1.6.0 via go-cty
			tfversion.SkipBelow(tfversion.Version1_6_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
					resource "framework_invalid_refinement" "test" {}
				
					resource "terraform_data" "test_out" {
						count = startswith(framework_invalid_refinement.test.string_with_prefix, "prefix://") ? 1 : 0
					}
				`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("framework_invalid_refinement.test", tfjsonpath.New("string_with_prefix")),
						plancheck.ExpectResourceAction("terraform_data.test_out[0]", plancheck.ResourceActionCreate),
					},
				},
				ExpectError: regexp.MustCompile(`Error: Provider produced inconsistent result after apply`),
			},
		},
	})
}
