// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSchemaResource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_schema" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_schema.test", "bool_attribute", "true"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					bool_attribute = false
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_schema.test", "bool_attribute", "false"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					float64_attribute = 2234.5
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "float64_attribute", "2234.5"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckResourceAttr("framework_schema.test", "int64_attribute", "1234"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					int64_attribute = 2345
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckResourceAttr("framework_schema.test", "int64_attribute", "2345"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_attribute.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_attribute.0", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					list_attribute = ["value2"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_attribute.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_attribute.0", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_attribute.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_attribute.0.list_nested_attribute_attribute", "value1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					list_nested_attribute = [
						{
							list_nested_attribute_attribute = "value2"
						},
					]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_attribute.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_attribute.0.list_nested_attribute_attribute", "value2"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.0.list_nested_block_attribute", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					list_nested_block {
						list_nested_block_attribute = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.0.list_nested_block_attribute", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_attribute.%", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_attribute.key1", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_attribute.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					map_attribute = {
						key1 = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_attribute.%", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_attribute.key1", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_attribute.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_nested_attribute.%", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_nested_attribute.key1.map_nested_attribute_attribute", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_attribute.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					map_nested_attribute = {
						"key1" = {
							map_nested_attribute_attribute = "value2"
						},
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_nested_attribute.%", "1"),
					resource.TestCheckResourceAttr("framework_schema.test", "map_nested_attribute.key1.map_nested_attribute_attribute", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_attribute.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "number_attribute", "1234.5"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					number_attribute = 2234.5
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "number_attribute", "2234.5"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "object_attribute.object_attribute_attribute", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					object_attribute = {
						object_attribute_attribute = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "object_attribute.object_attribute_attribute", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_attribute.#", "1"),
					resource.TestCheckTypeSetElemAttr("framework_schema.test", "set_attribute.*", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					set_attribute = ["value2"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_attribute.#", "1"),
					resource.TestCheckTypeSetElemAttr("framework_schema.test", "set_attribute.*", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_attribute.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_schema.test", "set_nested_attribute.*",
						map[string]string{"set_nested_attribute_attribute": "value1"},
					),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					set_nested_attribute = [
						{
							set_nested_attribute_attribute = "value2"
						},
					]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_attribute.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_schema.test", "set_nested_attribute.*",
						map[string]string{"set_nested_attribute_attribute": "value2"},
					),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_schema.test", "set_nested_block.*",
						map[string]string{"set_nested_block_attribute": "value1"},
					),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					set_nested_block {
						set_nested_block_attribute = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_schema.test", "set_nested_block.*",
						map[string]string{"set_nested_block_attribute": "value2"},
					),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "single_nested_attribute.single_nested_attribute_attribute", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_attribute = {
						single_nested_attribute_attribute = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckResourceAttr("framework_schema.test", "single_nested_attribute.single_nested_attribute_attribute", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "single_nested_block.single_nested_block_attribute", "value1"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					single_nested_block {
						single_nested_block_attribute = "value2"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "single_nested_block.single_nested_block_attribute", "value2"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "string_attribute"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckResourceAttr("framework_schema.test", "string_attribute", "value1"),
				),
			},
			{
				Config: `resource "framework_schema" "test" {
					string_attribute = "value2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("framework_schema.test", "bool_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "float64_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "int64_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "list_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "list_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "map_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "number_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "object_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "set_nested_attribute"),
					resource.TestCheckResourceAttr("framework_schema.test", "set_nested_block.#", "0"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_attribute"),
					resource.TestCheckNoResourceAttr("framework_schema.test", "single_nested_block"),
					resource.TestCheckResourceAttr("framework_schema.test", "string_attribute", "value2"),
				),
			},
		},
	})
}
