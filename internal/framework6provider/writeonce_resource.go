// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-corner/internal/cornertesting"
)

var _ resource.Resource = WriteOnceResource{}

func NewWriteOnceResource() resource.Resource {
	return &WriteOnceResource{}
}

type WriteOnceResource struct{}

func (r WriteOnceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonce"
}

func (r WriteOnceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"trigger_attr": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// The only place that validation can reference prior state (which is required to determine if the planned action is
			// a create, i.e, prior state is null) is during plan modification. So the plan modifier implementation is responsible
			// for applying the "write-once" validation
			"writeonce_string": schema.StringAttribute{
				Optional:  true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					cornertesting.RequiredOnCreate(),
				},
			},
		},
	}
}

func (r WriteOnceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config WriteOnceResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(assertWriteOnlyVal(ctx, req.Config, path.Root("writeonce_string"), types.StringValue("fakepassword"))...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Since all attributes are in configuration, we write it back directly to test that the write-only attributes
	// are nulled out before sending back to TF Core.
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Once created, the only operation that can occur is replacement (delete/create)
}

func (r WriteOnceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type WriteOnceResourceModel struct {
	TriggerAttr     types.String `tfsdk:"trigger_attr"`
	WriteOnlyString types.String `tfsdk:"writeonce_string"`
}
