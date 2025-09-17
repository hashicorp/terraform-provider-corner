// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ action.Action = &ModifyFileAction{}
var _ action.ActionWithModifyPlan = &ModifyFileAction{}

func NewModifyFileAction() action.Action {
	return &ModifyFileAction{}
}

type ModifyFileAction struct{}

func (u *ModifyFileAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filename": schema.StringAttribute{
				Required: true,
			},
			"content": schema.StringAttribute{
				Required: true,
			},
			"plan_error": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

func (u *ModifyFileAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_modify_file_action"
}

func (u *ModifyFileAction) ModifyPlan(ctx context.Context, req action.ModifyPlanRequest, resp *action.ModifyPlanResponse) {
	if req.Config.Raw.IsNull() {
		return
	}

	var planError *bool

	diags := req.Config.GetAttribute(ctx, path.Root("plan_error"), &planError)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if planError != nil && *planError {
		resp.Diagnostics.AddError("ModifyPlan error", "plan_error attribute was set to true")
		return
	}
}

func (u *ModifyFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	resp.SendProgress = func(event action.InvokeProgressEvent) {
		event.Message = "starting provider defined action"
	}

	var filename string
	var content string

	diags := req.Config.GetAttribute(ctx, path.Root("filename"), &filename)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.Config.GetAttribute(ctx, path.Root("content"), &content)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fi, err := os.Create(filename)
	if err != nil {
		resp.Diagnostics.AddError("Error creating file", fmt.Sprintf("There was an error creating the file %s: "+
			"original error %s", filename, err))
		return
	}

	_, err = fi.Write([]byte(content))
	if err != nil {
		resp.Diagnostics.AddError("Error writing to file", fmt.Sprintf("There was an error writing to file %s: "+
			"original error %s", filename, err))
		return
	}

}
