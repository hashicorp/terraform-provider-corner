// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccDataSourceBigint(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configDataSourceBigint,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.corner_bigint.foo", "int64", "7227701560655103598")),
			},
		},
	}
}

const configDataSourceBigint = `
data "corner_bigint" "foo" {
}
`
