// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"

	protocol "github.com/hashicorp/terraform-provider-corner/internal/protocolprovider"
	sdkv2 "github.com/hashicorp/terraform-provider-corner/internal/sdkv2provider"
)

func main() {
	ctx := context.Background()

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	providers := []func() tfprotov5.ProviderServer{
		protocol.Server,
		sdkv2.New().GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatalf("unable to create provider: %s", err)
	}

	if debug {
		err = tf5server.Serve("registry.terraform.io/hashicorp/corner", muxServer.ProviderServer, tf5server.WithManagedDebug())
	} else {
		err = tf5server.Serve("registry.terraform.io/hashicorp/corner", muxServer.ProviderServer)
	}

	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
