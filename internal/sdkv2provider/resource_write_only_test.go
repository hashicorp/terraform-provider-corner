package sdkv2

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// MAINTAINER NOTE: All the write-only data in these tests are hardcoded in the resource itself to verify
// the config data is passed to the resource Create/Update functions.
func TestWriteOnlyResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "hello!"
				  writeonly_bool = true
				  writeonly_string = "fakepassword"
				  writeonly_int = 1234
				  nested_list_block {
				  	string_attr = "hello!"
				  	writeonly_string = "fakepassword"
				  }
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("hello!")),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("string_attr"),
							knownvalue.StringExact("hello!"),
						),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("opt_or_computed_string_attr"),
							knownvalue.StringExact("computed value!"),
						),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("writeonly_string"),
							knownvalue.Null(),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("hello!")),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("string_attr"),
						knownvalue.StringExact("hello!"),
					),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("opt_or_computed_string_attr"),
						knownvalue.StringExact("computed value!"),
					),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("writeonly_string"),
						knownvalue.Null(),
					),
				},
			},
			{
				Config: `resource "corner_writeonly" "test" {
				  string_attr = "world!"
				  writeonly_bool = true
				  writeonly_string = "fakepassword"
				  writeonly_int = 1234
				  nested_list_block {
				  	string_attr = "world!"
					opt_or_computed_string_attr = "config value!"
				  	writeonly_string = "fakepassword"
				  }
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("world!")),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
						plancheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("string_attr"),
							knownvalue.StringExact("world!"),
						),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("opt_or_computed_string_attr"),
							knownvalue.StringExact("config value!"),
						),
						plancheck.ExpectKnownValue(
							"corner_writeonly.test",
							tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("writeonly_string"),
							knownvalue.Null(),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("string_attr"), knownvalue.StringExact("world!")),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_bool"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_string"), knownvalue.Null()),
					statecheck.ExpectKnownValue("corner_writeonly.test", tfjsonpath.New("writeonly_int"), knownvalue.Null()),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("string_attr"),
						knownvalue.StringExact("world!"),
					),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("opt_or_computed_string_attr"),
						knownvalue.StringExact("config value!"),
					),
					statecheck.ExpectKnownValue(
						"corner_writeonly.test",
						tfjsonpath.New("nested_list_block").AtSliceIndex(0).AtMapKey("writeonly_string"),
						knownvalue.Null(),
					),
				},
			},
		},
	})
}
