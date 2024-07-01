// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = Float32PrecisionResource{}

func NewFloat32PrecisionResource() resource.Resource {
	return &Float32PrecisionResource{}
}

// Float32PrecisionResource is for testing Float32/cty.Number quirks
// https://github.com/hashicorp/terraform-plugin-framework/issues/815
type Float32PrecisionResource struct{}

func (r Float32PrecisionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_float32_precision"
}

func (r Float32PrecisionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"float32_attribute": schema.Float32Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (r Float32PrecisionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Float32PrecisionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Test semantic equality by losing the precision of the initial *big.Float
	data.Float32Attribute = types.Float32Value(data.Float32Attribute.ValueFloat32())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float32PrecisionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Float32PrecisionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float32PrecisionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Float32PrecisionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Test semantic equality by losing the precision of the initial *big.Float
	data.Float32Attribute = types.Float32Value(data.Float32Attribute.ValueFloat32())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r Float32PrecisionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type Float32PrecisionResourceModel struct {
	Float32Attribute types.Float32 `tfsdk:"float32_attribute"`
}
