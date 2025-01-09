// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWriteOnlyUpgradeResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"corner": func() (tfprotov5.ProviderServer, error) { //nolint
						return NewWithUpgradeVersion(0).GRPCProvider(), nil
					},
				},
				Config: `resource "corner_writeonly_upgrade" "test" {
					string_attr = "hello!"
					writeonly_string = "fakepassword"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_upgrade.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
				},
			},
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"corner": func() (tfprotov5.ProviderServer, error) { //nolint
						return NewWithUpgradeVersion(1).GRPCProvider(), nil
					},
				},
				Config: `resource "corner_writeonly_upgrade" "test" {
					string_attr = "world!"
					writeonly_string = "fakepassword"
				}`,
				// TODO: This test will start erroring when core fixes the state upgrade bug (currently not checking write-only data sent back)
				// TODO: This test will then pass once SDKv2 fixes the bug of UpgradeResourceState not nulling out write-only data
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly_upgrade.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
				},
			},
		},
	})
}
