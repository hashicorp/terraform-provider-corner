// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/echoprovider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// This test is a smoke test for the ephemeral resource lifecycle (Open, Renew, and Close).
func TestEphemeralLifecycleResource_basic(t *testing.T) {
	t.Parallel()

	spyClient := &EphemeralResourceSpyClient{}
	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(NewWithEphemeralSpy(spyClient)),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				ephemeral "framework_lifecycle" "test" {
					name = "John Doe"
				}
				resource "time_sleep" "wait_20_seconds" {
					depends_on = [ephemeral.framework_lifecycle.test]
					create_duration = "20s"
				}`,
			},
		},
		CheckDestroy: func(_ *terraform.State) error {
			// We only really care that renew was being invoked multiple times, it should always be 4 invocations (with no skew), but we'll give a little leeway here.
			if spyClient.RenewInvocations() < 3 {
				t.Errorf("Renew lifecycle handler should have been executed at least 3 times (5s renewals in 20s), but was only executed %d times", spyClient.RenewInvocations())
			}

			// Close will be invoked 6 times (due to all of the planning/refreshing of the testing framework), but we only care that it was executed once.
			if spyClient.CloseInvocations() < 1 {
				t.Errorf("Close lifecycle handler should have been executed at least once")
			}

			return nil
		},
	})
}

// This test ensures that Terraform will skip invoking an ephemeral resource when unknown values are present in configuration.
// The framework_lifecycle ephemeral resource will return a diagnostic if an unknown value is encountered in "name".
func TestEphemeralLifecycleResource_SkipWithUnknown(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "hashicorp/random",
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralLifecycleConfig(`
					resource "random_string" "str" {
						length = 12
					}

					ephemeral "framework_lifecycle" "test" {
						name = "John ${random_string.str.result}"
					}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.lifecycle_test", echoDataPath.AtMapKey("name"), knownvalue.StringRegexp(regexp.MustCompile(`^John\s.{12}$`))),
					statecheck.ExpectKnownValue("echo.lifecycle_test", echoDataPath.AtMapKey("token"), knownvalue.StringExact("fake-token-12345")),
				},
			},
		},
	})
}

// Adds the test echo provider to enable using state checks with ephemeral resources
func addEchoToEphemeralLifecycleConfig(cfg string) string {
	return fmt.Sprintf(`
	%s
	provider "echo" {
		data = ephemeral.framework_lifecycle.test
	}
	resource "echo" "lifecycle_test" {}
	`, cfg)
}
