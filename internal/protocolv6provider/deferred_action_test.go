// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccDataSourceDeferredAction_InvalidDeferredResponse(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Only run this test for Terraform clients that do not support deferred actions
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipAbove(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(false), nil
			},
		},
		// Test that terraform-plugin-go throws an error diagnostic when a
		// deferral is sent without the deferred action client capability
		Steps: []resource.TestStep{
			{
				Config:      `data "corner_v6_deferred_action" "foo" {}`,
				ExpectError: regexp.MustCompile("Error: Invalid Deferred Response"),
			},
		},
	})
}
