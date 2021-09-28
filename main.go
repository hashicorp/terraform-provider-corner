package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"
	protocol "github.com/hashicorp/terraform-provider-corner/internal/protocolprovider"
	sdkv2 "github.com/hashicorp/terraform-provider-corner/internal/sdkv2provider"
)

func main() {
	ctx := context.Background()
	muxed, err := tfmux.NewSchemaServerFactory(ctx, sdkv2.New().GRPCProvider, protocol.Server)
	if err != nil {
		panic(err)
	}

	err = tf5server.Serve("registry.terraform.io/hashicorp/corner", func() tfprotov5.ProviderServer {
		return muxed.Server()
	})
	if err != nil {
		panic(err)
	}
}
