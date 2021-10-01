package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

type resourceUserType struct{}

func (r resourceUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"email": {
				Type:          types.StringType,
				Required:      true,
				PlanModifiers: []tfsdk.AttributePlanModifier{tfsdk.RequiresReplace()},
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
		},
	}, nil
}

func (r resourceUserType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceUser{
		p: *(p.(*provider)),
	}, nil
}

type resourceUser struct {
	p provider
}

type user struct {
	Email string `tfsdk:"email"`
	Name  string `tfsdk:"name"`
	Age   int    `tfsdk:"age"`
	Id    string `tfsdk:"id"`
}

func (r resourceUser) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

	err := r.p.client.CreateUser(newUser)
	if err != nil {
		resp.Diagnostics.AddError("Error creating user", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state user
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	p, err := r.p.client.ReadUser(state.Email)
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", err.Error())
		return
	}

	if p == nil {
		return
	}

	state.Name = p.Name
	state.Age = p.Age

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

	err := r.p.client.UpdateUser(newUser)
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

func (r resourceUser) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
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

	err := r.p.client.DeleteUser(userToDelete)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting user", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceUser) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("name"), req, resp)
}
