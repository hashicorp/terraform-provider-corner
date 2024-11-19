// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewSchemaEphemeralResource() ephemeral.EphemeralResource {
	return &SchemaEphemeralResource{}
}

// SchemaEphemeralResource is for testing all schema types
type SchemaEphemeralResource struct{}

func (e SchemaEphemeralResource) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema"
}

func (e SchemaEphemeralResource) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bool_attribute": schema.BoolAttribute{
				Optional: true,
			},
			"dynamic_attribute": schema.DynamicAttribute{
				Optional: true,
			},
			"float32_attribute": schema.Float32Attribute{
				Optional: true,
			},
			"float64_attribute": schema.Float64Attribute{
				Optional: true,
			},
			"int32_attribute": schema.Int32Attribute{
				Optional: true,
			},
			"int64_attribute": schema.Int64Attribute{
				Optional: true,
			},
			"list_attribute": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"map_attribute": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"number_attribute": schema.NumberAttribute{
				Optional: true,
			},
			"object_attribute": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"object_attribute_attribute": types.StringType,
				},
				Optional: true,
			},
			"set_attribute": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"string_attribute": schema.StringAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			"list_nested_block": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"list_nested_block_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"set_nested_block": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"set_nested_block_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"single_nested_block": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"single_nested_block_attribute": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

type SchemaEphemeralResourceModel struct {
	BoolAttribute     types.Bool    `tfsdk:"bool_attribute"`
	DynamicAttribute  types.Dynamic `tfsdk:"dynamic_attribute"`
	Float32Attribute  types.Float32 `tfsdk:"float32_attribute"`
	Float64Attribute  types.Float64 `tfsdk:"float64_attribute"`
	Int32attribute    types.Int32   `tfsdk:"int32_attribute"`
	Int64Attribute    types.Int64   `tfsdk:"int64_attribute"`
	ListAttribute     types.List    `tfsdk:"list_attribute"`
	ListNestedBlock   types.List    `tfsdk:"list_nested_block"`
	MapAttribute      types.Map     `tfsdk:"map_attribute"`
	NumberAttribute   types.Number  `tfsdk:"number_attribute"`
	ObjectAttribute   types.Object  `tfsdk:"object_attribute"`
	SetAttribute      types.Set     `tfsdk:"set_attribute"`
	SetNestedBlock    types.Set     `tfsdk:"set_nested_block"`
	SingleNestedBlock types.Object  `tfsdk:"single_nested_block"`
	StringAttribute   types.String  `tfsdk:"string_attribute"`
}

func (e SchemaEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data SchemaEphemeralResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
