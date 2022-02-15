package tf6muxprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-provider-corner/internal/tf6muxprovider/provider1"
	"github.com/hashicorp/terraform-provider-corner/internal/tf6muxprovider/provider2"
)

func New() (func() tfprotov6.ProviderServer, error) {
	ctx := context.Background()
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return tfsdk.NewProtocol6Server(provider1.New())
		},
		func() tfprotov6.ProviderServer {
			return tfsdk.NewProtocol6Server(provider2.New())
		},
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		return nil, err
	}

	return muxServer.ProviderServer, nil
}
