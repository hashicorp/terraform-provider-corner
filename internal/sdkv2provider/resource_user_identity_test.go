// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccResourceUserIdentity(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		// Latest alpha version that this JSON data is available in
		// https://github.com/hashicorp/terraform/releases/tag/v1.12.0-alpha20250319
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.12.0-alpha20250319"))),
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasicIdentity,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity("corner_user_identity.foo", map[string]knownvalue.Check{
						"email": knownvalue.StringExact("ford@prefect.co"),
					}),
				},
			},
		},
	}
}

const configResourceBasicIdentity = `
resource "corner_user_identity" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
}
`
