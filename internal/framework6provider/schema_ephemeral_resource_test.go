// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/echoprovider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

var echoDataPath = tfjsonpath.New("data")

func TestSchemaEphemeralResource_basic(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_BoolAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					bool_attribute = true
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_DynamicAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					dynamic_attribute = {
						"attribute_one": "value1",
						"attribute_two": false,
						"attribute_three": 1234.5,
						"attribute_four": [true, 1234.5],
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"),
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
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_Float32Attribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					float32_attribute = 1234.5
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Float32Exact(1234.5)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_Float64Attribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					float64_attribute = 1234.5
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Float64Exact(1234.5)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_Int32Attribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					int32_attribute = 1234
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Int32Exact(1234)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_Int64Attribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					int64_attribute = 1234
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Int64Exact(1234)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_ListAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					list_attribute = ["value1"]
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_ListNestedAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					list_nested_attribute = [
						{
							list_nested_attribute_attribute = "value1"
						},
					]
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_attribute_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_ListNestedBlock(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					list_nested_block {
						list_nested_block_attribute = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"),
						knownvalue.ListExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"list_nested_block_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_MapAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					map_attribute = {
						key1 = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"key1": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_MapNestedAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					map_nested_attribute = {
						"key1" = {
							map_nested_attribute_attribute = "value1"
						},
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"),
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
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_NumberAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					number_attribute = 1234.5
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.NumberExact(big.NewFloat(1234.5))),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_ObjectAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					object_attribute = {
						object_attribute_attribute = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"object_attribute_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_ObjectAttributeWithDynamic(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					object_attribute_with_dynamic = {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"),
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
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SetAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					set_attribute = ["value1"]
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SetNestedAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					set_nested_attribute = [
						{
							set_nested_attribute_attribute = "value1"
						},
					]
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_attribute_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SetNestedBlock(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					set_nested_block {
						set_nested_block_attribute = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"),
						knownvalue.SetExact(
							[]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"set_nested_block_attribute": knownvalue.StringExact("value1"),
								}),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SingleNestedAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					single_nested_attribute = {
						single_nested_attribute_attribute = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_attribute_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SingleNestedAttributeWithDynamic(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					single_nested_attribute_with_dynamic = {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"),
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
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SingleNestedBlock(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					single_nested_block {
						single_nested_block_attribute = "value1"
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"),
						knownvalue.ObjectExact(
							map[string]knownvalue.Check{
								"single_nested_block_attribute": knownvalue.StringExact("value1"),
							},
						),
					),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_SingleNestedBlockWithDynamic(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					single_nested_block_with_dynamic {
						dynamic_attribute = {
							"attribute_one": "value1",
							"attribute_two": false,
							"attribute_three": 1234.5,
							"attribute_four": [true, 1234.5],
						}
					}
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"),
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
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestSchemaEphemeralResource_StringAttribute(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		// Ephemeral resources are only available in 1.10 and later
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
			"echo":      echoprovider.NewProviderServer(),
		},
		Steps: []resource.TestStep{
			{
				Config: addEchoToEphemeralSchemaConfig(`ephemeral "framework_schema" "test" {
					string_attribute = "value1"
				}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("string_attribute"), knownvalue.StringExact("value1")),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("bool_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("dynamic_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("float64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int32_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("int64_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("list_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("map_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("number_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("object_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("set_nested_block"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_attribute_with_dynamic"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block"), knownvalue.Null()),
					statecheck.ExpectKnownValue("echo.schema_test", echoDataPath.AtMapKey("single_nested_block_with_dynamic"), knownvalue.Null()),
				},
			},
		},
	})
}

// Adds the test echo provider to enable using state checks with ephemeral resources
func addEchoToEphemeralSchemaConfig(cfg string) string {
	return fmt.Sprintf(`
	%s
	provider "echo" {
		data = ephemeral.framework_schema.test
	}
	resource "echo" "schema_test" {}
	`, cfg)
}
