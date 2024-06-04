// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestTimeTypesResource_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_timetypes" "test" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_timetypes.test", tfjsonpath.New("go_duration"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_timetypes.test", tfjsonpath.New("rfc3339"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestTimeTypesResource_GoDuration_Valid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "framework_timetypes" "test" {
  go_duration = "1h2m3s"
}
`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_timetypes.test", tfjsonpath.New("go_duration"), knownvalue.StringExact("1h2m3s")),
				},
			},
		},
	})
}

func TestTimeTypesResource_GoDuration_Invalid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "framework_timetypes" "test" {
  go_duration = "invalid"
}
`,
				ExpectError: regexp.MustCompile(`Invalid Time Duration String Value`),
			},
		},
	})
}

func TestTimeTypesResource_RFC3339_Valid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "framework_timetypes" "test" {
  rfc3339 = "2000-01-02T03:04:05Z"
}
`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_timetypes.test", tfjsonpath.New("rfc3339"), knownvalue.StringExact("2000-01-02T03:04:05Z")),
				},
			},
		},
	})
}

func TestTimeTypesResource_RFC3339_Invalid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "framework_timetypes" "test" {
  rfc3339 = "invalid"
}
`,
				ExpectError: regexp.MustCompile(`Invalid RFC3339 String Value`),
			},
		},
	})
}
