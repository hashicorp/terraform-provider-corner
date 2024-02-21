// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamic6provider_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

// With DynamicPseudoType as the schema for an attribute, Terraform will not perform any type conversion on literal expressions.
// - https://developer.hashicorp.com/terraform/language/expressions/types#literal-expressions
//
// For complex types, Terraform only has literal expressions for creating tuple and object values.
// - https://developer.hashicorp.com/terraform/language/expressions/type-constraints#complex-type-literals
// - https://developer.hashicorp.com/terraform/language/expressions/types#more-about-complex-types
//
// Lists, maps, and sets cannot directly be represented in Terraform config without using a type conversion function like `tolist()`.
// - https://developer.hashicorp.com/terraform/language/functions/tolist

func Test_Dynamic_ComplexTypeLiterals_Tuple(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test asserts the type that Terraform passes to a DynamicPseudoType attribute when using a literal expression for a tuple.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_attr = ["it's", "a", "tuple"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_attr"),
						// The type is stored as a tuple, but we can still use `ListExact` to verify state
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("it's"),
							knownvalue.StringExact("a"),
							knownvalue.StringExact("tuple"),
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
											Name:     "dynamic_attr",
											Type:     tftypes.DynamicPseudoType,
											Required: true,
										},
									},
								},
							},
						},
						PlanChangeFunc: verifyDynamicAttributeByPath(
							tftypes.NewAttributePath().WithAttributeName("dynamic_attr"),
							tftypes.NewValue(
								tftypes.Tuple{
									ElementTypes: []tftypes.Type{
										tftypes.String,
										tftypes.String,
										tftypes.String,
									},
								}, []tftypes.Value{
									tftypes.NewValue(tftypes.String, "it's"),
									tftypes.NewValue(tftypes.String, "a"),
									tftypes.NewValue(tftypes.String, "tuple"),
								},
							),
						),
					},
				},
			}),
		},
	})
}

func Test_Dynamic_ComplexTypeLiterals_Object(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test asserts the type that Terraform passes to a DynamicPseudoType attribute when using a literal expression for an object.
		//
		// Since object will be the type for `dynamic_attr` and the value of `prop3` is the literal `null`, which doesn't have a defined type in Terraform,
		// the resulting type for `prop3` will be DynamicPseudoType as the type has not yet been determined.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_attr = {
						prop1 = "hello"
						prop2 = "world"
						prop3 = null
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_attr"),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"prop1": knownvalue.StringExact("hello"),
							"prop2": knownvalue.StringExact("world"),
							"prop3": knownvalue.Null(),
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
											Name:     "dynamic_attr",
											Type:     tftypes.DynamicPseudoType,
											Required: true,
										},
									},
								},
							},
						},
						PlanChangeFunc: verifyDynamicAttributeByPath(
							tftypes.NewAttributePath().WithAttributeName("dynamic_attr"),
							tftypes.NewValue(
								tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"prop1": tftypes.String,
										"prop2": tftypes.String,
										// `null` literal is passed as DynamicPseudoType, since the type of `prop3` has not been determined yet
										"prop3": tftypes.DynamicPseudoType,
									},
								}, map[string]tftypes.Value{
									"prop1": tftypes.NewValue(tftypes.String, "hello"),
									"prop2": tftypes.NewValue(tftypes.String, "world"),
									"prop3": tftypes.NewValue(tftypes.DynamicPseudoType, nil),
								}),
						),
					},
				},
			}),
		},
	})
}

func Test_Dynamic_TypeConversion_Map(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test asserts the type that Terraform passes to a DynamicPseudoType attribute when using a map type conversion function on an object literal.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_attr = tomap({
						prop1 = "hello"
						prop2 = "world"
						prop3 = null
					})
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_attr"),
						knownvalue.MapExact(map[string]knownvalue.Check{
							"prop1": knownvalue.StringExact("hello"),
							"prop2": knownvalue.StringExact("world"),
							"prop3": knownvalue.Null(),
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
											Name:     "dynamic_attr",
											Type:     tftypes.DynamicPseudoType,
											Required: true,
										},
									},
								},
							},
						},
						PlanChangeFunc: verifyDynamicAttributeByPath(
							tftypes.NewAttributePath().WithAttributeName("dynamic_attr"),
							tftypes.NewValue(
								tftypes.Map{
									ElementType: tftypes.String,
								},
								map[string]tftypes.Value{
									"prop1": tftypes.NewValue(tftypes.String, "hello"),
									"prop2": tftypes.NewValue(tftypes.String, "world"),
									"prop3": tftypes.NewValue(tftypes.String, nil),
								},
							),
						),
					},
				},
			}),
		},
	})
}

func Test_Dynamic_TypeConversion_List(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test asserts the type that Terraform passes to a DynamicPseudoType attribute when using a list type conversion function on a tuple literal.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_attr = tolist(["it's", "a", "list"])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_attr"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("it's"),
							knownvalue.StringExact("a"),
							knownvalue.StringExact("list"),
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
											Name:     "dynamic_attr",
											Type:     tftypes.DynamicPseudoType,
											Required: true,
										},
									},
								},
							},
						},
						PlanChangeFunc: verifyDynamicAttributeByPath(
							tftypes.NewAttributePath().WithAttributeName("dynamic_attr"),
							tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.String,
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.String, "it's"),
									tftypes.NewValue(tftypes.String, "a"),
									tftypes.NewValue(tftypes.String, "list"),
								},
							),
						),
					},
				},
			}),
		},
	})
}

func Test_Dynamic_TypeConversion_Set(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// This test asserts the type that Terraform passes to a DynamicPseudoType attribute when using a set type conversion function on a tuple literal.
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dynamic_attr = toset([true, false])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("dynamic_attr"),
						knownvalue.SetExact([]knownvalue.Check{
							knownvalue.Bool(false),
							knownvalue.Bool(true),
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
											Name:     "dynamic_attr",
											Type:     tftypes.DynamicPseudoType,
											Required: true,
										},
									},
								},
							},
						},
						PlanChangeFunc: verifyDynamicAttributeByPath(
							tftypes.NewAttributePath().WithAttributeName("dynamic_attr"),
							tftypes.NewValue(
								tftypes.Set{
									ElementType: tftypes.Bool,
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Bool, true),
									tftypes.NewValue(tftypes.Bool, false),
								},
							),
						),
					},
				},
			}),
		},
	})
}

func verifyDynamicAttributeByPath(path *tftypes.AttributePath, expectedVal tftypes.Value) func(context.Context, resource.PlanChangeRequest, *resource.PlanChangeResponse) {
	return func(ctx context.Context, req resource.PlanChangeRequest, resp *resource.PlanChangeResponse) {
		if req.ProposedNewState.IsNull() {
			return
		}

		val, _, err := tftypes.WalkAttributePath(req.Config, path)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Verify Dynamic Attribute Failed",
				Detail:    fmt.Sprintf("error finding dynamic type path: %s", err.Error()),
				Attribute: path,
			})
		}
		tftypeVal, ok := val.(tftypes.Value)
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Verify Dynamic Attribute Failed",
				Detail:    fmt.Sprintf("error reading dynamic value, expected tftypes.Value, got %T", val),
				Attribute: path,
			})
		}

		// Verify that the types match
		if !tftypeVal.Type().Equal(expectedVal.Type()) {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Verify Dynamic Attribute Type Failed",
				Detail:    fmt.Sprintf("expected: %s, got: %s", expectedVal.Type(), tftypeVal.Type()),
				Attribute: path,
			})
		}

		// Verify that the values match
		if !tftypeVal.Equal(expectedVal) {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Verify Dynamic Attribute Value Failed",
				Detail:    fmt.Sprintf("expected: %s, got: %s", expectedVal, tftypeVal),
				Attribute: path,
			})
		}
	}
}
