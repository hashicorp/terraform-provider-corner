// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6muxprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6muxprovider": func() (tfprotov6.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "age", "123"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "email", "example1@example.com"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "id", "h"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "name", "Example Name 1"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "age", "234"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "email", "example2@example.com"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "id", "h"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "name", "Example Name 2"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf6muxprovider_user1" "example" {
  age   = 123
  email = "example1@example.com"
  id    = "h"
  name  = "Example Name 1"
}

resource "tf6muxprovider_user2" "example" {
  age   = 234
  email = "example2@example.com"
  id    = "h"
  name  = "Example Name 2"
}
`
