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

func TestWriteOnlyUpgradeResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"framework": providerserver.NewProtocol5WithError(NewWithUpgradeVersion(0)),
				},
				Config: `resource "framework_writeonly_upgrade" "test" {
					string_attr = "hello!"
					writeonly_string = "fakepassword"
				}`,
			},
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"framework": providerserver.NewProtocol5WithError(NewWithUpgradeVersion(1)),
				},
				Config: `resource "framework_writeonly_upgrade" "test" {
					string_attr = "world!"
					writeonly_string = "fakepassword"
				}`,
				// TODO: Remove this expect error once Framework is updated to null out write-only attributes.
				ExpectError: regexp.MustCompile(`Error: Provider produced invalid object`),
			},
			// TODO: Remove this additional step once Framework is updated to null out write-only attributes.
			// Back to the original config to avoid a destroy clean-up error.
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"framework": providerserver.NewProtocol5WithError(NewWithUpgradeVersion(0)),
				},
				Config: `resource "framework_writeonly_upgrade" "test" {
					string_attr = "hello!"
					writeonly_string = "fakepassword"
				}`,
			},
		},
	})
}
