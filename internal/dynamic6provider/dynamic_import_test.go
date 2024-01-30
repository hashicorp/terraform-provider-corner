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

func Test_Dynamic_ImportState(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ResourceName:  "corner_dynamic_thing.foo",
				Config:        `resource "corner_dynamic_thing" "foo" {}`,
				ImportState:   true,
				ImportStateId: "id-123456",
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
						ImportStateResponse: &resource.ImportStateResponse{
							State: tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"id": tftypes.String,
									"obj": tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"number": tftypes.Number,
										},
									},
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
