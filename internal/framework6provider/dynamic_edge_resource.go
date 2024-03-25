// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = DynamicEdgeResource{}

func NewDynamicEdgeResource() resource.Resource {
	return &DynamicEdgeResource{}
}

// DynamicEdgeResource is for testing specific scenarios for dynamic schema types.
type DynamicEdgeResource struct{}

func (r DynamicEdgeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dynamic_edge"
}

func (r DynamicEdgeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"required_dynamic": schema.DynamicAttribute{
				Required: true,
			},
			// This computed dynamic attribute changes type during refresh
			"computed_dynamic_type_changes": schema.DynamicAttribute{
				Computed: true,
			},
			// id attribute is required for acceptance testing.
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r DynamicEdgeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DynamicEdgeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Created as a boolean type
	data.ComputedDynamicTypeChanges = types.DynamicValue(types.BoolValue(true))

	data.Id = types.StringValue("test")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicEdgeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DynamicEdgeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicEdgeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DynamicEdgeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Updated to a number type
	data.ComputedDynamicTypeChanges = types.DynamicValue(types.Int64Value(200))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicEdgeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type DynamicEdgeResourceModel struct {
	RequiredDynamic            types.Dynamic `tfsdk:"required_dynamic"`
	ComputedDynamicTypeChanges types.Dynamic `tfsdk:"computed_dynamic_type_changes"`
	Id                         types.String  `tfsdk:"id"`
}
