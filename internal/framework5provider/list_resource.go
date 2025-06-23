// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = ListResource{}
var _ resource.ResourceWithIdentity = ListResource{}
var _ resource.ResourceWithImportState = ListResource{}
var _ list.ListResource = ListResource{}

func NewListResource() resource.Resource {
	return &ListResource{}
}

func NewListResourceAsListResource() list.ListResource {
	return &ListResource{}
}

type ListResource struct{}

func (r ListResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r ListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list_resource"
}

func (r ListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r ListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughWithIdentity(ctx, path.Root("id"), path.Root("id"), req, resp)
}

func (r ListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ComputeInstanceResource

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue("id-123")
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

	diags = resp.Identity.Set(ctx, data.ComputeInstanceIdentity)
	resp.Diagnostics.Append(diags...)
}

func (r ListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ComputeInstanceResource

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ID.ValueString() != "id-123" {
		resp.Diagnostics.AddAttributeError(
			path.Root("id"), "Unexpected ID value", fmt.Sprintf("Expected ID to be \"id-123\", got: %s", data.ID.String()),
		)
		return
	}

	data.Name = types.StringValue("platform")
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

	diags = resp.Identity.Set(ctx, data.ComputeInstanceIdentity)
	resp.Diagnostics.Append(diags...)
}

func (r ListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ComputeInstanceResource

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r ListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r ListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{
		Attributes: map[string]listschema.Attribute{
			"filter": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r ListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data ComputeInstanceListResource

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	identities := map[string]ComputeInstanceIdentity{
		"plateau":   {ID: types.StringValue("id-001")},
		"platinum":  {ID: types.StringValue("id-002")},
		"platypus":  {ID: types.StringValue("id-003")},
		"bookworm":  {ID: types.StringValue("id-004")},
		"bookshelf": {ID: types.StringValue("id-005")},
		"bookmark":  {ID: types.StringValue("id-006")},
	}

	items := map[string]ComputeInstanceResource{
		"plateau":   {ComputeInstanceIdentity: identities["plateau"], Name: types.StringValue("plateau")},
		"platinum":  {ComputeInstanceIdentity: identities["platinum"], Name: types.StringValue("platinum")},
		"platypus":  {ComputeInstanceIdentity: identities["platypus"], Name: types.StringValue("platypus")},
		"bookworm":  {ComputeInstanceIdentity: identities["bookworm"], Name: types.StringValue("bookworm")},
		"bookshelf": {ComputeInstanceIdentity: identities["bookshelf"], Name: types.StringValue("bookshelf")},
		"bookmark":  {ComputeInstanceIdentity: identities["bookmark"], Name: types.StringValue("bookmark")},
	}

	stream.Results = func(push func(list.ListResult) bool) {
		for name, item := range items {
			if !strings.HasPrefix(name, data.Filter.ValueString()) {
				continue
			}

			result := req.NewListResult()
			result.DisplayName = item.Name.ValueString()

			if diags := result.Resource.Set(ctx, item); diags.HasError() {
				result.Diagnostics.Append(diags...)
			}

			if diags := result.Identity.Set(ctx, identities[name]); diags.HasError() {
				result.Diagnostics.Append(diags...)
			}

			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
			}

			if !push(result) {
				return
			}
		}
	}
}

type ComputeInstanceResource struct {
	ComputeInstanceIdentity
	Name types.String `tfsdk:"name"`
}

type ComputeInstanceIdentity struct {
	ID types.String `tfsdk:"id"`
}

type ComputeInstanceListResource struct {
	Filter types.String `tfsdk:"filter"`
}
