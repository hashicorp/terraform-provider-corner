// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = DynamicSchemaResource{}

func NewDynamicSchemaResource() resource.Resource {
	return &DynamicSchemaResource{}
}

// DynamicSchemaResource is for testing the dynamic schema type.
//
// This is separated from the standard `SchemaResource` for protocol v5 because
// of a bug in Terraform v0.12.x around handling null values. See this resource's acceptance tests
// for more details.
type DynamicSchemaResource struct{}

func (r DynamicSchemaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dynamic_schema"
}

func (r DynamicSchemaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dynamic_attribute": schema.DynamicAttribute{
				Optional: true,
			},
			"object_attribute_with_dynamic": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"dynamic_attribute": types.DynamicType,
				},
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			"single_nested_block_with_dynamic": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"dynamic_attribute": schema.DynamicAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r DynamicSchemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DynamicSchemaResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicSchemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DynamicSchemaResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicSchemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DynamicSchemaResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r DynamicSchemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type DynamicSchemaResourceModel struct {
	DynamicAttribute             types.Dynamic `tfsdk:"dynamic_attribute"`
	ObjectAttributeWithDynamic   types.Object  `tfsdk:"object_attribute_with_dynamic"`
	SingleNestedBlockWithDynamic types.Object  `tfsdk:"single_nested_block_with_dynamic"`
}
