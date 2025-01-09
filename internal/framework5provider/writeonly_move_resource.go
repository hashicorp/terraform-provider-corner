// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = WriteOnlyMoveResource{}
var _ resource.ResourceWithMoveState = WriteOnlyMoveResource{}

func NewWriteOnlyMoveResource() resource.Resource {
	return &WriteOnlyMoveResource{}
}

type WriteOnlyMoveResource struct{}

func (r WriteOnlyMoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonly_move"
}

func (r WriteOnlyMoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r WriteOnlyMoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config WriteOnlyMoveResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyMoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyMoveResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyMoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r WriteOnlyMoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r WriteOnlyMoveResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, WriteOnlyMoveResourceModel{
					StringAttr:      types.StringValue("hello world!"),
					WriteOnlyString: types.StringValue("this shouldn't cause an error"),
				})...)
			},
		},
	}
}

type WriteOnlyMoveResourceModel struct {
	StringAttr      types.String `tfsdk:"string_attr"`
	WriteOnlyString types.String `tfsdk:"writeonly_string"`
}
