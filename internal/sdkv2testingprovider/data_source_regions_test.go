// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2testingprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccDataSourceRegions(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configDataSourceBasic,
				//nolint:staticcheck //Deprecated functions
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.corner_regions.foo", "names.#")),
			},
		},
	}
}

const configDataSourceBasic = `
data "corner_regions" "foo" {
}
`
