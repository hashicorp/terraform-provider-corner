// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = RefinementConsumerResource{}
var _ resource.ResourceWithConfigValidators = RefinementConsumerResource{}

func NewRefinementConsumer() resource.Resource {
	return &RefinementConsumerResource{}
}

type RefinementConsumerResource struct{}

func (r RefinementConsumerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_refinement_consumer"
}

func (r RefinementConsumerResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("conflicting_bool_one"),
			path.MatchRoot("conflicting_bool_two"),
		),
	}
}

func (r RefinementConsumerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"conflicting_bool_one": schema.BoolAttribute{
				Optional: true,
			},
			"conflicting_bool_two": schema.BoolAttribute{
				Optional: true,
			},
			"at_most_int64": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtMost(9),
				},
			},
			"at_least_float64": schema.Float64Attribute{
				Optional: true,
				Validators: []validator.Float64{
					float64validator.AtLeast(20.235),
				},
			},
			"at_least_list_size": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(6),
				},
			},
			"at_most_list_size": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"at_least_set_size": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(6),
				},
			},
			"at_most_string_length": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(8),
				},
			},
		},
	}
}

func (r RefinementConsumerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RefinementConsumerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementConsumerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RefinementConsumerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementConsumerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RefinementConsumerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r RefinementConsumerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type RefinementConsumerResourceModel struct {
	ConflictingBoolOne types.Bool    `tfsdk:"conflicting_bool_one"`
	ConflictingBoolTwo types.Bool    `tfsdk:"conflicting_bool_two"`
	AtMostInt64        types.Int64   `tfsdk:"at_most_int64"`
	AtLeastFloat64     types.Float64 `tfsdk:"at_least_float64"`
	AtLeastListSize    types.List    `tfsdk:"at_least_list_size"`
	AtMostListSize     types.List    `tfsdk:"at_most_list_size"`
	AtLeastSetSize     types.Set     `tfsdk:"at_least_set_size"`
	AtMostStringLength types.String  `tfsdk:"at_most_string_length"`
}
