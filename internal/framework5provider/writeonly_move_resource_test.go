// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWriteOnlyMoveResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "terraform_data" "test" {
					input = "hello world!"
				}`,
			},
			{
				Config: `resource "framework_writeonly_move" "test" {
					string_attr = "hello world!"
					writeonly_string = "fakepassword"
				}

				moved {
					from = terraform_data.test
					to   = framework_writeonly_move.test
				}`,
				// TODO: Remove this expect error once Framework is updated to null out write-only attributes.
				ExpectError: regexp.MustCompile(`Error: Provider returned invalid value`),
			},
			// TODO: Remove this additional step once Framework is updated to null out write-only attributes.
			// Back to the original config to avoid a destroy clean-up error.
			{
				Config: `resource "terraform_data" "test" {
					input = "hello world!"
				}`,
			},
		},
	})
}
