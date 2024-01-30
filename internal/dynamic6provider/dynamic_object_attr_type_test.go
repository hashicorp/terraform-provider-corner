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

func Test_Dynamic_Attribute_ObjectType_Valid(t *testing.T) {
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
											Type: tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"bar": tftypes.DynamicPseudoType,
													"foo": tftypes.Number,
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

func Test_Dynamic_Attribute_ObjectTypeInCollections_Valid(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					list_with_dpt = [{ bar = "hello", foo = 4 }, { bar = "world", foo = 5 }]
					set_with_dpt = [{ bar = "hello", foo = 4 }, { bar = "world", foo = 5 }]
					map_with_dpt = {
						"key1" = {
							bar = "hello"
							foo = 4
						}
						"key2" = {
							bar = "world"
							foo = 5
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
											Name:     "list_with_dpt",
											Required: true,
											Type: tftypes.List{
												ElementType: tftypes.Object{
													AttributeTypes: map[string]tftypes.Type{
														"bar": tftypes.DynamicPseudoType,
														"foo": tftypes.Number,
													},
												},
											},
										},
										{
											Name:     "set_with_dpt",
											Required: true,
											Type: tftypes.Set{
												ElementType: tftypes.Object{
													AttributeTypes: map[string]tftypes.Type{
														"bar": tftypes.DynamicPseudoType,
														"foo": tftypes.Number,
													},
												},
											},
										},
										{
											Name:     "map_with_dpt",
											Required: true,
											Type: tftypes.Map{
												ElementType: tftypes.Object{
													AttributeTypes: map[string]tftypes.Type{
														"bar": tftypes.DynamicPseudoType,
														"foo": tftypes.Number,
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
