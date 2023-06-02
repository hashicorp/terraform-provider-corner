// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser2(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6muxprovider": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "age", "123"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "email", "example@example.com"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "id", "h"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "name", "Example Name"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf6muxprovider_user2" "example" {
  age   = 123
  email = "example@example.com"
  id    = "h"
  name  = "Example Name"
}
`
