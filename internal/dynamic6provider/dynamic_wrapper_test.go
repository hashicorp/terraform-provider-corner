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

func TestDynamic_Wrap_V6(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that if a provider developer sets a state value which is marked as dynamic pseudo-type in the schema to a concrete type.
		// The logic that ensures that the final value is wrapped in a dynamic pseudo-type is implemented in the `NewDynamicValue` function in `terraform-plugin-go`,
		// which uses the schema to determine how to encode the type to Terraform.
		//
		//  https://github.com/hashicorp/terraform-plugin-framework/blob/68e33ef13ddcb23d0a85797648129048cacf8da2/internal/toproto5/dynamic_value.go#L32
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_collection = ["hey", "there", "tuple"]
				}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_primitive", "hello world"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.0", "hey"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.1", "there"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.2", "tuple"),
				),
			},
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_collection = ["hey", "there", "tuple"]
				}`,
				PlanOnly: true,
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
											Name:     "dynamic_collection",
											Type:     tftypes.DynamicPseudoType,
											Optional: true,
											Computed: true,
										},
										{
											Name:     "dynamic_primitive",
											Type:     tftypes.DynamicPseudoType,
											Computed: true,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									// The schema's are dynamic, however these are all concrete types
									"dynamic_collection": tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
											tftypes.String,
											tftypes.String,
										},
									},
									"dynamic_primitive": tftypes.String,
								},
							}, map[string]tftypes.Value{
								"dynamic_collection": tftypes.NewValue(
									tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
											tftypes.String,
											tftypes.String,
										},
									}, []tftypes.Value{
										tftypes.NewValue(tftypes.String, "hey"),
										tftypes.NewValue(tftypes.String, "there"),
										tftypes.NewValue(tftypes.String, "tuple"),
									}),
								"dynamic_primitive": tftypes.NewValue(tftypes.String, "hello world"),
							}),
						},
					},
				},
			}),
		},
	})
}

func TestDynamic_ListType_Preservation_V6(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that Terraform core will preserve the type of a dynamic attribute for future Terraform operations
		//
		// https://github.com/zclconf/go-cty/blob/main/docs/json.md#type-preserving-json-serialization
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.0", "hey"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.1", "there"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_collection.2", "list"),
				),
			},
			{
				Config:   `resource "corner_dynamic_thing" "foo" {}`,
				PlanOnly: true,
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
											Name: "dynamic_collection",
											// Type:     tftypes.DynamicPseudoType,
											Type: tftypes.List{
												ElementType: tftypes.String,
											},
											Optional: true,
											Computed: true,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_collection": tftypes.List{
										ElementType: tftypes.String,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_collection": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
									},
									[]tftypes.Value{
										tftypes.NewValue(tftypes.String, "hey"),
										tftypes.NewValue(tftypes.String, "there"),
										tftypes.NewValue(tftypes.String, "list"),
									},
								),
							}),
						},
						ReadResponse: &resource.ReadResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_collection": tftypes.List{
										ElementType: tftypes.String,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_collection": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
									},
									[]tftypes.Value{
										tftypes.NewValue(tftypes.String, "hey"),
										tftypes.NewValue(tftypes.String, "there"),
										tftypes.NewValue(tftypes.String, "tuple"),
									},
								),
							}),
						},
					},
				},
			}),
		},
	})
}
