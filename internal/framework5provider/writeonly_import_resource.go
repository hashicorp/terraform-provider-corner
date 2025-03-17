// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = WriteOnlyImportResource{}
var _ resource.ResourceWithImportState = WriteOnlyImportResource{}

func NewWriteOnlyImportResource() resource.Resource {
	return &WriteOnlyImportResource{}
}

type WriteOnlyImportResource struct{}

func (r WriteOnlyImportResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonly_import"
}

func (r WriteOnlyImportResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"string_attr": schema.StringAttribute{
				Required: true,
			},
			"writeonly_string": schema.StringAttribute{
				Optional:  true,
				WriteOnly: true,
			},
		},
	}
}

func (r WriteOnlyImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config WriteOnlyImportResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.StringAttr = types.StringValue("hello world!")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r WriteOnlyImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r WriteOnlyImportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resp.Diagnostics.Append(resp.State.Set(
		ctx,
		WriteOnlyUpgradeResourceModel{
			StringAttr:      types.StringValue("hello world!"),
			WriteOnlyString: types.StringValue("this shouldn't cause an error"),
		},
	)...)
}

type WriteOnlyImportResourceModel struct {
	StringAttr      types.String `tfsdk:"string_attr"`
	WriteOnlyString types.String `tfsdk:"writeonly_string"`
}
