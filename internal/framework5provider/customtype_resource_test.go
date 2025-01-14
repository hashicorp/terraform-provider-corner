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

// Not error test and check state for the address
// Semantic equality later, probably with Austin, read documentation on it
// Need to use custom type like IPv6 or NormalizeJSON

func TestSchemaResource_CustomTypeIPv4Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_customtype" "test" {
					ipv4test_attribute = "test value"
				}`,
				ExpectError: regexp.MustCompile("Invalid IPv4 Address String Value"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv4test_attribute"), knownvalue.StringExact("test value")),
				},
			},
			{
				Config: `resource "framework_customtype" "test" {
					ipv4test_attribute = "192.0.2.146"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv4test_attribute"), knownvalue.StringExact("192.0.2.146")),
				},
			},
		},
	})
}

func TestSchemaResource_CustomTypeIPv6Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			/*			{
							Config: `resource "framework_customtype" "test" {
								ipv6test_attribute = "test value"
							}`,
							ExpectError: regexp.MustCompile("Invalid IPv6 Address String Value"),
							ConfigStateChecks: []statecheck.StateCheck{
								statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv6test_attribute"), knownvalue.StringExact("test value")),
							},
						},
						{
							Config: `resource "framework_customtype" "test" {
								ipv6test_attribute = "2001:db8:3333:4444:5555:6666:7777:8888"
							}`,
							ConfigStateChecks: []statecheck.StateCheck{
								statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv6test_attribute"), knownvalue.StringExact("2001:db8:3333:4444:5555:6666:7777:8888")),
							},
						},
						{
							Config: `resource "framework_customtype" "test" {
								ipv6test_attribute = "1050:0000:0000:0000:0005:0600:300c:326b"
							}`,
							ConfigStateChecks: []statecheck.StateCheck{
								statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv6test_attribute"), knownvalue.StringExact("1050:0000:0000:0000:0005:0600:300c:326b")),
							},
						},*/
			{
				Config: `resource "framework_customtype" "test" {
					ipv6test_attribute = "ff06:0:0:0:0:0:0:c3" 
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_customtype.test", tfjsonpath.New("ipv6test_attribute"), knownvalue.StringExact("ff06:0:0:0:0:0:0:c3")),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_customtype.test", "ipv6test_attribute", "ff06:0:0:0:0:0:0:c3"),
				),
			},
			/*{ // ff06:0:0:0:0:0:0:c3 is ff06::c3, does sematic equality work in this case?
				Config: `resource "framework_customtype" "test" {
					ipv6test_attribute = "ff06::c3"
				}`,
				ExpectError: regexp.MustCompile("Attribute 'ipv6test_attribute' expected \"ff06:0:0:0:0:0:0:c3\", got \"ff06::c3\""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_customtype.test", "ipv6test_attribute", "ff06:0:0:0:0:0:0:c3"),
				),
			},*/
		},
	})
}
