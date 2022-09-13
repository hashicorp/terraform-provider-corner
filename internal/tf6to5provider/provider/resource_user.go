package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

var (
	_ resource.Resource                = &resourceUser{}
	_ resource.ResourceWithConfigure   = &resourceUser{}
	_ resource.ResourceWithImportState = &resourceUser{}
)

func NewUserResource() resource.Resource {
	return &resourceUser{}
}

type resourceUser struct {
	client *backend.Client
}

func (r *resourceUser) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *resourceUser) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"email": {
				Type:          types.StringType,
				Required:      true,
				PlanModifiers: []tfsdk.AttributePlanModifier{resource.RequiresReplace()},
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"age": {
				Type:     types.NumberType,
				Required: true,
			},
			// included only for compatibility with SDKv2 test framework
			"id": {
				Type:     types.StringType,
				Optional: true,
			},
			"date_joined": {
				Type:     types.StringType,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.UseStateForUnknown(),
				},
			},
			"language": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.UseStateForUnknown(),
				},
			},
		},
	}, nil
}

func (r *resourceUser) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backend.Client)
}

type user struct {
	Email      string       `tfsdk:"email"`
	Name       string       `tfsdk:"name"`
	Age        int          `tfsdk:"age"`
	Id         string       `tfsdk:"id"`
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
	if !plan.Language.Unknown {
		newUser.Language = plan.Language.Value
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
	plan.DateJoined = types.String{Value: p.DateJoined}
	plan.Language = types.String{Value: p.Language}

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
	state.DateJoined = types.String{Value: p.DateJoined}
	state.Language = types.String{Value: p.Language}

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
	if !plan.Language.Unknown {
		newUser.Language = plan.Language.Value
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
