// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5muxprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"tf5muxprovider": func() (tfprotov5.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf5muxprovider_user1.example", "age", "123"),
					resource.TestCheckResourceAttr("tf5muxprovider_user1.example", "email", "example1@example.com"),
					resource.TestCheckResourceAttr("tf5muxprovider_user1.example", "name", "Example Name 1"),
					resource.TestCheckResourceAttr("tf5muxprovider_user2.example", "age", "234"),
					resource.TestCheckResourceAttr("tf5muxprovider_user2.example", "email", "example2@example.com"),
					resource.TestCheckResourceAttr("tf5muxprovider_user2.example", "name", "Example Name 2"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf5muxprovider_user1" "example" {
  age   = 123
  email = "example1@example.com"
  name  = "Example Name 1"
}

resource "tf5muxprovider_user2" "example" {
  age   = 234
  email = "example2@example.com"
  name  = "Example Name 2"
}
`
