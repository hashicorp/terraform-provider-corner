package tf5muxprovider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
)

func main() {
	provider, err := New()

	if err != nil {
		log.Fatalf("unable to create provider: %s", err)
	}

	err = tf5server.Serve("registry.terraform.io/hashicorp/corner", provider)

	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
