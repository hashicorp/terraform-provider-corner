// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"

	protocolv6 "github.com/hashicorp/terraform-provider-corner/internal/protocolv6provider"
)

// MAINTAINER NOTE: The current main function does not include all of the available corner resource types
// as the corner provider is rarely needed to be built as a single binary. The corner CI testing suite executes
// Go test directly against the internal packages which contain provider servers that have conflicting type names
// and different provider namespaces (testing both v5 and v6 protocols). Debugging those provider servers can also
// be achieved with Go's testing tools and a debugger such as delve.
//
// In the future, if we want to adjust this provider to be built as a single binary, we will need to refactor all of the
// internal provider packages and resource type names to avoid conflicts, as well as allow the provider binary to be built
// with protocol v5 or v6 conditionally.
func main() {

	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	debugEnvFilePath := flag.String("debug-env-file", "", "Path to the debug environment file to which reattach config gets written.")
	flag.Parse()

	var serveOpts []tf6server.ServeOpt

	if *debugFlag {
		serveOpts = append(
			serveOpts,
			tf6server.WithManagedDebug(),
		)
	}

	if *debugEnvFilePath != "" {
		if !*debugFlag {
			log.Fatalf("debug environment file path provided without debug flag, please also set -debug")
		}

		serveOpts = append(
			serveOpts,
			tf6server.WithManagedDebugEnvFilePath(*debugEnvFilePath),
		)
	}

	err := tf6server.Serve("registry.terraform.io/hashicorp/corner",
		func() tfprotov6.ProviderServer {
			return protocolv6.Server(false)
		},
		serveOpts...,
	)

	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
