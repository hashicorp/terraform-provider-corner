// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			"float64_attribute": schema.Float64Attribute{
				Optional: true,
			},
			// id attribute is required for acceptance testing.
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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

func (r SchemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SchemaResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue("test")

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
	BoolAttribute     types.Bool    `tfsdk:"bool_attribute"`
	Float64Attribute  types.Float64 `tfsdk:"float64_attribute"`
	Id                types.String  `tfsdk:"id"`
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
