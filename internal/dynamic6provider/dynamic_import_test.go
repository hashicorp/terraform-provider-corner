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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

func Test_Dynamic_ImportState(t *testing.T) {
	resourceName := "corner_dynamic_thing.foo"
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("id"),
						knownvalue.StringExact("id-123456"),
					),
					statecheck.ExpectKnownValue(
						"corner_dynamic_thing.foo",
						tfjsonpath.New("obj"),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"number": knownvalue.Int64Exact(123456),
						}),
					),
				},
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     "id-123456",
				ImportState:       true,
				ImportStateVerify: true,
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
											Name:     "id",
											Computed: true,
											Type:     tftypes.DynamicPseudoType,
										},
										{
											Name:     "obj",
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
									"id":  tftypes.DynamicPseudoType,
									"obj": tftypes.DynamicPseudoType,
								},
							}, map[string]tftypes.Value{
								"id": tftypes.NewValue(tftypes.String, "id-123456"),
								"obj": tftypes.NewValue(tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"number": tftypes.Number,
									},
								}, map[string]tftypes.Value{
									"number": tftypes.NewValue(tftypes.Number, 123456),
								}),
							}),
						},
						ImportStateResponse: &resource.ImportStateResponse{
							State: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"id":  tftypes.DynamicPseudoType,
									"obj": tftypes.DynamicPseudoType,
								},
							}, map[string]tftypes.Value{
								"id": tftypes.NewValue(tftypes.String, "id-123456"),
								"obj": tftypes.NewValue(tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"number": tftypes.Number,
									},
								}, map[string]tftypes.Value{
									"number": tftypes.NewValue(tftypes.Number, 123456),
								}),
							}),
						},
					},
				},
			}),
		},
	})
}
