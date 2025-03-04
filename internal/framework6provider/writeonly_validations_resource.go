// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = WriteOnlyValidationsResource{}
var _ resource.ResourceWithConfigValidators = WriteOnlyValidationsResource{}

func NewWriteOnlyValidationsResource() resource.Resource {
	return &WriteOnlyValidationsResource{}
}

type WriteOnlyValidationsResource struct{}

func (r WriteOnlyValidationsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonly_validations"
}

func (r WriteOnlyValidationsResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.PreferWriteOnlyAttribute(
			path.MatchRoot("old_password_attr"),
			path.MatchRoot("writeonly_password"),
		),
		resourcevalidator.Conflicting(
			path.MatchRoot("old_password_attr"),
			path.MatchRoot("password_version"),
		),
		resourcevalidator.ExactlyOneOf(
			path.MatchRoot("old_password_attr"),
			path.MatchRoot("writeonly_password"),
		),
		resourcevalidator.RequiredTogether(
			path.MatchRoot("password_version"),
			path.MatchRoot("writeonly_password"),
		),
	}
}

func (r WriteOnlyValidationsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"old_password_attr": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("password_version")),
				},
			},
			"password_version": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"writeonly_password": schema.StringAttribute{
				Optional:  true,
				WriteOnly: true,
			},
		},
	}
}

func (r WriteOnlyValidationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config WriteOnlyValidationsResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.WriteOnlyPassword.IsNull() {
		if config.WriteOnlyPassword.ValueString() != "newpassword" && config.WriteOnlyPassword.ValueString() != "newpassword2" {
			resp.Diagnostics.AddAttributeError(
				path.Root("writeonly_password"),
				"Unexpected writeonly_password value",
				fmt.Sprintf("expected `writeonly_password` to be `newpassword` or `newpassword2`, got: %q", config.WriteOnlyPassword.ValueString()),
			)
			return
		}
	} else {
		if config.OldPasswordAttr.ValueString() != "oldpassword" && config.OldPasswordAttr.ValueString() != "oldpassword2" {
			resp.Diagnostics.AddAttributeError(
				path.Root("old_password_attr"),
				"Unexpected old_password_attr value",
				fmt.Sprintf("expected `old_password_attr` to be `oldpassword` or `oldpassword2`, got: %q", config.OldPasswordAttr.ValueString()),
			)
			return
		}
	}

	// Since all attributes are in configuration, we write it back directly to test that the write-only attributes
	// are nulled out before sending back to TF Core.
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyValidationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyValidationsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyValidationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Once created, the only operation that can occur is replacement (delete/create)
}

func (r WriteOnlyValidationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type WriteOnlyValidationsResourceModel struct {
	OldPasswordAttr   types.String `tfsdk:"old_password_attr"`
	PasswordVersion   types.String `tfsdk:"password_version"`
	WriteOnlyPassword types.String `tfsdk:"writeonly_password"`
}
