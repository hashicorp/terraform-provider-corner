// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamic6provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

func Test_Dynamic_UsageInFunctions(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that attributes marked as DynamicPseudoType can still be utilized with built-in Terraform functions if the eventual types match up
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {}
				
				output "cidrhost" {
					value = cidrhost(corner_dynamic_thing.foo.prefix, corner_dynamic_thing.foo.hostnum)
				}
				
				output "b64_decode" {
					value = base64decode(corner_dynamic_thing.foo.sensitive_obj.b64)
					sensitive = true
				}
				
				output "combined_collections" {
					value = flatten([corner_dynamic_thing.foo.list, tolist(corner_dynamic_thing.foo.set)])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("cidrhost", knownvalue.StringExact("10.12.112.16")),
					statecheck.ExpectKnownOutputValue("b64_decode", knownvalue.StringExact("hello world")),
					statecheck.ExpectKnownOutputValue("combined_collections",
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("1"),
							knownvalue.StringExact("2"),
							knownvalue.StringExact("3"),
							knownvalue.StringExact("4"),
						}),
					),
				},
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": providerserver.NewProviderServer(testprovider.Provider{
				Resources: map[string]testprovider.Resource{
					"corner_dynamic_thing": {
						SchemaResponse: &resource.SchemaResponse{
							Schema: &tfprotov6.Schema{
								Block: &tfprotov6.SchemaBlock{
									Attributes: []*tfprotov6.SchemaAttribute{
										{
											Name:     "prefix",
											Computed: true,
											Type:     tftypes.DynamicPseudoType,
										},
										{
											Name:     "hostnum",
											Computed: true,
											Type:     tftypes.DynamicPseudoType,
										},
										{
											Name:      "sensitive_obj",
											Computed:  true,
											Sensitive: true,
											Type:      tftypes.DynamicPseudoType,
										},
										{
											Name:     "list",
											Computed: true,
											Type:     tftypes.DynamicPseudoType,
										},
										{
											Name:     "set",
											Computed: true,
											Type:     tftypes.DynamicPseudoType,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"prefix":  tftypes.String,
									"hostnum": tftypes.Number,
									"sensitive_obj": tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"b64": tftypes.DynamicPseudoType,
										},
									},
									"list": tftypes.List{
										ElementType: tftypes.String,
									},
									"set": tftypes.Set{
										ElementType: tftypes.String,
									},
								},
							}, map[string]tftypes.Value{
								"prefix":  tftypes.NewValue(tftypes.String, "10.12.112.0/20"),
								"hostnum": tftypes.NewValue(tftypes.Number, 16),
								"sensitive_obj": tftypes.NewValue(tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"b64": tftypes.String,
									},
								}, map[string]tftypes.Value{
									// "hello world"
									"b64": tftypes.NewValue(tftypes.String, "aGVsbG8gd29ybGQ="),
								}),
								"list": tftypes.NewValue(tftypes.List{
									ElementType: tftypes.String,
								}, []tftypes.Value{
									tftypes.NewValue(tftypes.String, "1"),
									tftypes.NewValue(tftypes.String, "2"),
								}),
								"set": tftypes.NewValue(tftypes.Set{
									ElementType: tftypes.String,
								}, []tftypes.Value{
									tftypes.NewValue(tftypes.String, "3"),
									tftypes.NewValue(tftypes.String, "4"),
								}),
							}),
						},
					},
				},
			}),
		},
	})
}
