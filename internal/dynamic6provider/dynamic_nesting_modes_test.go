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

func Test_Dynamic_Block_NestingModeList_Bug(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema (similar to nesting mode `Set`), but currently if a block with nesting mode `List`
				// contains a DynamicPseudoType, Terraform core will send the entire block as a DynamicPseudoType, resulting in an error in `terraform-plugin-go`.
				//
				// Related issues:
				// - https://github.com/hashicorp/terraform-plugin-go/issues/267
				// - https://github.com/hashicorp/terraform/issues/34574
				//
				// In this test, the dynamic `bar` attribute will cause the entire block to be sent over as a DynamicPseudoType, instead of just the bar attribute.
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
											Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
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

func Test_Dynamic_Attribute_NestingModeList_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema (similar to nesting mode `Set`), it currently behaves as expected
				// but may not be exposed to provider developers to avoid confusion.
				//
				// Related issue: https://github.com/hashicorp/terraform/issues/34574
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = [{
						bar = "hello"
						foo = 4
					}]
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
									Attributes: []*tfprotov6.SchemaAttribute{
										{
											Name:     "attribute_with_dpt",
											Required: true,
											NestedType: &tfprotov6.SchemaObject{
												Nesting: tfprotov6.SchemaObjectNestingModeList,
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

func Test_Dynamic_Block_NestingModeMap_Bug(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema (similar to nesting mode `Set`), but currently if a block with nesting mode `Map`
				// contains a DynamicPseudoType, Terraform core will send the entire block as a DynamicPseudoType, resulting in an error in `terraform-plugin-go`.
				//
				// Related issues:
				// - https://github.com/hashicorp/terraform-plugin-go/issues/267
				// - https://github.com/hashicorp/terraform/issues/34574
				//
				// In this test, the dynamic `bar` attribute will cause the entire block to be sent over as a DynamicPseudoType, instead of just the bar attribute.
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
											Nesting:  tfprotov6.SchemaNestedBlockNestingModeMap,
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

func Test_Dynamic_Attribute_NestingModeMap_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema (similar to nesting mode `Set`), it currently behaves as expected
				// but may not be exposed to provider developers to avoid confusion.
				//
				// Related issue: https://github.com/hashicorp/terraform/issues/34574
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = {
						"test1" = {
							bar = "hello"
							foo = 4
						}
						"test2" = {
							bar = "world"
							foo = 6
						}
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
									Attributes: []*tfprotov6.SchemaAttribute{
										{
											Name:     "attribute_with_dpt",
											Required: true,
											NestedType: &tfprotov6.SchemaObject{
												Nesting: tfprotov6.SchemaObjectNestingModeMap,
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

func Test_Dynamic_Block_NestingModeSet_Invalid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// Blocks with a nesting mode of `Set` are considered invalid by Terraform Core when containing a DynamicPseudoType.
				// https://github.com/hashicorp/terraform/blob/a9b43f332ea2b8fcf152a74a60af1d3a4a26e5f7/internal/configs/configschema/internal_validate.go#L73-L81
				Config: `resource "corner_dynamic_thing" "foo" {
					block_with_dpt {
						bar = "hello"
						foo = 4
					}
				}`,
				ExpectError: regexp.MustCompile(`NestingSet blocks may not contain attributes of cty.DynamicPseudoType`),
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
											Nesting:  tfprotov6.SchemaNestedBlockNestingModeSet,
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

func Test_Dynamic_Attribute_NestingModeSet_Invalid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// Attributes with a nesting mode of `Set` are considered invalid by Terraform Core when containing a DynamicPseudoType.
				// https://github.com/hashicorp/terraform/blob/a9b43f332ea2b8fcf152a74a60af1d3a4a26e5f7/internal/configs/configschema/internal_validate.go#L140-L148
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = [{
						bar = "hello"
						foo = 4
					}]
				}`,
				ExpectError: regexp.MustCompile(`NestingSet blocks may not contain attributes of cty.DynamicPseudoType`),
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
											Name:     "attribute_with_dpt",
											Required: true,
											NestedType: &tfprotov6.SchemaObject{
												Nesting: tfprotov6.SchemaObjectNestingModeSet,
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

func Test_Dynamic_Block_NestingModeSingle_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
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
											Nesting:  tfprotov6.SchemaNestedBlockNestingModeSingle,
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

func Test_Dynamic_Attribute_NestingModeSingle_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = {
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
									Attributes: []*tfprotov6.SchemaAttribute{
										{
											Name:     "attribute_with_dpt",
											Required: true,
											NestedType: &tfprotov6.SchemaObject{
												Nesting: tfprotov6.SchemaObjectNestingModeSingle,
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
