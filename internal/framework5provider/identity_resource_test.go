// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestIdentityResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.12.0-beta2"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "test" {
					name = "john"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.test", map[string]knownvalue.Check{
						"id":           knownvalue.StringExact("id-123"),
						"date_created": knownvalue.StringExact("2025-04-16"),
					}),
					statecheck.ExpectKnownValue("framework_identity.test", tfjsonpath.New("id"), knownvalue.StringExact("id-123")),
				},
			},
			// -> terraform import "id-123"
			{
				ResourceName:      "framework_identity.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// -> Plannable import with ID
			//
			// 	resource "framework_identity" "test" {
			// 		name = "john"
			// 	}
			//
			// 	import {
			// 	  to = framework_identity.test
			// 	  id = "id-123"
			// 	}
			//
			{
				ResourceName:    "framework_identity.test",
				ImportStateKind: resource.ImportBlockWithID,
				ImportState:     true,
			},
			// -> Plannable import with identity
			//
			// 	resource "framework_identity" "test" {
			// 		name = "john"
			// 	}
			//
			// 	import {
			// 	  to = framework_identity.test
			// 	  identity = {
			// 	    id = "id-123"
			// 	  }
			// 	}
			//
			{
				ResourceName:    "framework_identity.test",
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
				ImportState:     true,
			},
		},
	})
}

func TestIdentityResource_config(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.12.0-beta2"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "test" {
					name = "john"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.test", map[string]knownvalue.Check{
						"id":           knownvalue.StringExact("id-123"),
						"date_created": knownvalue.StringExact("2025-04-16"),
					}),
				},
			},
			// -> Plannable import with identity
			//
			// 	resource "framework_identity" "test" {
			// 		name = "john"
			// 	}
			//
			// 	import {
			// 	  to = framework_identity.test
			// 	  identity = {
			// 	    id = "id-123"
			// 	  }
			// 	}
			//
			{
				ResourceName:    "framework_identity.test",
				ConfigFile:      config.StaticFile("testdata/test.tf"),
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
				ImportState:     true,
			},
		},
	})
}
