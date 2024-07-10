// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestSchemaResource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_BoolAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					bool_attribute = true
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					bool_attribute = false
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_DynamicAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					dynamic_attribute = "value1"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.StringExact("value1")),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					dynamic_attribute = tolist(["value1", "value2"])
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
								knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					dynamic_attribute = {
						"attribute_one": "value1",
						"attribute_two": false,
						"attribute_three": 1234.5,
						"attribute_four": [true, 1234.5],
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"attribute_one":   knownvalue.StringExact("value1"),
								"attribute_two":   knownvalue.Bool(false),
								"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
								"attribute_four": knownvalue.TupleExact(
									[]knownvalue.Check{
										knownvalue.Bool(true),
										knownvalue.NumberExact(big.NewFloat(1234.5)),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_Float32Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					float32_attribute = 1234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Float32Exact(1234.5)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					float32_attribute = 2234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Float32Exact(2234.5)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_Float64Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					float64_attribute = 1234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Float64Exact(1234.5)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					float64_attribute = 2234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Float64Exact(2234.5)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_Int32Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					int32_attribute = 1234
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Int32Exact(1234)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					int32_attribute = 2345
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Int32Exact(2345)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_Int64Attribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					int64_attribute = 1234
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Int64Exact(1234)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					int64_attribute = 2345
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Int64Exact(2345)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_ListAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					list_attribute = ["value1"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					list_attribute = ["value2"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_ListNestedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					list_nested_attribute = [
						{
							list_nested_attribute_attribute = "value1"
						},
					]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_attribute_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					list_nested_attribute = [
						{
							list_nested_attribute_attribute = "value2"
						},
					]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_attribute_attribute": knownvalue.StringExact("value2"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_ListNestedBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					list_nested_block {
						list_nested_block_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_block_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					list_nested_block {
						list_nested_block_attribute = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_block_attribute": knownvalue.StringExact("value2"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_MapAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					map_attribute = {
						key1 = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key1": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					map_attribute = {
						key1 = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key1": knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_MapNestedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					map_nested_attribute = {
						"key1" = {
							map_nested_attribute_attribute = "value1"
						},
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key1": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"map_nested_attribute_attribute": knownvalue.StringExact("value1"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					map_nested_attribute = {
						"key1" = {
							map_nested_attribute_attribute = "value2"
						},
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key1": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"map_nested_attribute_attribute": knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_NumberAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					number_attribute = 1234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.NumberExact(big.NewFloat(1234.5))),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					number_attribute = 2234.5
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.NumberExact(big.NewFloat(2234.5))),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_ObjectAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					object_attribute = {
						object_attribute_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"object_attribute_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					object_attribute = {
						object_attribute_attribute = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"object_attribute_attribute": knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_ObjectAttributeWithDynamic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = tolist(["value1", "value2"])
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ListExact(
									[]knownvalue.Check{
										knownvalue.StringExact("value1"),
										knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"attribute_one":   knownvalue.StringExact("value1"),
										"attribute_two":   knownvalue.Bool(false),
										"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"attribute_four": knownvalue.TupleExact(
											[]knownvalue.Check{
												knownvalue.Bool(true),
												knownvalue.NumberExact(big.NewFloat(1234.5)),
											},
										),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SetAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					set_attribute = ["value1"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					set_attribute = ["value2"]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SetNestedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					set_nested_attribute = [
						{
							set_nested_attribute_attribute = "value1"
						},
					]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_attribute_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					set_nested_attribute = [
						{
							set_nested_attribute_attribute = "value2"
						},
					]
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_attribute_attribute": knownvalue.StringExact("value2"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SetNestedBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					set_nested_block {
						set_nested_block_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_block_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					set_nested_block {
						set_nested_block_attribute = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_block_attribute": knownvalue.StringExact("value2"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SingleNestedAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute = {
						single_nested_attribute_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_attribute_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute = {
						single_nested_attribute_attribute = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_attribute_attribute": knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SingleNestedAttributeWithDynamic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute_with_dynamic = {
						dynamic_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute_with_dynamic = {
						dynamic_attribute = tolist(["value1", "value2"])
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ListExact(
									[]knownvalue.Check{
										knownvalue.StringExact("value1"),
										knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute_with_dynamic = {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"attribute_one":   knownvalue.StringExact("value1"),
										"attribute_two":   knownvalue.Bool(false),
										"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"attribute_four": knownvalue.TupleExact(
											[]knownvalue.Check{
												knownvalue.Bool(true),
												knownvalue.NumberExact(big.NewFloat(1234.5)),
											},
										),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SingleNestedBlock(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block {
						single_nested_block_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_block_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block {
						single_nested_block_attribute = "value2"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_block_attribute": knownvalue.StringExact("value2"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_SingleNestedBlockWithDynamic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = "value1"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = tolist(["value1", "value2"])
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ListExact(
									[]knownvalue.Check{
										knownvalue.StringExact("value1"),
										knownvalue.StringExact("value2"),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"dynamic_attribute": knownvalue.ObjectExact(
									map[string]knownvalue.Check{
										"attribute_one":   knownvalue.StringExact("value1"),
										"attribute_two":   knownvalue.Bool(false),
										"attribute_three": knownvalue.NumberExact(big.NewFloat(1234.5)),
										"attribute_four": knownvalue.TupleExact(
											[]knownvalue.Check{
												knownvalue.Bool(true),
												knownvalue.NumberExact(big.NewFloat(1234.5)),
											},
										),
									},
								),
							},
						),
					),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaResource_StringAttribute(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {
					string_attribute = "value1"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.StringExact("value1")),
				},
			},
			{
				Config: `resource "framework_schema" "test" {
					string_attribute = "value2"
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("framework_schema.test", tfjsonpath.New("string_attribute"), knownvalue.StringExact("value2")),
				},
			},
		},
	})
}
