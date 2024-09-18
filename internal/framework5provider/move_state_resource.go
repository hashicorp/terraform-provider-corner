// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = MoveStateResource{}
var _ resource.ResourceWithMoveState = MoveStateResource{}

func NewMoveStateResource() resource.Resource {
	return &MoveStateResource{}
}

// MoveStateResource is for testing the MoveResourceState RPC
// https://developer.hashicorp.com/terraform/plugin/framework/resources/state-move
type MoveStateResource struct{}

func (r MoveStateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_move_state"
}

func (r MoveStateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"moved_random_string": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r MoveStateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MoveStateResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MoveStateResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MoveStateResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r MoveStateResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			SourceSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"result": schema.StringAttribute{},
				},
			},
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if !strings.HasSuffix(req.SourceProviderAddress, "hashicorp/random") || req.SourceTypeName != "random_string" {
					resp.Diagnostics.AddError(
						"Invalid Move State Request",
						fmt.Sprintf("This test can only migrate resource state from the \"random_string\" managed resource from the \"hashicorp/random\" provider:\n\n"+
							"req.SourceProviderAddress: %q\n"+
							"req.SourceTypeName: %q\n",
							req.SourceProviderAddress,
							req.SourceTypeName,
						),
					)
				}

				var oldState RandomStringResourceModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &oldState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.SetAttribute(ctx, path.Root("moved_random_string"), oldState.Result)...)
			},
		},
	}
}

// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string
type RandomStringResourceModel struct {
	Result types.String `tfsdk:"result"`
}

type MoveStateResourceModel struct {
	MovedRandomString types.String `tfsdk:"moved_random_string"`
}
