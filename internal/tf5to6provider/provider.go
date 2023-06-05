// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5to6provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-provider-corner/internal/tf5to6provider/provider"
)

func New() (func() tfprotov6.ProviderServer, error) {
	ctx := context.Background()

	upgradeServer, err := tf5to6server.UpgradeServer(ctx, provider.New().GRPCProvider)

	if err != nil {
		return nil, err
	}

	return func() tfprotov6.ProviderServer {
		return upgradeServer
	}, nil
}
