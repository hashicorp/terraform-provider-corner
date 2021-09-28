package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccDataSourceRegionsCty(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configDataSourceRegionsCtyBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.corner_regions_cty.foo", "names.#")),
			},
		},
	}
}

const configDataSourceRegionsCtyBasic = `
data "corner_regions_cty" "foo" {
  filter = "foo"
}
`
