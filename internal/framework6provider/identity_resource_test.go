// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestIdentityResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "test" {
					name = "tom"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("framework_identity.test", map[string]knownvalue.Check{
						"id":   knownvalue.StringExact("id-123"),
						"name": knownvalue.StringExact("tom"),
					}),
					statecheck.ExpectKnownValue("framework_identity.test", tfjsonpath.New("id"), knownvalue.StringExact("id-123")),
					statecheck.ExpectKnownValue("framework_identity.test", tfjsonpath.New("name"), knownvalue.StringExact("tom")),
				},
			},
			// Typically you don't need to test all of these different import methods,
			// but this just a smoke test for passing state + identity data through.
			{
				ImportState:     true,
				ResourceName:    "framework_identity.test",
				ImportStateKind: resource.ImportCommandWithID,
			},
			{
				ImportState:     true,
				ResourceName:    "framework_identity.test",
				ImportStateKind: resource.ImportBlockWithID,
			},
			{
				ImportState:     true,
				ResourceName:    "framework_identity.test",
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
			},
		},
	})
}

func TestIdentityResource_identity_changes(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_identity" "test" {
					name = "jerry"
				}`,
				// The resource is hardcoded to refresh with the same identity, based off the name attribute during create.
				// Resources are currently not allowed to change identities at any time, so framework will return an error message
				// after the post-apply refresh.
				ExpectError: regexp.MustCompile(`Error: Unexpected Identity Change`),
			},
		},
	})
}
