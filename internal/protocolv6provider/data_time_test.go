package protocolv6

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccV6DataSourceTime(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": func() (tfprotov6.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTimeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.corner_v6_time.foo", "current", regexp.MustCompile(`[0-9]+`)),
				),
			},
		},
	})
}

var testAccDataSourceTimeConfig = `data "corner_v6_time" "foo" {

}`
