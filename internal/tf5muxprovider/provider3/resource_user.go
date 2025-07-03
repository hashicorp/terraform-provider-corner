// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider3

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

var (
	_ resource.Resource                = &resourceUser{}
	_ resource.ResourceWithConfigure   = &resourceUser{}
	_ resource.ResourceWithImportState = &resourceUser{}
	_ resource.ResourceWithIdentity    = &resourceUser{}
	_ list.ListResource                = &UserListResource{}
)

func (r resourceUser) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
			"name": identityschema.StringAttribute{
				OptionalForImport: true,
			},
		},
	}
}

//hard code method results

func (u UserListResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "tf5muxprovider_user1"
}

func (u UserListResource) ListResourceConfigSchema(ctx context.Context, request list.ListResourceSchemaRequest, response *list.ListResourceSchemaResponse) {
	response.Schema = listschema.Schema{
		Attributes: map[string]listschema.Attribute{
			"filter": listschema.StringAttribute{
				Required: true,
			},
		},
	}
}

type UserListResourceModel struct {
	Filter types.String `tfsdk:"filter"`
}

type UserResource struct {
	UserListResourceIdentity
	Name types.String `tfsdk:"name"`
}

type UserListResourceIdentity struct {
	ID types.String `tfsdk:"id"`
}

type UserListResource struct {
	Filter types.String `tfsdk:"filter"`
}

var identities = map[string]UserListResourceIdentity{
	"plateau":   {ID: types.StringValue("id-001")},
	"platinum":  {ID: types.StringValue("id-002")},
	"platypus":  {ID: types.StringValue("id-003")},
	"bookworm":  {ID: types.StringValue("id-004")},
	"bookshelf": {ID: types.StringValue("id-005")},
	"bookmark":  {ID: types.StringValue("id-006")},
}

var items = map[string]UserResource{
	"plateau":   {UserListResourceIdentity: identities["plateau"], Name: types.StringValue("plateau")},
	"platinum":  {UserListResourceIdentity: identities["platinum"], Name: types.StringValue("platinum")},
	"platypus":  {UserListResourceIdentity: identities["platypus"], Name: types.StringValue("platypus")},
	"bookworm":  {UserListResourceIdentity: identities["bookworm"], Name: types.StringValue("bookworm")},
	"bookshelf": {UserListResourceIdentity: identities["bookshelf"], Name: types.StringValue("bookshelf")},
	"bookmark":  {UserListResourceIdentity: identities["bookmark"], Name: types.StringValue("bookmark")},
}

func (u UserListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data UserListResourceModel

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
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

var _ list.ListResource = UserListResource{}

func NewUserResource() resource.Resource {
	return &resourceUser{}
}

type resourceUser struct {
	client *backend.Client
}

func (r *resourceUser) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user3"
}

func (r *resourceUser) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"age": schema.NumberAttribute{
				Required: true,
			},
			"date_joined": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"language": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *resourceUser) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*backend.Client)
	if !ok {
		return
	}

	r.client = client
}

type user struct {
	Email      string       `tfsdk:"email"`
	Name       string       `tfsdk:"name"`
	Age        int          `tfsdk:"age"`
	DateJoined types.String `tfsdk:"date_joined"`
	Language   types.String `tfsdk:"language"`
}

func (r resourceUser) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan user
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newUser := &backend.User{
		Email: plan.Email,
		Name:  plan.Name,
		Age:   plan.Age,
	}
	if !plan.Language.IsUnknown() {
		newUser.Language = plan.Language.ValueString()
	}

	err := r.client.CreateUser(newUser)
	if err != nil {
		resp.Diagnostics.AddError("Error creating user", err.Error())
		return
	}

	p, err := r.client.ReadUser(newUser.Email)
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", err.Error())
		return
	}

	if p == nil {
		resp.Diagnostics.AddError("Error reading user", "could not find user after it was created")
		return
	}
	plan.DateJoined = types.StringValue(p.DateJoined)
	plan.Language = types.StringValue(p.Language)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state user
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	p, err := r.client.ReadUser(state.Email)
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", err.Error())
		return
	}

	if p == nil {
		return
	}

	state.Name = p.Name
	state.Age = p.Age
	state.DateJoined = types.StringValue(p.DateJoined)
	state.Language = types.StringValue(p.Language)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan user
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newUser := &backend.User{
		Email: plan.Email,
		Name:  plan.Name,
		Age:   plan.Age,
	}
	if !plan.Language.IsUnknown() {
		newUser.Language = plan.Language.ValueString()
	}

	err := r.client.UpdateUser(newUser)
	if err != nil {
		resp.Diagnostics.AddError("Error updating user", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state user
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userToDelete := &backend.User{
		Email: state.Email,
		Name:  state.Name,
		Age:   state.Age,
	}

	err := r.client.DeleteUser(userToDelete)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting user", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceUser) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
