// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceTime(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			//nolint:unparam // False positive in unparam related to map: https://github.com/mvdan/unparam/issues/40
			"corner": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTimeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.corner_time.foo", "current", regexp.MustCompile(`[0-9]+`)),
				),
			},
		},
	})
}

var testAccDataSourceTimeConfig = `data "corner_time" "foo" {

}`
