// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccResourceUserCty(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceUserCtyBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"corner_user.foo", "email", "ford@prefect.co"),
					resource.TestCheckResourceAttr(
						"corner_user.foo", "name", "Ford Prefect"),
					resource.TestCheckResourceAttr(
						"corner_user.foo", "age", "200"),
				),
			},
			{
				Config: configResourceUserCtyUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"corner_user.foo", "email", "ford@prefect.co"),
					resource.TestCheckResourceAttr(
						"corner_user.foo", "name", "Ford Prefect II"),
					resource.TestCheckResourceAttr(
						"corner_user.foo", "age", "300"),
				),
			},
		},
	}
}

const configResourceUserCtyBasic = `
resource "corner_user" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
}
`

const configResourceUserCtyUpdate = `
resource "corner_user" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect II"
  age = 300
}
`
