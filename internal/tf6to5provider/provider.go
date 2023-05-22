// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6to5provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf6to5server"
	"github.com/hashicorp/terraform-provider-corner/internal/tf6to5provider/provider"
)

func New() (func() tfprotov5.ProviderServer, error) {
	ctx := context.Background()

	downgradeServer, err := tf6to5server.DowngradeServer(ctx, providerserver.NewProtocol6(provider.New()))

	if err != nil {
		return nil, err
	}

	return func() tfprotov5.ProviderServer {
		return downgradeServer
	}, nil
}
