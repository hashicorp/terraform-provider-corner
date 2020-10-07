package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	sdkv2 "github.com/hashicorp/terraform-provider-corner/internal/sdkv2provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: sdkv2.New})
}
