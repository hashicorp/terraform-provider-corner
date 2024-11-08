// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ ephemeral.EphemeralResourceWithConfigure = &EphemeralLifecycleResource{}
	_ ephemeral.EphemeralResourceWithRenew     = &EphemeralLifecycleResource{}
	_ ephemeral.EphemeralResourceWithClose     = &EphemeralLifecycleResource{}
)

func NewEphemeralLifecycleResource() ephemeral.EphemeralResource {
	return &EphemeralLifecycleResource{}
}

// EphemeralLifecycleResource is for testing the ephemeral resource lifecycle (Open, Renew, Close)
type EphemeralLifecycleResource struct {
	spyClient *EphemeralResourceSpyClient
}

type EphemeralLifecycleResourceModel struct {
	Name  types.String `tfsdk:"name"`
	Token types.String `tfsdk:"token"`
}

func (e *EphemeralLifecycleResource) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lifecycle"
}

func (e *EphemeralLifecycleResource) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"token": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (e *EphemeralLifecycleResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	spyClient, ok := req.ProviderData.(*EphemeralResourceSpyClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			fmt.Sprintf("Expected *EphemeralResourceSpyClient, got: %T.", req.ProviderData),
		)

		return
	}

	e.spyClient = spyClient
}

func (e *EphemeralLifecycleResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data EphemeralLifecycleResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Name.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"Unknown value encountered in Open lifecycle handler",
			`The "name" attribute should never be unknown, Terraform core should skip executing the Open lifecycle handler until the value becomes known.`,
		)
		return
	}

	data.Token = types.StringValue("fake-token-12345")

	// Renew in 5 seconds
	resp.RenewAt = time.Now().Add(5 * time.Second)

	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}

func (e *EphemeralLifecycleResource) Renew(ctx context.Context, req ephemeral.RenewRequest, resp *ephemeral.RenewResponse) {
	e.spyClient.Renew()

	// Renew again in 5 seconds
	resp.RenewAt = time.Now().Add(5 * time.Second)
}

func (e *EphemeralLifecycleResource) Close(ctx context.Context, req ephemeral.CloseRequest, resp *ephemeral.CloseResponse) {
	e.spyClient.Close()
}
