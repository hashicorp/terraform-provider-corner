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

func Test_Dynamic_TypedValueToState(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that for a DynamicPseudoType attribute, the state value is properly set when provided a matching typed value.
		//
		// The logic that ensures that the value passed to Terraform is wrapped in a DynamicPseudoType is implemented in the `tfprotov6.NewDynamicValue`
		// function in `terraform-plugin-go`, which uses the schema to determine how to encode the DynamicPseudoType to Terraform.
		//
		//  https://github.com/hashicorp/terraform-plugin-framework/blob/68e33ef13ddcb23d0a85797648129048cacf8da2/internal/toproto6/dynamic_value.go#L32
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = ["hey", "there", "tuple"]
				}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_attr", "hello world"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_config_attr.0", "hey"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_config_attr.1", "there"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_config_attr.2", "tuple"),
				),
			},
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = ["hey", "there", "tuple"]
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
											Name:     "dynamic_computed_attr",
											Type:     tftypes.DynamicPseudoType,
											Computed: true,
										},
										{
											Name:     "dynamic_config_attr",
											Type:     tftypes.DynamicPseudoType,
											Optional: true,
											Computed: true,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							// Although the schema types are DynamicPseudoType, the values provided for NewState are all concrete types
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_computed_attr": tftypes.String,
									"dynamic_config_attr": tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
											tftypes.String,
											tftypes.String,
										},
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_attr": tftypes.NewValue(tftypes.String, "hello world"),
								"dynamic_config_attr": tftypes.NewValue(
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
							}),
						},
					},
				},
			}),
		},
	})
}

func Test_Dynamic_TypePreservedInState(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that Terraform will preserve the type of a dynamic attribute for future Terraform operations.
		//
		// For DynamicPseudoType, Terraform/go-cty JSON serialization will preserve the type in state, for this example, state would look like:
		//
		// 	{
		// 		"schema_version": 0,
		// 		"attributes": {
		// 		  "dynamic_computed_list": {
		// 			"value": [
		// 			  "it's",
		// 			  "a",
		// 			  "list"
		// 			],
		// 			"type": [
		// 			  "list",
		// 			  "string"
		// 			]
		// 		  },
		// 		  "dynamic_computed_map": {
		// 			"value": {
		// 			  "prop1": 15,
		// 			  "prop2": 1.23
		// 			},
		// 			"type": [
		// 			  "map",
		// 			  "number"
		// 			]
		// 		  }
		// 		},
		// 		"sensitive_attributes": []
		// 	}
		//
		// - https://github.com/zclconf/go-cty/blob/main/docs/json.md#type-preserving-json-serialization
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.0", "it's"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.1", "a"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.2", "list"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_map.prop1", "15"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_map.prop2", "1.23"),
				),
			},
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.0", "still"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.1", "a"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_list.2", "list"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_map.prop1", "10"),
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_computed_map.prop2", "1.23"),
				),
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
											Name:     "dynamic_computed_list",
											Type:     tftypes.DynamicPseudoType,
											Computed: true,
										},
										{
											Name:     "dynamic_computed_map",
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
									"dynamic_computed_list": tftypes.List{
										ElementType: tftypes.String,
									},
									"dynamic_computed_map": tftypes.Map{
										ElementType: tftypes.Number,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_list": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
									},
									[]tftypes.Value{
										tftypes.NewValue(tftypes.String, "it's"),
										tftypes.NewValue(tftypes.String, "a"),
										tftypes.NewValue(tftypes.String, "list"),
									},
								),
								"dynamic_computed_map": tftypes.NewValue(
									tftypes.Map{
										ElementType: tftypes.Number,
									},
									map[string]tftypes.Value{
										"prop1": tftypes.NewValue(tftypes.Number, 15),
										"prop2": tftypes.NewValue(tftypes.Number, 1.23),
									},
								),
							}),
						},
						ReadResponse: &resource.ReadResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_computed_list": tftypes.List{
										ElementType: tftypes.String,
									},
									"dynamic_computed_map": tftypes.Map{
										ElementType: tftypes.Number,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_list": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
									},
									[]tftypes.Value{
										tftypes.NewValue(tftypes.String, "still"),
										tftypes.NewValue(tftypes.String, "a"),
										tftypes.NewValue(tftypes.String, "list"),
									},
								),
								"dynamic_computed_map": tftypes.NewValue(
									tftypes.Map{
										ElementType: tftypes.Number,
									},
									map[string]tftypes.Value{
										"prop1": tftypes.NewValue(tftypes.Number, 10),
										"prop2": tftypes.NewValue(tftypes.Number, 1.23),
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

func Test_Dynamic_TypeChangesInState(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test verifies that for a DynamicPseudoType attribute, the state value type can change during refresh. While this behavior is valid,
		// it can potentially cause drift because Terraform will not use prior state type information to convert future plan values.
		//
		// i.e. If a literal tuple is defined in config, refreshing the value as a list type will not convince Terraform the literal is a list :)
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = ["change me to a list"]
				}`,
				// TODO: switch to use nice new state checks :)
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_config_attr.0", "change me to a list"),
				),
				// State will drift because the literal is always determined as a tuple type by Terraform, but the read will set state to a list type
				ExpectNonEmptyPlan: true,
				// ConfigPlanChecks: r.ConfigPlanChecks{
				// 	PostApplyPostRefresh: []plancheck.PlanCheck{
				// 		plancheck.ExpectResourceAction("corner_dynamic_thing.foo", plancheck.ResourceActionUpdate),
				// 	},
				// },
			},
			{
				// Adding a type conversion will prevent the drift
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = tolist(["change me to a list"])
				}`,
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("corner_dynamic_thing.foo", "dynamic_config_attr.0", "change me to a list"),
				),
				// ConfigPlanChecks: r.ConfigPlanChecks{
				// 	PostApplyPostRefresh: []plancheck.PlanCheck{
				// 		plancheck.ExpectEmptyPlan(),
				// 	},
				// },
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
											Name:     "dynamic_config_attr",
											Type:     tftypes.DynamicPseudoType,
											Optional: true,
											Computed: true,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							// Create as a tuple
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_config_attr": tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
										},
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_config_attr": tftypes.NewValue(
									tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
										},
									}, []tftypes.Value{
										tftypes.NewValue(tftypes.String, "change me to a list"),
									}),
							}),
						},
						// Read as a list
						ReadResponse: &resource.ReadResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_config_attr": tftypes.List{
										ElementType: tftypes.String,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_config_attr": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
									}, []tftypes.Value{
										tftypes.NewValue(tftypes.String, "change me to a list"),
									}),
							}),
						},
					},
				},
			}),
		},
	})
}
