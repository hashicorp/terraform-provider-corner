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

var _ resource.Resource = MoveStateResourceWithIdentity{}
var _ resource.ResourceWithIdentity = MoveStateResourceWithIdentity{}
var _ resource.ResourceWithMoveState = MoveStateResourceWithIdentity{}

func NewMoveStateResourceWithIdentity() resource.Resource {
	return &MoveStateResourceWithIdentity{}
}

// MoveStateResourceWithIdentity is for testing the MoveResourceState RPC
// https://developer.hashicorp.com/terraform/plugin/framework/resources/state-move
type MoveStateResourceWithIdentity struct{}

func (r MoveStateResourceWithIdentity) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_move_state_with_identity"
}

func (r MoveStateResourceWithIdentity) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"moved_random_string": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r MoveStateResourceWithIdentity) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r MoveStateResourceWithIdentity) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MoveStateResourceWithIdentityModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResourceWithIdentity) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MoveStateResourceWithIdentityModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResourceWithIdentity) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MoveStateResourceWithIdentityModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r MoveStateResourceWithIdentity) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r MoveStateResourceWithIdentity) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			SourceSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"result": schema.StringAttribute{},
				},
			},
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				switch req.SourceProviderAddress {
				case "registry.terraform.io/hashicorp/framework": // Corner provider (testing identity moves)
					if req.SourceTypeName != "framework_identity" {
						resp.Diagnostics.AddError(
							"Invalid Move State Request",
							fmt.Sprintf("The \"framework_move_state\" resource can only be sourced from the \"random_string\" or \"framework_identity\" managed resources:\n\n"+
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

type MoveStateResourceWithIdentityModel struct {
	MovedRandomString types.String `tfsdk:"moved_random_string"`
}
