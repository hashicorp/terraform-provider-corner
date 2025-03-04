// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestSchemaResource_CustomTypeJSONNormalizedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
					json_normalized_attribute = "test value"
				}`,
				ExpectError: regexp.MustCompile("Invalid JSON String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("json_normalized_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					json_normalized_attribute = "{\n\t\t\t\t\t\t\t\t\t\t\t\t  \"version\": 0,\n\t\t\t\t\t\t\t\t\t\t\t\t  \"block\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"attributes\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t  \"bool_attribute\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"bool\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description\": \"example bool attribute\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description_kind\": \"markdown\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"optional\": true\n\t\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("json_normalized_attribute"), knownvalue.StringExact("{\n\t\t\t\t\t\t\t\t\t\t\t\t  \"version\": 0,\n\t\t\t\t\t\t\t\t\t\t\t\t  \"block\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"attributes\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t  \"bool_attribute\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"bool\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description\": \"example bool attribute\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description_kind\": \"markdown\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"optional\": true\n\t\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeJSONExact(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
					json_exact_attribute = "test value"
				}`,
				ExpectError: regexp.MustCompile("Invalid JSON String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("json_exact_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					json_exact_attribute = "{\n\t\t\t\t\t\t\t\t\t\t\t\t  \"version\": 0,\n\t\t\t\t\t\t\t\t\t\t\t\t  \"block\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"attributes\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t  \"bool_attribute\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"bool\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description\": \"example bool attribute\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description_kind\": \"markdown\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"optional\": true\n\t\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("json_exact_attribute"), knownvalue.StringExact("{\n\t\t\t\t\t\t\t\t\t\t\t\t  \"version\": 0,\n\t\t\t\t\t\t\t\t\t\t\t\t  \"block\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"attributes\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t  \"bool_attribute\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"bool\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description\": \"example bool attribute\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"description_kind\": \"markdown\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"optional\": true\n\t\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t  }\n\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeIPv4Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
					ip_v4_attribute = "test value"
				}`,
				ExpectError: regexp.MustCompile("Invalid IPv4 Address String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v4_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					ip_v4_attribute = "192.0.2.146"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v4_attribute"), knownvalue.StringExact("192.0.2.146")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeIPv6Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
								ip_v6_attribute = "test value"
							}`,
				ExpectError: regexp.MustCompile("Invalid IPv6 Address String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v6_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
								ip_v6_attribute = "1050:0000:0000:0000:0005:0600:300c:326b"
							}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v6_attribute"), knownvalue.StringExact("1050:0000:0000:0000:0005:0600:300c:326b")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					ip_v6_attribute = "ff06::c3" 
				}`,

				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v6_attribute"), knownvalue.StringExact("ff06::c3")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeIPv4CIDRAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
					ip_v4_cidr_attribute = "test value"
				}`,
				ExpectError: regexp.MustCompile("Invalid IPv4 CIDR String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v4_cidr_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					ip_v4_cidr_attribute = "192.0.2.146/24"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v4_cidr_attribute"), knownvalue.StringExact("192.0.2.146/24")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeIPv6CIDRAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
								ip_v6_cidr_attribute = "test value"
							}`,
				ExpectError: regexp.MustCompile("Invalid IPv6 CIDR String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v6_cidr_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
								ip_v6_cidr_attribute = "1050:0000:0000:0000:0005:0600:300c:326b/64"
							}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ip_v6_cidr_attribute"), knownvalue.StringExact("1050:0000:0000:0000:0005:0600:300c:326b/64")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeTimeRFC3339Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
								time_rfc3339_attribute = "test value"
							}`,
				ExpectError: regexp.MustCompile("Invalid RFC3339 String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("time_rfc3339_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
								time_rfc3339_attribute = "1985-04-12T23:20:50.52Z"
							}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("time_rfc3339_attribute"), knownvalue.StringExact("1985-04-12T23:20:50.52Z")),
				},
			},
		},
	})
}
