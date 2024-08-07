// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = SchemaResource{}

func NewSchemaResource() resource.Resource {
	return &SchemaResource{}
}

// SchemaResource is for testing all schema types.
type SchemaResource struct{}

func (r SchemaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema"
}

func (r SchemaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"list_nested_attribute": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"list_nested_attribute_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
			"map_attribute": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"map_nested_attribute": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"map_nested_attribute_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
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
			"object_attribute_with_dynamic": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"dynamic_attribute": types.DynamicType,
				},
				Optional: true,
			},
			"set_attribute": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"set_nested_attribute": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"set_nested_attribute_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
			"single_nested_attribute": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"single_nested_attribute_attribute": schema.StringAttribute{
						Optional: true,
					},
				},
				Optional: true,
			},
			"single_nested_attribute_with_dynamic": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"dynamic_attribute": schema.DynamicAttribute{
						Optional: true,
					},
				},
				Optional: true,
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
			"single_nested_block_with_dynamic": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"dynamic_attribute": schema.DynamicAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r SchemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SchemaResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SchemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SchemaResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SchemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SchemaResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SchemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type SchemaResourceModel struct {
	BoolAttribute                    types.Bool    `tfsdk:"bool_attribute"`
	DynamicAttribute                 types.Dynamic `tfsdk:"dynamic_attribute"`
	Float32Attribute                 types.Float32 `tfsdk:"float32_attribute"`
	Float64Attribute                 types.Float64 `tfsdk:"float64_attribute"`
	Int32Attribute                   types.Int32   `tfsdk:"int32_attribute"`
	Int64Attribute                   types.Int64   `tfsdk:"int64_attribute"`
	ListAttribute                    types.List    `tfsdk:"list_attribute"`
	ListNestedAttribute              types.List    `tfsdk:"list_nested_attribute"`
	ListNestedBlock                  types.List    `tfsdk:"list_nested_block"`
	MapAttribute                     types.Map     `tfsdk:"map_attribute"`
	MapNestedAttribute               types.Map     `tfsdk:"map_nested_attribute"`
	NumberAttribute                  types.Number  `tfsdk:"number_attribute"`
	ObjectAttribute                  types.Object  `tfsdk:"object_attribute"`
	ObjectAttributeWithDynamic       types.Object  `tfsdk:"object_attribute_with_dynamic"`
	SetAttribute                     types.Set     `tfsdk:"set_attribute"`
	SetNestedAttribute               types.Set     `tfsdk:"set_nested_attribute"`
	SetNestedBlock                   types.Set     `tfsdk:"set_nested_block"`
	SingleNestedAttribute            types.Object  `tfsdk:"single_nested_attribute"`
	SingleNestedAttributeWithDynamic types.Object  `tfsdk:"single_nested_attribute_with_dynamic"`
	SingleNestedBlock                types.Object  `tfsdk:"single_nested_block"`
	SingleNestedBlockWithDynamic     types.Object  `tfsdk:"single_nested_block_with_dynamic"`
	StringAttribute                  types.String  `tfsdk:"string_attribute"`
}
