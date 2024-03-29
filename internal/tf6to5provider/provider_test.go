// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6to5provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"tf6to5provider": func() (tfprotov5.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "age", "123"),
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "email", "example@example.com"),
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "name", "Example Name"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf6to5provider_user" "example" {
  age   = 123
  email = "example@example.com"
  name  = "Example Name"
}
`

func TestAccFunctionString(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"tf6to5provider": func() (tfprotov5.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test1" {
					value = provider::tf6to5provider::string("str")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test1", knownvalue.StringExact("str")),
				},
			},
		},
	})
}
