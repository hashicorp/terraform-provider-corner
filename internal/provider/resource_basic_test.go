package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccResourceBasic(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"basic_resource.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	}
}

const configResourceBasic = `
resource "basic_resource" "foo" {
  sample_attribute = "bar"
}
`
