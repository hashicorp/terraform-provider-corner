package sdkv2testingprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccResourceBigint(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBigint,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("corner_bigint.foo", "number", "7227701560655103598"),
					resource.TestCheckResourceAttr("corner_bigint.foo", "int64", "7227701560655103598"),
				),
			},
		},
	}
}

const configResourceBigint = `
resource "corner_bigint" "foo" {
  number = 7227701560655103598
}
`
