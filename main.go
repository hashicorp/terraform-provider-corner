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
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	debugEnvFilePath := flag.String("debug-env-file", "", "Path to the debug environment file to which reattach config gets written.")
	flag.Parse()

	ctx := context.Background()
	providers := []func() tfprotov5.ProviderServer{
		protocol.Server,
		sdkv2.New().GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatalf("unable to create provider: %s", err)
	}

	var serveOpts []tf5server.ServeOpt

	if *debugFlag {
		serveOpts = append(
			serveOpts,
			tf5server.WithManagedDebug(),
		)
	}

	if *debugEnvFilePath != "" {
		if !*debugFlag {
			log.Fatalf("debug environment file path provided without debug flag, please also set -debug")
		}

		serveOpts = append(
			serveOpts,
			tf5server.WithManagedDebugEnvFilePath(*debugEnvFilePath),
		)
	}

	err = tf5server.Serve("registry.terraform.io/hashicorp/corner",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
