// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = DynamicComputedTypeChangeResource{}

func NewDynamicComputedTypeChangeResource() resource.Resource {
	return &DynamicComputedTypeChangeResource{}
}

// DynamicComputedTypeChangeResource is for testing the ability of a computed dynamic attribute type to change on apply (update) when unknown
// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/969
type DynamicComputedTypeChangeResource struct{}

func (r DynamicComputedTypeChangeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dynamic_computed_type_change"
}

func (r DynamicComputedTypeChangeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"required_dynamic": schema.DynamicAttribute{
				Required: true,
			},
			// This computed dynamic attribute changes type during update
			"computed_dynamic_type_changes": schema.DynamicAttribute{
				Computed: true,
			},
		},
	}
}

func (r DynamicComputedTypeChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DynamicComputedTypeChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Created as a boolean type
	data.ComputedDynamicTypeChanges = types.DynamicValue(types.BoolValue(true))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicComputedTypeChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DynamicComputedTypeChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicComputedTypeChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DynamicComputedTypeChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Updated to a number type
	data.ComputedDynamicTypeChanges = types.DynamicValue(types.Int64Value(200))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicComputedTypeChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type DynamicComputedTypeChangeResourceModel struct {
	RequiredDynamic            types.Dynamic `tfsdk:"required_dynamic"`
	ComputedDynamicTypeChanges types.Dynamic `tfsdk:"computed_dynamic_type_changes"`
}
