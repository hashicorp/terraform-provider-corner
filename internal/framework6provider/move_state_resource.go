// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ resource.Resource = MoveStateResource{}
var _ resource.ResourceWithIdentity = MoveStateResource{}
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

func (r MoveStateResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
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
				switch req.SourceProviderAddress {
				case "registry.terraform.io/hashicorp/random": // Random provider (testing state moves)
					if req.SourceTypeName != "random_string" {
						resp.Diagnostics.AddError(
							"Invalid Move State Request",
							fmt.Sprintf("This test can only migrate resource state from the \"random_string\" managed resource from the \"hashicorp/random\" provider:\n\n"+
								"req.SourceProviderAddress: %q\n"+
								"req.SourceTypeName: %q\n",
								req.SourceProviderAddress,
								req.SourceTypeName,
							),
						)
						return
					}

					var oldState RandomStringResourceModel
					resp.Diagnostics.Append(req.SourceState.Get(ctx, &oldState)...)
					if resp.Diagnostics.HasError() {
						return
					}

					resp.Diagnostics.Append(resp.TargetState.SetAttribute(ctx, path.Root("moved_random_string"), oldState.Result)...)
				case "registry.terraform.io/hashicorp/framework": // Corner provider (testing identity moves)
					if req.SourceTypeName != "framework_identity" {
						resp.Diagnostics.AddError(
							"Invalid Move State Request",
							fmt.Sprintf("This test can only migrate resource state from the \"framework_identity\" managed resource from the \"hashicorp/framework\" provider:\n\n"+
								"req.SourceProviderAddress: %q\n"+
								"req.SourceTypeName: %q\n",
								req.SourceProviderAddress,
								req.SourceTypeName,
							),
						)
						return
					}

					oldIdentityVal, err := req.SourceIdentity.Unmarshal(
						tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"id":   tftypes.String,
								"name": tftypes.String,
							},
						},
					)
					if err != nil {
						resp.Diagnostics.AddError(
							"Unexpected Move State Error",
							fmt.Sprintf("Error decoding source identity: %s", err.Error()),
						)
						return
					}

					var sourceIdentityObj map[string]tftypes.Value
					var sourceID, sourceName string

					oldIdentityVal.As(&sourceIdentityObj)     //nolint:errcheck // This is just a quick test of grabbing raw identity data
					sourceIdentityObj["id"].As(&sourceID)     //nolint:errcheck // This is just a quick test of grabbing raw identity data
					sourceIdentityObj["name"].As(&sourceName) //nolint:errcheck // This is just a quick test of grabbing raw identity data

					resp.Diagnostics.Append(resp.TargetState.SetAttribute(ctx, path.Root("moved_random_string"), sourceName)...)
					resp.Diagnostics.Append(resp.TargetIdentity.SetAttribute(ctx, path.Root("id"), sourceID)...)
				default:
					resp.Diagnostics.AddError(
						"Invalid Move State Request",
						fmt.Sprintf("This test can only migrate resource state from hardcoded provider/resource types:\n\n"+
							"req.SourceProviderAddress: %q\n"+
							"req.SourceTypeName: %q\n",
							req.SourceProviderAddress,
							req.SourceTypeName,
						),
					)
				}
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
