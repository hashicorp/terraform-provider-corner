// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccResourceUserIdentityUpgrade(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"corner": func() (tfprotov5.ProviderServer, error) { //nolint
						return NewWithIdentityUpgradeVersion(0).GRPCProvider(), nil
					},
				},
				Config: `resource "corner_user_identity_upgrade" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 200
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("corner_user_identity_upgrade.foo", map[string]knownvalue.Check{
						"email": knownvalue.StringExact("ford@prefect.co"),
					}),
				},
			},
			{
				ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
					"corner": func() (tfprotov5.ProviderServer, error) { //nolint
						return NewWithIdentityUpgradeVersion(1).GRPCProvider(), nil
					},
				},
				Config: `resource "corner_user_identity_upgrade" "foo" {
					email = "ford@prefect.co"
					name = "Ford Prefect"
					age = 200
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("corner_user_identity_upgrade.foo", map[string]knownvalue.Check{
						"local_part": knownvalue.StringExact("ford"),
						"domain":     knownvalue.StringExact("prefect.co"),
					}),
				},
			},
		},
	}
}
