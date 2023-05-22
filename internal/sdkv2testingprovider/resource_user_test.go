// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2testingprovider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccResourceUser(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"corner_user.foo", "name", regexp.MustCompile("^For")),
				),
			},
		},
	}
}

const configResourceBasic = `
resource "corner_user" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
}
`
