// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = IdentityResource{}
var _ resource.ResourceWithIdentity = IdentityResource{}
var _ resource.ResourceWithImportState = IdentityResource{}

func NewIdentityResource() resource.Resource {
	return &IdentityResource{}
}

type IdentityResource struct{}

// -> 1. Define an identity schema
func (r IdentityResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
			"date_created": identityschema.StringAttribute{
				OptionalForImport: true,
			},
		},
	}
}

type IdentityResourceIdentityModel struct {
	ID          types.String `tfsdk:"id"`
	DateCreated types.String `tfsdk:"date_created"`
}

func (r IdentityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity"
}

func (r IdentityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

type IdentityResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r IdentityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IdentityResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newId := types.StringValue("id-123")

	// -> 2. Store the identity
	newIdentity := IdentityResourceIdentityModel{
		ID:          newId,
		DateCreated: types.StringValue(time.Now().Format("2006-01-02")),
	}

	data.ID = newId

	resp.Diagnostics.Append(resp.Identity.Set(ctx, newIdentity)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r IdentityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IdentityResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Name = types.StringValue("john")

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r IdentityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IdentityResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r IdentityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r IdentityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// -> 3. Implement import (pass-through of identity or import ID)
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	// importhelper.ImportStatePassthrough(ctx, path.Root("id"), path.Root("id"), req, resp)

	if !req.Identity.Raw.IsNull() {
		// -> 4. OptionalForImport attributes can be populated during import or refresh, whichever makes sense for the provider.
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("date_created"), types.StringValue(time.Now().Format("2006-01-02")))...)
	}
}
