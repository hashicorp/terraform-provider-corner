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

var _ resource.Resource = InvalidRefinementResource{}

func NewInvalidRefinement() resource.Resource {
	return &InvalidRefinementResource{}
}

type InvalidRefinementResource struct{}

func (r InvalidRefinementResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_invalid_refinement"
}

func (r InvalidRefinementResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"string_with_prefix": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.WillHavePrefix("prefix://"),
				},
			},
		},
	}
}

func (r InvalidRefinementResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InvalidRefinementResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.StringWithPrefix = types.StringValue("not correct")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r InvalidRefinementResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InvalidRefinementResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.StringWithPrefix = types.StringValue("not correct")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r InvalidRefinementResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data InvalidRefinementResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.StringWithPrefix = types.StringValue("not correct")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r InvalidRefinementResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type InvalidRefinementResourceModel struct {
	StringWithPrefix types.String `tfsdk:"string_with_prefix"`
}
