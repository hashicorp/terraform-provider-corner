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

var _ resource.Resource = Float64PrecisionResource{}

func NewFloat64PrecisionResource() resource.Resource {
	return &Float64PrecisionResource{}
}

// Float64PrecisionResource is for testing Float64/cty.Number quirks
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
type Float64PrecisionResource struct{}

func (r Float64PrecisionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_float64_precision"
}

func (r Float64PrecisionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// id attribute is required for acceptance testing.
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"float64_attribute": schema.Float64Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (r Float64PrecisionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Float64PrecisionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Test semantic equality by losing the precision of the initial *big.Float
	data.Float64Attribute = types.Float64Value(data.Float64Attribute.ValueFloat64())
	data.Id = types.StringValue("test")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float64PrecisionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Float64PrecisionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float64PrecisionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Float64PrecisionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Test semantic equality by losing the precision of the initial *big.Float
	data.Float64Attribute = types.Float64Value(data.Float64Attribute.ValueFloat64())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float64PrecisionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type Float64PrecisionResourceModel struct {
	Id               types.String  `tfsdk:"id"`
	Float64Attribute types.Float64 `tfsdk:"float64_attribute"`
}
