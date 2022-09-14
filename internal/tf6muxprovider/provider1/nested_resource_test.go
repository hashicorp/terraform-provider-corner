package provider1

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceNested(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6muxprovider": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceNestedBasic,
			},
		},
	})
}

const configResourceNestedBasic = `
resource "tf6muxprovider_nested" "example" {
  set {
    id = "one"

    list {

    }

    list {

    }
  }

  set {

  }
}
`
