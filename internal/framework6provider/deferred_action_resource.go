// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*DeferredActionResource)(nil)
var _ resource.ResourceWithModifyPlan = (*DeferredActionResource)(nil)
var _ resource.ResourceWithImportState = (*DeferredActionResource)(nil)

func NewDeferredActionResource() resource.Resource {
	return &DeferredActionResource{
		typeName:         "_deferred_action",
		planModification: false,
	}
}

func NewDeferredActionPlanModificationResource() resource.Resource {
	return &DeferredActionResource{
		typeName:         "_deferred_action_plan_modification",
		planModification: true,
	}
}

// DeferredActionResource is for testing all schema types.
type DeferredActionResource struct {
	typeName         string
	planModification bool
}

func (r *DeferredActionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + r.typeName
	resp.ResourceBehavior.ProviderDeferred.EnablePlanModification = r.planModification
}

func (r *DeferredActionResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan DeferredActionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.ID.IsNull() && !plan.ID.IsUnknown() && plan.ID.ValueString() != "test_id" {
		resp.Diagnostics.AddError("invalid id value", "id should be test_id")
		return
	}

	if plan.ModifyPlanDeferral.ValueBool() && req.ClientCapabilities.DeferralAllowed {
		resp.Deferred = &resource.Deferred{
			Reason: resource.DeferredReasonResourceConfigUnknown,
		}
	}
}

func (r *DeferredActionResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"modify_plan_deferral": schema.BoolAttribute{
				Optional: true,
			},
			"read_deferral": schema.BoolAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *DeferredActionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DeferredActionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DeferredActionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DeferredActionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ReadDeferral.ValueBool() && req.ClientCapabilities.DeferralAllowed {
		resp.Deferred = &resource.Deferred{
			Reason: resource.DeferredReasonResourceConfigUnknown,
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DeferredActionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DeferredActionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DeferredActionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type DeferredActionResourceModel struct {
	ModifyPlanDeferral types.Bool   `tfsdk:"modify_plan_deferral"`
	ReadDeferral       types.Bool   `tfsdk:"read_deferral"`
	ID                 types.String `tfsdk:"id"`
}

func (r *DeferredActionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ClientCapabilities.DeferralAllowed && req.ID == "deferral-id" {
		resp.Deferred = &resource.Deferred{
			Reason: resource.DeferredReasonResourceConfigUnknown,
		}
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
