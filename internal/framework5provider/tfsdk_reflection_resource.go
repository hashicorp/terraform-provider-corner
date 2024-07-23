// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = TFSDKReflectionResource{}

func NewTFSDKReflectionResource() resource.Resource {
	return &TFSDKReflectionResource{}
}

// TFSDKReflectionResource is a smoke test for reflection logic on objects using the `tfsdk` field tags.
type TFSDKReflectionResource struct{}

func (r TFSDKReflectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tfsdk_reflection"
}

var nestedObjType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"nested_string": types.StringType,
	},
}

func (r TFSDKReflectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_bool": schema.BoolAttribute{
				Required: true,
			},
			"computed_bool": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"config_dynamic": schema.DynamicAttribute{
				Required: true,
			},
			"computed_dynamic": schema.DynamicAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Dynamic{
					dynamicplanmodifier.UseStateForUnknown(),
				},
			},
			"config_float64": schema.Float64Attribute{
				Required: true,
			},
			"computed_float64": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"config_int64": schema.Int64Attribute{
				Required: true,
			},
			"computed_int64": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"config_list": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"computed_list": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"config_map": schema.MapAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"computed_map": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
			"config_object": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"nested_string": types.StringType,
				},
				Required: true,
			},
			"computed_object": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"nested_string": types.StringType,
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
			},
			"config_string": schema.StringAttribute{
				Required: true,
			},
			"computed_string": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"config_set_nested_block": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"nested_string": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"computed_list_nested_block": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"nested_string": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r TFSDKReflectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TFSDKReflectionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedBool = types.BoolValue(true)
	data.ComputedDynamic = types.DynamicValue(types.StringValue("dynamic string"))
	data.ComputedFloat64 = types.Float64Value(1.2)
	data.ComputedInt64 = types.Int64Value(100)
	data.ComputedString = types.StringValue("computed string")

	data.ComputedList = types.ListValueMust(
		types.StringType,
		[]attr.Value{
			types.StringValue("computed"),
			types.StringValue("list"),
		},
	)
	data.ComputedMap = types.MapValueMust(
		types.StringType,
		map[string]attr.Value{
			"key6": types.StringValue("val6"),
			"key7": types.StringValue("val7"),
		},
	)

	data.ComputedObject = types.ObjectValueMust(
		nestedObjType.AttrTypes,
		map[string]attr.Value{
			"nested_string": types.StringValue("computed string"),
		},
	)

	data.ComputedListNestedBlock = types.ListValueMust(
		nestedObjType,
		[]attr.Value{
			types.ObjectValueMust(
				nestedObjType.AttrTypes,
				map[string]attr.Value{
					"nested_string": types.StringValue("computed string one"),
				},
			),
			types.ObjectValueMust(
				nestedObjType.AttrTypes,
				map[string]attr.Value{
					"nested_string": types.StringValue("computed string two"),
				},
			),
		},
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TFSDKReflectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TFSDKReflectionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TFSDKReflectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TFSDKReflectionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r TFSDKReflectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type TFSDKReflectionResourceModel struct {
	ConfigDynamic   types.Dynamic `tfsdk:"config_dynamic"`
	ComputedDynamic types.Dynamic `tfsdk:"computed_dynamic"`

	primitiveAttrs
	CollectionAttrs
	objectAttrs
	CollectionBlocks

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
	*EmbedIgnore   `tfsdk:"-"`
	embedIgnore    `tfsdk:"-"`
}

type embedIgnore struct {
	Field1 string
}

type EmbedIgnore struct {
	Field2 string
}

type primitiveAttrs struct {
	ConfigBool   *bool      `tfsdk:"config_bool"`
	ComputedBool types.Bool `tfsdk:"computed_bool"`

	ConfigFloat64   *float64      `tfsdk:"config_float64"`
	ComputedFloat64 types.Float64 `tfsdk:"computed_float64"`

	ConfigInt64   *int64      `tfsdk:"config_int64"`
	ComputedInt64 types.Int64 `tfsdk:"computed_int64"`

	ConfigString   *string      `tfsdk:"config_string"`
	ComputedString types.String `tfsdk:"computed_string"`

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}

type CollectionAttrs struct {
	ConfigList   []string   `tfsdk:"config_list"`
	ComputedList types.List `tfsdk:"computed_list"`

	MapAttrs

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}

type MapAttrs struct {
	ConfigMap   map[string]string `tfsdk:"config_map"`
	ComputedMap types.Map         `tfsdk:"computed_map"`

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}

type objectAttrs struct {
	ConfigObject   types.Object `tfsdk:"config_object"`
	ComputedObject types.Object `tfsdk:"computed_object"`

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}

type CollectionBlocks struct {
	setBlock
	listBlock

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}

type setBlock struct {
	ConfigSetNestedBlock types.Set `tfsdk:"config_set_nested_block"`

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}
type listBlock struct {
	ComputedListNestedBlock types.List `tfsdk:"computed_list_nested_block"`

	ExplicitIgnore string `tfsdk:"-"`
	implicitIgnore string //nolint
}
