package sdkv2

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	framework "github.com/hashicorp/terraform-provider-corner/internal/framework5provider"
)

func TestSchemaWriteOnly_basic(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(framework.New()),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `ephemeral "framework_schema" "test" {
					string_attribute = "test"
				}

				resource "corner_user_writeonly" "foo" {
				  email = "ford@prefect.co"
				  name = "Ford Prefect"
				  age = 200
				  password = ephemeral.framework_schema.test.string_attribute
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("password"), knownvalue.Null()),
						plancheck.ExpectUnknownValue("corner_user_writeonly.foo", tfjsonpath.New("saved_password")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("password"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("saved_password"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestSchemaWriteOnly_randomPassword(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(framework.New()),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `ephemeral "random_password" "test" {
  					length = 20
				}

				resource "corner_user_writeonly" "foo" {
				  email    = "ford@prefect.co"
				  name     = "Ford Prefect"
				  age      = 200
				  password = ephemeral.random_password.test.result
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("password"), knownvalue.Null()),
						plancheck.ExpectUnknownValue("corner_user_writeonly.foo", tfjsonpath.New("saved_password")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("password"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_user_writeonly.foo", tfjsonpath.New("saved_password"), knownvalue.NotNull()),
				},
			},
		},
	})
}
