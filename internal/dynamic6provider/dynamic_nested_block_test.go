// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamic6provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

// There is some unexpected behavior that occurs when attempting to utilize DynamicPseudoType for an attribute
// in a nested block with a block nesting mode of list or map. The original reported issue describes some of
// these observations: https://github.com/hashicorp/terraform-plugin-go/issues/267.
//
// The tests in this file recreate the bug and also includes a test which works successfully with block nesting mode of single.

func Test_Dynamic_NestingModeList_Bug(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// The dynamic `bar` attribute will cause the entire block to be sent over as a DynamicPseudoType, instead of just the bar attribute.
				// The msgpack data of the config below is received as:
				//
				// 	{
				// 		"block_with_dpt": [
				// 		  "[\"tuple\",[[\"object\",{\"bar\":\"string\",\"foo\":\"number\"}]]]",
				// 		  [
				// 			{
				// 			  "bar": "hello",
				// 			  "foo": 4
				// 			}
				// 		  ]
				// 		]
				// 	}
				//
				Config: `resource "corner_dynamic_thing" "foo" {
					block_with_dpt {
						bar = "hello"
						foo = 4
					}
				}`,
				ExpectError: regexp.MustCompile(`unexpected code=c4 decoding map length`),
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": providerserver.NewProviderServer(testprovider.Provider{
				Resources: map[string]testprovider.Resource{
					"corner_dynamic_thing": {
						SchemaResponse: &resource.SchemaResponse{
							Schema: &tfprotov6.Schema{
								Block: &tfprotov6.SchemaBlock{
									BlockTypes: []*tfprotov6.SchemaNestedBlock{
										{
											TypeName: "block_with_dpt",
											// List nesting mode
											Nesting: tfprotov6.SchemaNestedBlockNestingModeList,
											Block: &tfprotov6.SchemaBlock{
												Attributes: []*tfprotov6.SchemaAttribute{
													{
														Name:     "bar",
														Type:     tftypes.DynamicPseudoType,
														Optional: true,
													},
													{
														Name:     "foo",
														Type:     tftypes.Number,
														Optional: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}),
		},
	})
}

func Test_Dynamic_NestingModeMap_Bug(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// The dynamic `bar` attribute will cause the entire block to be sent over as a DynamicPseudoType, instead of just the bar attribute.
				// The msgpack data of the config below is received as:
				//
				// 	{
				// 		"block_with_dpt": [
				// 		  "[\"object\",{\"test\":[\"object\",{\"bar\":\"string\",\"foo\":\"number\"}]}]",
				// 		  {
				// 			"test": {
				// 			  "bar": "hello",
				// 			  "foo": 4
				// 			}
				// 		  }
				// 		]
				// 	}
				//
				Config: `resource "corner_dynamic_thing" "foo" {
					block_with_dpt "test" {
						bar = "hello"
						foo = 4
					}
				}`,
				ExpectError: regexp.MustCompile(`unexpected code=92 decoding map length`),
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": providerserver.NewProviderServer(testprovider.Provider{
				Resources: map[string]testprovider.Resource{
					"corner_dynamic_thing": {
						SchemaResponse: &resource.SchemaResponse{
							Schema: &tfprotov6.Schema{
								Block: &tfprotov6.SchemaBlock{
									BlockTypes: []*tfprotov6.SchemaNestedBlock{
										{
											TypeName: "block_with_dpt",
											// Map nesting mode
											Nesting: tfprotov6.SchemaNestedBlockNestingModeMap,
											Block: &tfprotov6.SchemaBlock{
												Attributes: []*tfprotov6.SchemaAttribute{
													{
														Name:     "bar",
														Type:     tftypes.DynamicPseudoType,
														Optional: true,
													},
													{
														Name:     "foo",
														Type:     tftypes.Number,
														Optional: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}),
		},
	})
}

func Test_Dynamic_NestingModeSingle_Success(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// The dynamic `bar` attribute exhibits the behavior that `terraform-plugin-go` currently is expecting for DynamicPseudoType,
				// where the `bar` attribute is the only DynamicPseudoType received from Terraform. The msgpack data of the config below is received as:
				//
				// 	{
				// 		"block_with_dpt": {
				// 		  "bar": [
				// 			"\"string\"",
				// 			"hello"
				// 		  ],
				// 		  "foo": 4
				// 		}
				// 	}
				//
				Config: `resource "corner_dynamic_thing" "foo" {
					block_with_dpt {
						bar = "hello"
						foo = 4
					}
				}`,
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": providerserver.NewProviderServer(testprovider.Provider{
				Resources: map[string]testprovider.Resource{
					"corner_dynamic_thing": {
						SchemaResponse: &resource.SchemaResponse{
							Schema: &tfprotov6.Schema{
								Block: &tfprotov6.SchemaBlock{
									BlockTypes: []*tfprotov6.SchemaNestedBlock{
										{
											TypeName: "block_with_dpt",
											// Single nesting mode
											Nesting: tfprotov6.SchemaNestedBlockNestingModeSingle,
											Block: &tfprotov6.SchemaBlock{
												Attributes: []*tfprotov6.SchemaAttribute{
													{
														Name:     "bar",
														Type:     tftypes.DynamicPseudoType,
														Optional: true,
													},
													{
														Name:     "foo",
														Type:     tftypes.Number,
														Optional: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}),
		},
	})
}
