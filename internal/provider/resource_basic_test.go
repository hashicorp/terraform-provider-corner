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
						"corner_person.foo", "name", regexp.MustCompile("^For")),
				),
			},
		},
	}
}

const configResourceBasic = `
resource "corner_person" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
}
`
