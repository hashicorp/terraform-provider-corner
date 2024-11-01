package framework

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-corner/internal/echoprovider"
)

// Test that echos an ephemeral resource to state for testing purposes
func Test_EchoEntireEphemeralResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.10.0"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			// Target provider we want to test with an ephemeral resource
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			// Provided test "echo" provider, would either be published in registry or from Go module in `terraform-plugin-testing`
			"echo": echoprovider.NewServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				ephemeral "framework_thing" "this" {
					name = "thing-one"
				}

				provider "echo" {
					data = ephemeral.framework_thing.this
				}

				resource "echo_resource" "echo" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo_resource.echo", tfjsonpath.New("data"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"name":  knownvalue.StringExact("thing-one"),
								"token": knownvalue.StringExact("token-abc-123-do-re-mi"),
							},
						),
					),
				},
			},
		},
	})
}

// Test that echos a single attribute from an ephemeral resource to state for testing purposes
func Test_EchoSingleEphemeralAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.10.0"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			// Target provider we want to test with an ephemeral resource
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			// Provided test "echo" provider, would either be published in registry or from Go module in `terraform-plugin-testing`
			"echo": echoprovider.NewServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				ephemeral "framework_thing" "this" {
					name = "thing-one"
				}

				provider "echo" {
					data = ephemeral.framework_thing.this.token
				}

				resource "echo_resource" "echo" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo_resource.echo", tfjsonpath.New("data"), knownvalue.StringExact("token-abc-123-do-re-mi")),
				},
			},
		},
	})
}
