// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = SetSemanticEqualityResource{}

func NewSetSemanticEqualityResource() resource.Resource {
	return &SetSemanticEqualityResource{}
}

// This resource tests that semantic equality for elements inside of a set are correctly executed
// Original bug: https://github.com/hashicorp/terraform-plugin-framework/issues/1061
type SetSemanticEqualityResource struct{}

func (r SetSemanticEqualityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_set_semantic_equality"
}

func (r SetSemanticEqualityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"set_of_ipv6": schema.SetAttribute{
				ElementType: iptypes.IPv6AddressType{},
				Required:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"set_nested_block": schema.SetNestedBlock{
				Validators: []validator.Set{
					setvalidator.IsRequired(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"ipv6": schema.StringAttribute{
							CustomType: iptypes.IPv6AddressType{},
							Required:   true,
						},
					},
				},
			},
		},
	}
}

func (r SetSemanticEqualityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SetSemanticEqualityResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Simulate remote API returning semantically equivalent IPv6 addresses
	data.shiftAndShorten()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetSemanticEqualityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SetSemanticEqualityResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Simulate remote API returning semantically equivalent IPv6 addresses
	data.shiftAndShorten()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetSemanticEqualityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SetSemanticEqualityResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Simulate remote API returning semantically equivalent IPv6 addresses
	data.shiftAndShorten()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r SetSemanticEqualityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type SetSemanticEqualityResourceModel struct {
	SetOfIPv6      types.Set `tfsdk:"set_of_ipv6"`
	SetNestedBlock types.Set `tfsdk:"set_nested_block"`
}

var setObjectWithIPv6 = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"ipv6": iptypes.IPv6AddressType{},
	},
}

// Shifts + switches data to shortened IPv6 addresses, but is semantically equal to test config
func (m *SetSemanticEqualityResourceModel) shiftAndShorten() {
	m.SetOfIPv6 = types.SetValueMust(iptypes.IPv6AddressType{}, []attr.Value{
		iptypes.NewIPv6AddressValue("2001:DB8::8:800:200C:417A"),
		iptypes.NewIPv6AddressValue("::FFFF:192.168.255.255"),
		iptypes.NewIPv6AddressValue("::"),
		iptypes.NewIPv6AddressValue("::101"),
	})

	m.SetNestedBlock = types.SetValueMust(setObjectWithIPv6, []attr.Value{
		types.ObjectValueMust(setObjectWithIPv6.AttributeTypes(), map[string]attr.Value{
			"ipv6": iptypes.NewIPv6AddressValue("::FFFF:192.168.255.255"),
		}),
		types.ObjectValueMust(setObjectWithIPv6.AttributeTypes(), map[string]attr.Value{
			"ipv6": iptypes.NewIPv6AddressValue("FF01::"),
		}),
		types.ObjectValueMust(setObjectWithIPv6.AttributeTypes(), map[string]attr.Value{
			"ipv6": iptypes.NewIPv6AddressValue("2001:DB8::8:800:200C:417A"),
		}),
	})
}
