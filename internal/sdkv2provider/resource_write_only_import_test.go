// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWriteOnlyImportResource(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `resource "corner_writeonly_import" "test" {
				  string_attr = "hello world!"
				  writeonly_string = "fakepassword"
				}`,
			},
			{
				ResourceName:      "corner_writeonly_import.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
