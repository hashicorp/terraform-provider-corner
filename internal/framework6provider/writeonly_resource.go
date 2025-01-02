// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = WriteOnlyResource{}

func NewWriteOnlyResource() resource.Resource {
	return &WriteOnlyResource{}
}

type WriteOnlyResource struct{}

// WriteOnlyResource is a smoke test for schema attributes that contain write-only attributes
func (r WriteOnlyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_writeonly"
}

func (r WriteOnlyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"writeonly_custom_ipv6": schema.StringAttribute{
				Required:   true,
				WriteOnly:  true,
				CustomType: iptypes.IPv6AddressType{},
			},
			"writeonly_string": schema.StringAttribute{
				Required:  true,
				WriteOnly: true,
			},
			"writeonly_set": schema.SetAttribute{
				Required:    true,
				WriteOnly:   true,
				ElementType: types.StringType,
			},
			// TODO: At the moment, this raises an invalid plan error in Terraform core (Provider is successfully planning null, core
			// is rejecting this null value, instead saying it should be the config value, which is a bug)
			//  - https://hashicorp.slack.com/archives/C071HC4JJCC/p1734465766242319?thread_ts=1734465749.748579&cid=C071HC4JJCC
			//
			// "nested_object": schema.SingleNestedAttribute{
			// 	Required: true,
			// 	Attributes: map[string]schema.Attribute{
			// 		"string_attr": schema.StringAttribute{
			// 			Required:  true,
			// 		},
			// 		"writeonly_float64": schema.Float64Attribute{
			// 			Required:  true,
			// 			WriteOnly: true,
			// 		},
			// 	},
			// },
			"writeonly_nested_object": schema.SingleNestedAttribute{
				Required:  true,
				WriteOnly: true,
				Attributes: map[string]schema.Attribute{
					"writeonly_int64": schema.Int64Attribute{
						Required:  true,
						WriteOnly: true,
					},
					"writeonly_bool": schema.BoolAttribute{
						Required:  true,
						WriteOnly: true,
					},
					"writeonly_nested_list": schema.ListNestedAttribute{
						Required:  true,
						WriteOnly: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"writeonly_string": schema.StringAttribute{
									Required:  true,
									WriteOnly: true,
								},
							},
						},
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			"nested_block_list": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"string_attr": schema.StringAttribute{
							Required: true,
						},
						"writeonly_string": schema.StringAttribute{
							Required:  true,
							WriteOnly: true,
						},
					},
					Blocks: map[string]schema.Block{
						"double_nested_object": schema.SingleNestedBlock{
							Attributes: map[string]schema.Attribute{
								"bool_attr": schema.BoolAttribute{
									Required: true,
								},
								"writeonly_bool": schema.BoolAttribute{
									Required:  true,
									WriteOnly: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r WriteOnlyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WriteOnlyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var config WriteOnlyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(config.VerifyWriteOnlyData())
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r WriteOnlyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WriteOnlyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config WriteOnlyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(config.VerifyWriteOnlyData())
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r WriteOnlyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type WriteOnlyResourceModel struct {
	WriteOnlyCustomIPv6   iptypes.IPv6Address `tfsdk:"writeonly_custom_ipv6"`
	WriteOnlyString       types.String        `tfsdk:"writeonly_string"`
	WriteOnlySet          types.Set           `tfsdk:"writeonly_set"`
	WriteOnlyNestedObject types.Object        `tfsdk:"writeonly_nested_object"`
	NestedBlockList       types.List          `tfsdk:"nested_block_list"`
}

// VerifyWriteOnlyData compares the hardcoded test data for the write-only attributes in this resource, raising
// error diagnostics if the data differs from expectations.
func (m WriteOnlyResourceModel) VerifyWriteOnlyData() diag.Diagnostic {
	// Primitive write-only attributes
	expectedCustomIPv6 := "::"
	expectedString := "fakepassword"
	if m.WriteOnlyCustomIPv6.ValueString() != expectedCustomIPv6 {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("writeonly_custom_ipv6"),
			"Unexpected WriteOnly Value",
			fmt.Sprintf("wanted: %q, got: %q", expectedCustomIPv6, m.WriteOnlyCustomIPv6),
		)
	}
	if m.WriteOnlyString.ValueString() != expectedString {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("writeonly_string"),
			"Unexpected WriteOnly Value",
			fmt.Sprintf("wanted: %q, got: %q", expectedString, m.WriteOnlyString),
		)
	}

	// Collection write-only attribute
	expectedSet := types.SetValueMust(types.StringType, []attr.Value{types.StringValue("fake"), types.StringValue("password")})
	if !m.WriteOnlySet.Equal(expectedSet) {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("writeonly_set"),
			"Unexpected WriteOnly Value",
			fmt.Sprintf("wanted: %s, got: %s", expectedSet, m.WriteOnlySet),
		)
	}

	// Nested write-only attribute
	expectedListObjType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"writeonly_string": types.StringType,
		},
	}
	expectedObjectType := map[string]attr.Type{
		"writeonly_int64":       types.Int64Type,
		"writeonly_bool":        types.BoolType,
		"writeonly_nested_list": types.ListType{ElemType: expectedListObjType},
	}
	expectedObject := types.ObjectValueMust(expectedObjectType, map[string]attr.Value{
		"writeonly_int64": types.Int64Value(1234),
		"writeonly_bool":  types.BoolValue(true),
		"writeonly_nested_list": types.ListValueMust(expectedListObjType, []attr.Value{
			types.ObjectValueMust(expectedListObjType.AttributeTypes(), map[string]attr.Value{
				"writeonly_string": types.StringValue("fakepassword1"),
			}),
			types.ObjectValueMust(expectedListObjType.AttributeTypes(), map[string]attr.Value{
				"writeonly_string": types.StringValue("fakepassword2"),
			}),
		}),
	})

	if !m.WriteOnlyNestedObject.Equal(expectedObject) {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("writeonly_nested_object"),
			"Unexpected WriteOnly Value",
			fmt.Sprintf("wanted: %s, got: %s", expectedObject, m.WriteOnlyNestedObject),
		)
	}

	// Nested block with write-only attributes
	expectedBlockObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_attr":      types.BoolType,
			"writeonly_bool": types.BoolType,
		},
	}
	expectedBlockListObjType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"string_attr":          types.StringType,
			"writeonly_string":     types.StringType,
			"double_nested_object": expectedBlockObjectType,
		},
	}
	expectedListBlock := types.ListValueMust(expectedBlockListObjType, []attr.Value{
		types.ObjectValueMust(expectedBlockListObjType.AttributeTypes(), map[string]attr.Value{
			"string_attr":      types.StringValue("hello"),
			"writeonly_string": types.StringValue("fakepassword1"),
			"double_nested_object": types.ObjectValueMust(expectedBlockObjectType.AttributeTypes(), map[string]attr.Value{
				"bool_attr":      types.BoolValue(true),
				"writeonly_bool": types.BoolValue(false),
			}),
		}),
		types.ObjectValueMust(expectedBlockListObjType.AttributeTypes(), map[string]attr.Value{
			"string_attr":      types.StringValue("world"),
			"writeonly_string": types.StringValue("fakepassword2"),
			"double_nested_object": types.ObjectValueMust(expectedBlockObjectType.AttributeTypes(), map[string]attr.Value{
				"bool_attr":      types.BoolValue(false),
				"writeonly_bool": types.BoolValue(true),
			}),
		}),
	})

	if !m.NestedBlockList.Equal(expectedListBlock) {
		return diag.NewAttributeErrorDiagnostic(
			path.Root("nested_block_list"),
			"Unexpected WriteOnly Value",
			fmt.Sprintf("wanted: %s, got: %s", expectedListBlock, m.NestedBlockList),
		)
	}

	return nil
}
