// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = CustomTypeResource{}

func NewCustomTypeResource() resource.Resource {
	return &CustomTypeResource{}
}

// SchemaResource is for testing all schema types, excluding dynamic schema types. (see `DynamicSchemaResource`)
type CustomTypeResource struct{}

func (r CustomTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customtype"
}

func (r CustomTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ipv4test_attribute": schema.StringAttribute{
				Optional:   true,
				CustomType: iptypes.IPv4AddressType{},
			},
			"ipv6test_attribute": schema.StringAttribute{
				Optional:   true,
				CustomType: iptypes.IPv6AddressType{},
			},
		},
	}
}

func (r CustomTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CustomTypeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r CustomTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CustomTypeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r CustomTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CustomTypeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r CustomTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type CustomTypeResourceModel struct {
	CustomIPv4Attribute iptypes.IPv4Address `tfsdk:"ipv4test_attribute"`
	CustomIPv6Attribute iptypes.IPv6Address `tfsdk:"ipv6test_attribute"`
}
