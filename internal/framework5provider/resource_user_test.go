// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFrameworkResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"framework_user.foo", "name", "Ford Prefect"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "email", "ford@prefect.co"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "age", "200"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "id", "h"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "language", "en"),
				),
			},
		},
	})
}

func TestAccFrameworkResourceUser_language(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserLanguage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"framework_user.foo", "name", "J Doe"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "email", "jdoe@example.com"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "age", "18"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "id", "jdoe"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "language", "es"),
				),
			},
		},
	})
}

func TestAccFrameworkResourceUser_interpolateLanguage(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "hashicorp/random",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"framework_user.foo", "name", "Ford Prefect"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "email", "ford@prefect.co"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "age", "200"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "id", "h"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "language", "en"),
				),
			},
			{
				Config: configResourceUserLanguageInterpolated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"framework_user.foo", "name", "Ford Prefect"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "email", "ford@prefect.co"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "age", "200"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "id", "h"),
					resource.TestCheckResourceAttr(
						"framework_user.foo", "language", "es"),
				),
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-plugin-sdk/issues/935
func TestAccFrameworkResourceUser_TF_VAR_Environment_Variable(t *testing.T) {
	expectedUserName := "Ford Prefect"
	t.Setenv("TF_VAR_framework_user_name", expectedUserName)

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
# Sourced via TF_VAR_framework_user_name environment variable
variable "framework_user_name" {
  type = string
}

resource "framework_user" "test" {
  email = "ford@prefect.co"
  name  = var.framework_user_name
  age   = 200
  id    = "h"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_user.test", "name", expectedUserName),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "framework_user" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
  id = "h"
}
`

const configResourceUserLanguage = `
resource "framework_user" "foo" {
  email = "jdoe@example.com"
  name = "J Doe"
  age = 18
  id = "jdoe"
  language = "es"
}
`

const configResourceUserLanguageInterpolated = `
resource "random_shuffle" "foo" {
  input = ["es", "es"]
  result_count = 1
}

resource "framework_user" "foo" {
  email = "ford@prefect.co"
  name = "Ford Prefect"
  age = 200
  id = "h"
  language = random_shuffle.foo.result[0]
}
`
