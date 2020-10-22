package warningprovider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Currently the test framework does not provide an option for
// testing warnings, so this provider is mostly tested manually.

func TestAccWarningProviderResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"corner-warning": func() (*schema.Provider, error) {
				return New(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: config("foo"),
			},
			{
				ResourceName:      "corner_warning_only.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: config("foo1"),
			},
			{
				ResourceName:      "corner_warning_only.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(id string) string {
	return fmt.Sprintf(`
terraform {
	required_providers {
		corner = {
			source = "hashicorp/corner-warning"
		}
	}
}

data "corner_warning_only" "test" {
}

resource "corner_warning_only" "test" {
	set_id = %q
}  
`, id)
}
