// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamic6provider_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_attr"),
						knownvalue.StringExact("hello world"),
					),
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_config_attr"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("hey"),
							knownvalue.StringExact("there"),
							knownvalue.StringExact("tuple"),
						}),
					),
				},
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_list"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("it's"),
							knownvalue.StringExact("a"),
							knownvalue.StringExact("list"),
						}),
					),
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_map"),
						knownvalue.MapExact(map[string]knownvalue.Check{
							"prop1": knownvalue.Int64Exact(15),
							"prop2": knownvalue.Float64Exact(1.23),
						}),
					),
				},
			},
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_list"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("still"),
							knownvalue.StringExact("a"),
							knownvalue.StringExact("list"),
						}),
					),
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_map"),
						knownvalue.MapExact(map[string]knownvalue.Check{
							"prop1": knownvalue.Int64Exact(10),
							"prop2": knownvalue.Float64Exact(1.23),
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
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// This test fails on Terraform 1.0.x due to a plan renderer issue where the old state type is used to try and render the planned drift.
			// Since the type has been changed, and is causing the drift, it fails with an error message of:
			//
			//		dynamic_state_types_test.go:258: Step 1/2 error: Error retrieving second post-apply plan: exit status 1
			// 		Failed to marshal plan to json: error in marshalResourceDrift: failed to encode refreshed data for corner_dynamic_thing.foo as JSON: attribute "dynamic_config_attr": tuple required
			//
			// - https://github.com/hashicorp/terraform/commit/f0cf4235f9e8eafe1d13a6a6e0720f0f0bc67e7e#diff-aec0f9962c6764cbde3325886b9b81650e1007d96d7693a9dcf0df9b53f09d2e
			tfversion.SkipBelow(tfversion.Version1_1_0),
		},
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = ["change me to a list"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_config_attr"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("change me to a list"),
						}),
					),
				},
				// State will drift because the literal is always determined as a tuple type by Terraform, but the read will set state to a list type
				ExpectNonEmptyPlan: true,
				ConfigPlanChecks: r.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("corner_dynamic_thing.foo", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				// Adding a type conversion will prevent the drift
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = tolist(["change me to a list"])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_config_attr"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("change me to a list"),
						}),
					),
				},
				ConfigPlanChecks: r.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
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

func Test_Dynamic_PlanChangesType_Error(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// Provider developers typically experience Terraform's data consistency rules in relation to the "value" of an attribute, like changing
		// a planned value to a different value during apply. Terraform's data consistency rules also extend to the type of a DynamicPseudoType attribute,
		// where once a type is determined during a run, the provider is prevented from changing the type.
		//
		// This test shows an error, where the type is determined by Terraform in config ( tuple[string, string, string] ), then changed during the plan ( list[string]).
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_config_attr = ["turn", "to", "list"]
				}`,
				ExpectError: regexp.MustCompile(`planned an invalid value`),
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
						PlanChangeFunc: func(ctx context.Context, req resource.PlanChangeRequest, resp *resource.PlanChangeResponse) {
							if req.ProposedNewState.IsNull() {
								return
							}

							// Comes in as a tuple[string,string], plan as a list[string]
							resp.PlannedState = tftypes.NewValue(tftypes.Object{
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
										tftypes.NewValue(tftypes.String, "turn"),
										tftypes.NewValue(tftypes.String, "to"),
										tftypes.NewValue(tftypes.String, "list"),
									}),
							})
						},
					},
				},
			}),
		},
	})
}

func Test_Dynamic_ApplyChangesType_Error(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// Provider developers typically experience Terraform's data consistency rules in relation to the "value" of an attribute, like changing
		// a planned value to a different value during apply. Terraform's data consistency rules also extend to the type of a DynamicPseudoType attribute,
		// where once a type is determined during a run, the provider is prevented from changing the type.
		//
		// This test shows an error, where the type is determined by the provider during plan modification ( tuple[string, string, string] ),
		// then changed during the apply ( list[string] ).
		Steps: []r.TestStep{
			{
				Config:      `resource "corner_dynamic_thing" "foo" {}`,
				ExpectError: regexp.MustCompile(`.dynamic_computed_attr: wrong final value type: tuple required`),
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
									},
								},
							},
						},
						PlanChangeFunc: func(ctx context.Context, req resource.PlanChangeRequest, resp *resource.PlanChangeResponse) {
							if req.ProposedNewState.IsNull() {
								return
							}

							// Plan as a tuple
							resp.PlannedState = tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_computed_attr": tftypes.Tuple{
										ElementTypes: []tftypes.Type{
											tftypes.String,
											tftypes.String,
											tftypes.String,
										},
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_attr": tftypes.NewValue(
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
							})
						},
						CreateResponse: &resource.CreateResponse{
							// Apply as a list
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"dynamic_computed_attr": tftypes.List{
										ElementType: tftypes.String,
									},
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_attr": tftypes.NewValue(
									tftypes.List{
										ElementType: tftypes.String,
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

func Test_Dynamic_ComputedNull_ToNewType(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test shows that an underlying type of DynamicPseudoType will be valid until a type is determined.
		// In this test, the value is initially stored as null (with no determined type), then determined as a tftypes.String on refresh.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_attr"),
						knownvalue.Null(),
					),
				},
			},
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_computed_attr"),
						knownvalue.StringExact("refreshed to a string"),
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
											Name:     "dynamic_computed_attr",
											Type:     tftypes.DynamicPseudoType,
											Computed: true,
											Optional: true,
										},
									},
								},
							},
						},
						CreateResponse: &resource.CreateResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									// It will initially be created as a null DPT
									"dynamic_computed_attr": tftypes.DynamicPseudoType,
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_attr": tftypes.NewValue(tftypes.DynamicPseudoType, nil),
							}),
						},
						ReadResponse: &resource.ReadResponse{
							NewState: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									// Switch from DPT to String
									"dynamic_computed_attr": tftypes.String,
								},
							}, map[string]tftypes.Value{
								"dynamic_computed_attr": tftypes.NewValue(tftypes.String, "refreshed to a string"),
							}),
						},
					},
				},
			}),
		},
	})
}
