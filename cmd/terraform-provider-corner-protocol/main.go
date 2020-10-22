package main

import (
	"flag"
	// provider "github.com/hashicorp/terraform-provider-corner/internal/protocolprovider"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// TODO: fill this in with the server for the protocol provider
	panic("not implemented")
}
