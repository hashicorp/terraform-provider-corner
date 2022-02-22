package tf6to5provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf6to5server"
	"github.com/hashicorp/terraform-provider-corner/internal/tf6to5provider/provider"
)

func New() (func() tfprotov5.ProviderServer, error) {
	ctx := context.Background()

	downgradeServer, err := tf6to5server.DowngradeServer(ctx, func() tfprotov6.ProviderServer {
		return tfsdk.NewProtocol6Server(provider.New())
	})

	if err != nil {
		return nil, err
	}

	return func() tfprotov5.ProviderServer {
		return downgradeServer
	}, nil
}
