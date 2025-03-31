// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func testAccResourceUserIdentity(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
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
