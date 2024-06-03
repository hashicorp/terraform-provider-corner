// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = TimeTypesResource{}

func NewTimeTypesResource() resource.Resource {
	return &TimeTypesResource{}
}

// TimeTypesResource is for testing all schema types.
type TimeTypesResource struct{}

func (r TimeTypesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_timetypes"
}

func (r TimeTypesResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"go_duration": schema.StringAttribute{
				CustomType: timetypes.GoDurationType{},
				Optional:   true,
			},
			"rfc3339": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Optional:   true,
			},
		},
	}
}

func (r TimeTypesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TimeTypesResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TimeTypesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TimeTypesResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TimeTypesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TimeTypesResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TimeTypesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type TimeTypesResourceModel struct {
	GoDuration timetypes.GoDuration `tfsdk:"go_duration"`
	Rfc3339    timetypes.RFC3339    `tfsdk:"rfc3339"`
}
