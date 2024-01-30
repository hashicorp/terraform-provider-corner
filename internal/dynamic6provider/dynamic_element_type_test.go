// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamic6provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

func Test_Dynamic_Attribute_ListType_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema. It currently behaves as expected but may not be exposed to provider developers
				// to avoid confusion. While the element type is dynamic, all elements must still have the exact same type, which Terraform will achieve
				// by either performing a conversion (like below, converting 12345 to "12345"), or throw an error if conversion is impossible, for example:
				//
				// 		Error: Incorrect attribute value type
				//
				// 		  on terraform_plugin_test.tf line 12, in resource "corner_dynamic_thing" "foo":
				// 		  12: 					attribute_with_dpt = ["hey", { number = 12345 }]
				//
				// 		Inappropriate value for attribute "attribute_with_dpt": all list elements must have the same type.
				//
				// Related issue: https://github.com/hashicorp/terraform/issues/34574
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = ["hey", 12345]
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
											Type: tftypes.List{
												ElementType: tftypes.DynamicPseudoType,
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

func Test_Dynamic_Attribute_MapType_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema. It currently behaves as expected but may not be exposed to provider developers
				// to avoid confusion. While the element type is dynamic, all elements must still have the exact same type, which Terraform will achieve
				// by either performing a conversion (like below, converting 12345 to "12345"), or throw an error if conversion is impossible, for example:
				//
				// 	Error: Incorrect attribute value type
				//
				// 		on terraform_plugin_test.tf line 12, in resource "corner_dynamic_thing" "foo":
				// 		 					attribute_with_dpt = {
				// 		 						"key1" = "hey"
				// 		 						"key2" = {
				// 		 							number = 12345
				// 		 						}
				// 		 					}
				//
				//   Inappropriate value for attribute "attribute_with_dpt": all map elements must have the same type.
				//
				// Related issue: https://github.com/hashicorp/terraform/issues/34574
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = {
						"key1" = "hey"
						"key2" = 12345
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
											Type: tftypes.Map{
												ElementType: tftypes.DynamicPseudoType,
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

func Test_Dynamic_Attribute_SetType_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				// This may eventually be considered an invalid schema. It currently behaves as expected but may not be exposed to provider developers
				// to avoid confusion. While the element type is dynamic, all elements must still have the exact same type, which Terraform will achieve
				// by either performing a conversion (like below, converting 12345 to "12345"), or throw an error if conversion is impossible, for example:
				//
				// 		Error: Incorrect attribute value type
				//
				// 		  on terraform_plugin_test.tf line 12, in resource "corner_dynamic_thing" "foo":
				// 		  	 					attribute_with_dpt = ["hey", { number = 12345 }]
				//
				// 		Inappropriate value for attribute "attribute_with_dpt": all set elements must have the same type.
				//
				// Related issue: https://github.com/hashicorp/terraform/issues/34574
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = ["hey", 12345]
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
											Type: tftypes.Set{
												ElementType: tftypes.DynamicPseudoType,
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

func Test_Dynamic_Attribute_TupleType_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					attribute_with_dpt = ["hey", { number = 12345 }, ["there", "tuple"]]
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
											Type: tftypes.Tuple{
												ElementTypes: []tftypes.Type{
													tftypes.String,
													tftypes.DynamicPseudoType,
													tftypes.DynamicPseudoType,
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
