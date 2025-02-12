// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = WriteOnlyUpgradeResource{}
var _ resource.ResourceWithUpgradeState = WriteOnlyUpgradeResource{}

func NewWriteOnlyUpgradeResource(version int64) resource.Resource {
	return &WriteOnlyUpgradeResource{
		version: version,
	}
}

type WriteOnlyUpgradeResource struct {
	// version is allowing the calling code to determine what the schema version is, to allow us to use a single resource implementation
	// and simulate an upgrade.
	version int64
}

func (r WriteOnlyUpgradeResource) UpgradeState(context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {

				resp.Diagnostics.Append(resp.State.Set(
					ctx,
					WriteOnlyUpgradeResourceModel{
						StringAttr:      types.StringValue("world!"),
						WriteOnlyString: types.StringValue("this shouldn't cause an error"),
					},
				)...)
			},
		},
	}
}

func (r WriteOnlyUpgradeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonly_upgrade"
}

func (r WriteOnlyUpgradeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: r.version,
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

func (r WriteOnlyUpgradeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config WriteOnlyUpgradeResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyUpgradeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyUpgradeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyUpgradeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r WriteOnlyUpgradeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type WriteOnlyUpgradeResourceModel struct {
	StringAttr      types.String `tfsdk:"string_attr"`
	WriteOnlyString types.String `tfsdk:"writeonly_string"`
}
