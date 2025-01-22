// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = RefinementProducerResource{}

func NewRefinementProducer() resource.Resource {
	return &RefinementProducerResource{}
}

type RefinementProducerResource struct{}

func (r RefinementProducerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_refinement_producer"
}

func (r RefinementProducerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bool_with_not_null": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.WillNotBeNull(),
				},
			},
			"int64_with_bounds": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.WillBeAtLeast(10),
					int64planmodifier.WillBeAtMost(20),
				},
			},
			"float64_with_bounds": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.WillBeAtLeast(10.234),
					float64planmodifier.WillBeAtMost(20.234),
				},
			},
			"list_with_length_bounds": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.WillHaveSizeAtLeast(2),
					listplanmodifier.WillHaveSizeAtMost(5),
				},
			},
			"string_with_prefix": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.WillHavePrefix("prefix://"),
				},
			},
		},
	}
}

func (r RefinementProducerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RefinementProducerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.BoolWithNotNull = types.BoolValue(true)
	data.Int64WithBounds = types.Int64Value(15)
	data.Float64WithBounds = types.Float64Value(12.102)
	data.ListWithLengthBounds = types.ListValueMust(
		types.StringType,
		[]attr.Value{
			types.StringValue("hello"),
			types.StringValue("there"),
			types.StringValue("world!"),
		},
	)
	data.StringWithPrefix = types.StringValue("prefix://hello-world")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementProducerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RefinementProducerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.BoolWithNotNull = types.BoolValue(true)
	data.Int64WithBounds = types.Int64Value(15)
	data.Float64WithBounds = types.Float64Value(12.102)
	data.ListWithLengthBounds = types.ListValueMust(
		types.StringType,
		[]attr.Value{
			types.StringValue("hello"),
			types.StringValue("there"),
			types.StringValue("world!"),
		},
	)
	data.StringWithPrefix = types.StringValue("prefix://hello-world")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementProducerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RefinementProducerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.BoolWithNotNull = types.BoolValue(true)
	data.Int64WithBounds = types.Int64Value(15)
	data.Float64WithBounds = types.Float64Value(12.102)
	data.ListWithLengthBounds = types.ListValueMust(
		types.StringType,
		[]attr.Value{
			types.StringValue("hello"),
			types.StringValue("there"),
			types.StringValue("world!"),
		},
	)
	data.StringWithPrefix = types.StringValue("prefix://hello-world")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementProducerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type RefinementProducerResourceModel struct {
	BoolWithNotNull      types.Bool    `tfsdk:"bool_with_not_null"`
	Int64WithBounds      types.Int64   `tfsdk:"int64_with_bounds"`
	Float64WithBounds    types.Float64 `tfsdk:"float64_with_bounds"`
	ListWithLengthBounds types.List    `tfsdk:"list_with_length_bounds"`
	StringWithPrefix     types.String  `tfsdk:"string_with_prefix"`
}
