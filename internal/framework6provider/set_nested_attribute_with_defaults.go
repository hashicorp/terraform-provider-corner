// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = SetNestedAttributeWithDefaultsResource{}

func NewSetNestedAttributeWithDefaultsResource() resource.Resource {
	return &SetNestedAttributeWithDefaultsResource{}
}

// SetNestedAttributeWithDefaultsResource is used for a test asserting a bug that has yet to be fixed in plugin framework
// with defaults being used in an attribute inside of a set.
//
// This bug can be observed with various different outcomes: producing duplicate set element errors, incorrect diffs during plan,
// consistent diffs with values switching back and forth, etc. Example bug reports:
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/783
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/867
//   - https://github.com/hashicorp/terraform-plugin-framework/issues/1036
type SetNestedAttributeWithDefaultsResource struct{}

func (r SetNestedAttributeWithDefaultsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_set_nested_attribute_with_defaults"
}

func (r SetNestedAttributeWithDefaultsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"set": schema.SetNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString("zero"),
						},
						"default_value": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString("this is a default"),
						},
					},
				},
			},
		},
	}
}

func (r SetNestedAttributeWithDefaultsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SetNestedAttributeWithDefaultsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetNestedAttributeWithDefaultsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SetNestedAttributeWithDefaultsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetNestedAttributeWithDefaultsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SetNestedAttributeWithDefaultsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetNestedAttributeWithDefaultsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type SetNestedAttributeWithDefaultsResourceModel struct {
	Set types.Set `tfsdk:"set"`
}
