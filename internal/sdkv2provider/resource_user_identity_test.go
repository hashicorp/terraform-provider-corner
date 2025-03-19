// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccResourceUserIdentity(t *testing.T) resource.TestCase {
	return resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasicIdentity,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"corner_user_identity.foo", "name", regexp.MustCompile("^For")),
				),
				PlanOnly:           true, // Should have a plan with something in there and not error
				ExpectNonEmptyPlan: true,
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
