// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6to5provider": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "age", "123"),
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "email", "example@example.com"),
					resource.TestCheckResourceAttr("tf6to5provider_user.example", "id", "h"),
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
  id    = "h"
  name  = "Example Name"
}
`
