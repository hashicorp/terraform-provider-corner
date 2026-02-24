// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"
	"testing/fstest"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccStateStore_InMem(t *testing.T) {
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
							region = "us-east-1"
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
							region = "us-west-2"
							provider "framework" {}
						}
					}
				`,
			},
		},
	})
}

func TestAccStateStore_InMem_Validate_Error(t *testing.T) {
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
							region = "not-valid"
							provider "framework" {}
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Attribute region value must be one of: \["us-east-1" "us-west-2"\], got:\s"not-valid"`),
			},
		},
	})
}

// This test generates a state file that is around 100 MB to test the state store RPC
// chunking logic (which by default is chunked at 8 MB)
func TestAccStateStore_InMem_LargeState(t *testing.T) {
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
							region = "us-east-1"
							provider "framework" {}
						}
					}

					locals {
						# Generate padding, each iteration adds 50 characters, total: 1024 * 50 = ~51KB per string
						padding = join("", [for i in range(1024) : "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"])
						large_string = "This is a test string to generate a large state file around 100 MB for testing the state store RPC chunking logic. ${local.padding}"

						# Create a large list with padded strings, using 1024 items (max for range) with ~51KB strings each (> 50MB)
						large_list = [for i in range(1024) : "${local.large_string} - item ${i}"]
					}

					# The result of this will be > 100 MB of state data (50MB x 2, as the input attribute is duplicated to the output attribute)
					resource "terraform_data" "large" {
						input = local.large_list
					}
				`,
			},
		},
	})
}
