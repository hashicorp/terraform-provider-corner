// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceUser2(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("tf5muxprovider_user2.example", "name", regexp.MustCompile("^Exam")),
				),
			},
		},
	})
}

const configResourceBasic = `
resource "tf5muxprovider_user2" "example" {
  age   = 200
  email = "example@example.com"
  name  = "Example Name"
}
`
