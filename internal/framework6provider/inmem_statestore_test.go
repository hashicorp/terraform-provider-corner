// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"
	"testing/fstest"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccStateStore_InMem_Basic(t *testing.T) {
	// TODO: Remove this once we start running CI tests with stable TF releases
	t.Setenv("TF_ENABLE_PLUGGABLE_STATE_STORAGE", "1")

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_15_0),
			tfversion.SkipIfNotPrerelease(),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(NewWithStateStoreFS(fstest.MapFS{})),
		},
		Steps: []resource.TestStep{
			{
				StateStore: true,
				Config: `
					terraform {
						required_providers {
							framework = {
								source = "hashicorp/framework"
							}
						}
						state_store "framework_inmem" {
							provider "framework" {}
						}
					}
				`,
			},
		},
	})
}

func TestAccStateStore_InMem_VerifyLock(t *testing.T) {
	// TODO: Remove this once we start running CI tests with stable TF releases
	t.Setenv("TF_ENABLE_PLUGGABLE_STATE_STORAGE", "1")

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_15_0),
			tfversion.SkipIfNotPrerelease(),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(NewWithStateStoreFS(fstest.MapFS{})),
		},
		Steps: []resource.TestStep{
			{
				StateStore:           true,
				VerifyStateStoreLock: true,
				Config: `
					terraform {
						required_providers {
							framework = {
								source = "hashicorp/framework"
							}
						}
						state_store "framework_inmem" {
							provider "framework" {}
						}
					}
				`,
			},
		},
	})
}
