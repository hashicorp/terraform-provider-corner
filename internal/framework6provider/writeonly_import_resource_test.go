// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWriteOnlyImportResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		// Write-only attributes are only available in 1.11.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_11_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_writeonly_import" "test" {
					string_attr = "hello world!"
					writeonly_string = "fakepassword"
				}`,
			},
			{
				ResourceName:      "framework_writeonly_import.test",
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(`Error: Import returned a non-null value for a write-only attribute`),
			},
		},
	})
}
