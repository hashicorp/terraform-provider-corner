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
	"net/netip"
)

var _ resource.Resource = CustomTypeResource{}

func NewCustomTypeResource() resource.Resource {
	return &CustomTypeResource{}
}

type CustomTypeResource struct{}

func (r CustomTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customtype"
}

func (r CustomTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"json_normalized_attribute": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
			},
			"json_exact_attribute": schema.StringAttribute{
				CustomType: jsontypes.ExactType{},
				Optional:   true,
			},
			"ip_v4_attribute": schema.StringAttribute{
				CustomType: iptypes.IPv4AddressType{},
				Optional:   true,
			},
			"ip_v6_attribute": schema.StringAttribute{
				CustomType: iptypes.IPv6AddressType{},
				Optional:   true,
			},
			"ip_v4_cidr_attribute": schema.StringAttribute{
				CustomType: cidrtypes.IPv4PrefixType{},
				Optional:   true,
			},
			"ip_v6_cidr_attribute": schema.StringAttribute{
				CustomType: cidrtypes.IPv6PrefixType{},
				Optional:   true,
			},
			"time_rfc3339_attribute": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Optional:   true,
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

	// Testing IPv6 Semantic Equality
	if !data.CustomIPv6Attribute.IsNull() {
		tempIpAddr := data.CustomIPv6Attribute.ValueString()
		currentIpAddr, _ := netip.ParseAddr(tempIpAddr)
		expandedIpAddr := currentIpAddr.StringExpanded()
		newIpAddr := iptypes.NewIPv6AddressValue(expandedIpAddr)
		data.CustomIPv6Attribute = newIpAddr
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
	JSONNormalized      jsontypes.Normalized `tfsdk:"json_normalized_attribute"`
	JSONExact           jsontypes.Exact      `tfsdk:"json_exact_attribute"`
	CustomIPv4Attribute iptypes.IPv4Address  `tfsdk:"ip_v4_attribute"`
	CustomIPv6Attribute iptypes.IPv6Address  `tfsdk:"ip_v6_attribute"`
	IPv4CIDR            cidrtypes.IPv4Prefix `tfsdk:"ip_v4_cidr_attribute"`
	IPv6CIDR            cidrtypes.IPv6Prefix `tfsdk:"ip_v6_cidr_attribute"`
	TimeRFC3339         timetypes.RFC3339    `tfsdk:"time_rfc3339_attribute"`
}
