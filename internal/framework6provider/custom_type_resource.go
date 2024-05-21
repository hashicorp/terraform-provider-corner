// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = CustomTypeResource{}

func NewCustomTypeResource() resource.Resource {
	return &CustomTypeResource{}
}

// CustomTypeResource is for testing custom types.
type CustomTypeResource struct{}

func (r CustomTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_type"
}

func (r CustomTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"json_normalized": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				Computed:   true,
			},
			"json_exact": schema.StringAttribute{
				CustomType: jsontypes.ExactType{},
				Optional:   true,
				Computed:   true,
			},
			"ip_v4": schema.StringAttribute{
				CustomType: iptypes.IPv4AddressType{},
				Optional:   true,
				Computed:   true,
			},
			"ip_v6": schema.StringAttribute{
				CustomType: iptypes.IPv6AddressType{},
				Optional:   true,
				Computed:   true,
			},
			"ip_v4_cidr": schema.StringAttribute{
				CustomType: cidrtypes.IPv4PrefixType{},
				Optional:   true,
				Computed:   true,
			},
			"ip_v6_cidr": schema.StringAttribute{
				CustomType: cidrtypes.IPv6PrefixType{},
				Optional:   true,
				Computed:   true,
			},
			"time_rfc3339": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Optional:   true,
				Computed:   true,
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
	JSONNormalized jsontypes.Normalized `tfsdk:"json_normalized"`
	JSONExact      jsontypes.Exact      `tfsdk:"json_exact"`
	IPv4           iptypes.IPv4Address  `tfsdk:"ip_v4"`
	IPv6           iptypes.IPv6Address  `tfsdk:"ip_v6"`
	IPv4CIDR       cidrtypes.IPv4Prefix `tfsdk:"ip_v4_cidr"`
	IPv6CIDR       cidrtypes.IPv6Prefix `tfsdk:"ip_v6_cidr"`
	TimeRFC3339    timetypes.RFC3339    `tfsdk:"time_rfc3339"`
}
