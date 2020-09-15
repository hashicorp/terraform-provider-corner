package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceBasic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"basic_data_source.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const configDataSourceBasic = `
resource "basic_data_source" "foo" {
  sample_attribute = "bar"
}
`
