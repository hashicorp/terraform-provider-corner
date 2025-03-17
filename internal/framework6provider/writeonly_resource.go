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
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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
			"computed_attr": schema.StringAttribute{
				Computed: true,
			},
			"writeonly_custom_ipv6": schema.StringAttribute{
				Optional:   true,
				WriteOnly:  true,
				CustomType: iptypes.IPv6AddressType{},
			},
			"writeonly_string": schema.StringAttribute{
				Optional:  true,
				WriteOnly: true,
			},
			"writeonly_list": schema.ListAttribute{
				Optional:    true,
				WriteOnly:   true,
				ElementType: types.StringType,
			},
			"nested_map": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"string_attr": schema.StringAttribute{
							Required: true,
						},
						"writeonly_float64": schema.Float64Attribute{
							Optional:  true,
							WriteOnly: true,
						},
					},
				},
			},
			"writeonly_nested_object": schema.SingleNestedAttribute{
				Optional:  true,
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
							Optional:  true,
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
									Optional:  true,
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
	var config WriteOnlyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(VerifyWriteOnlyData(ctx, req.Config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.ComputedAttribute = types.StringValue("computed_val")

	// Since all attributes are in configuration, we write it back directly to test that the write-only attributes
	// are nulled out before sending back to TF Core.
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WriteOnlyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.WriteOnlyString = types.StringValue("this shouldn't cause an error!")
	data.ComputedAttribute = types.StringValue("computed_val")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r WriteOnlyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config WriteOnlyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(VerifyWriteOnlyData(ctx, req.Config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.ComputedAttribute = types.StringValue("computed_val")

	// Since all attributes are in configuration, we write it back directly to test that the write-only attributes
	// are nulled out before sending back to TF Core.
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func (r WriteOnlyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type WriteOnlyResourceModel struct {
	ComputedAttribute     types.String        `tfsdk:"computed_attr"`
	WriteOnlyCustomIPv6   iptypes.IPv6Address `tfsdk:"writeonly_custom_ipv6"`
	WriteOnlyString       types.String        `tfsdk:"writeonly_string"`
	WriteOnlyList         types.List          `tfsdk:"writeonly_list"`
	WriteOnlyNestedObject types.Object        `tfsdk:"writeonly_nested_object"`
	NestedMap             types.Map           `tfsdk:"nested_map"`
	NestedBlockList       types.List          `tfsdk:"nested_block_list"`
}

// VerifyWriteOnlyData compares the hardcoded test data for the write-only attributes in this resource, raising
// error diagnostics if the data differs from expectations.
func VerifyWriteOnlyData(ctx context.Context, cfg tfsdk.Config) diag.Diagnostics {
	var diags diag.Diagnostics
	// Primitive write-only attributes
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("writeonly_custom_ipv6"), iptypes.NewIPv6AddressValue("::"))...)
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("writeonly_string"), types.StringValue("fakepassword"))...)

	// Collection write-only attribute
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("writeonly_list"), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("fake"), types.StringValue("password")}))...)

	// Nested map with write-only attribute
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_map").AtMapKey("key1").AtName("writeonly_float64"), types.Float64Value(10))...)
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_map").AtMapKey("key2").AtName("writeonly_float64"), types.Float64Value(20))...)

	// Nested write-only object attribute
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

	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("writeonly_nested_object"), expectedObject)...)

	// Nested block with write-only attributes
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_block_list").AtListIndex(0).AtName("writeonly_string"), types.StringValue("fakepassword1"))...)
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_block_list").AtListIndex(1).AtName("writeonly_string"), types.StringValue("fakepassword2"))...)
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_block_list").AtListIndex(0).AtName("double_nested_object").AtName("writeonly_bool"), types.BoolValue(false))...)
	diags.Append(assertWriteOnlyVal(ctx, cfg, path.Root("nested_block_list").AtListIndex(1).AtName("double_nested_object").AtName("writeonly_bool"), types.BoolValue(true))...)

	return diags
}

// Asserts a write-only value in configuration, if the value is null it will return without an error (allowing the attribute to be optional)
func assertWriteOnlyVal[T attr.Value](ctx context.Context, cfg tfsdk.Config, p path.Path, expectedVal T) diag.Diagnostics {
	var writeOnlyVal T
	diags := cfg.GetAttribute(ctx, p, &writeOnlyVal)
	if diags.HasError() {
		// All the paths are hardcoded in the resource, so this scenario shouldn't occur unless there is a schema/path mismatch or addition
		return diags
	}

	if !writeOnlyVal.IsNull() && !writeOnlyVal.Equal(expectedVal) {
		return diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(
				p,
				"Unexpected WriteOnly Value",
				fmt.Sprintf("wanted: %s, got: %s", expectedVal, writeOnlyVal),
			),
		}
	}

	return nil
}
