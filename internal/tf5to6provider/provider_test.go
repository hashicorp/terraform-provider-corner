// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5to6provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf5to6provider": func() (tfprotov6.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf5to6provider_user.example", "age", "123"),
					resource.TestCheckResourceAttr("tf5to6provider_user.example", "email", "example@example.com"),
					resource.TestCheckResourceAttr("tf5to6provider_user.example", "name", "Example Name"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf5to6provider_user" "example" {
  age   = 123
  email = "example@example.com"
  name  = "Example Name"
}
`
