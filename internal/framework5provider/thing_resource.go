// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	sdkv2schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ resource.Resource = ThingResource{}

func NewThingResource() resource.Resource {
	return &ThingResource{}
}

type ThingResource struct{}

func (r ThingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

var thingFrameworkSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
	},
}

var thingSDKv2Resource = &sdkv2schema.Resource{
	// .. other fields

	Schema: map[string]*sdkv2schema.Schema{
		"id": {
			Type:     sdkv2schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     sdkv2schema.TypeString,
			Required: true,
		},
	},
}

func (r ThingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = thingFrameworkSchema
}

func (r ThingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// 1. Typical framework-only data handling
	// ----------------------------------------------------------------------------
	//    - Read stuff into a struct
	//    - Set additional data on said struct
	//    - Set entire struct back to response
	// var data ThingResourceModel
	// resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }
	// data.ID = types.StringValue("id-123")
	// resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	// ----------------------------------------------------------------------------

	// 2. A hardcoded way of expressing #1
	//    - Just doing this to express some of the data structures underneath
	// ----------------------------------------------------------------------------
	// resp.State = tfsdk.State{
	// 	Raw: tftypes.NewValue(
	// 		thingFrameworkSchema.Type().TerraformType(ctx), // The schema has the plugin-go type we want
	// 		map[string]tftypes.Value{
	// 			"id":   tftypes.NewValue(tftypes.String, "id-123"),
	// 			"name": tftypes.NewValue(tftypes.String, "Austin Valle"),
	// 		},
	// 	),
	// 	Schema: thingFrameworkSchema,
	// }
	// ----------------------------------------------------------------------------

	// 3. SDKv2 to Framework
	// ----------------------------------------------------------------------------
	resp.State = tfsdk.State{
		Raw: tftypes.NewValue(
			thingSDKv2Resource.GetProtoType(ctx), // The schema has the plugin-go type we want (new method introduced to SDKv2)
			map[string]tftypes.Value{ // TODO: Can we use (ResourceData) -> map[string]tftypes.Value
				"id":   tftypes.NewValue(tftypes.String, "id-123"),
				"name": tftypes.NewValue(tftypes.String, "Austin Valle"),
			},
		),
		// TODO: We need to decide if this field is actually needed for tfsdk.Resource, maybe we can craft a bare minimum schema inside a helper function?
		// This will cause a panic in framework if we don't have a schema, but technically it's not needed! We just wrote all the internal logic
		// of ApplyResourceChange/protocol marshaling to rely on the schema. The type is the important part, which we do have.
		//
		// Ideas:
		// - We could make it so that Schema is not required, or just remove it completely from the tfsdk.Resource implementation to ensure no internal logic actually uses it
		//    - When we go to write this to the protocol, we would need to add a new function that creates a dynamic value using just the raw data (doable, just typically uses the schema)
		//    - The logic I'm referring to that would panic without a schema:
		//       - https://github.com/hashicorp/terraform-plugin-framework/blob/6997cf5e479d4e695436debac2e6222f07daf539/internal/toproto5/dynamic_value.go#L32
		// - We could write a very simple tftypes.Type -> fwschema.Schema converter helper function and put it in the terraform-plugin-framework Go module.
		//    - This would essentially just create a schema that had types and nothing else, possibly making even more confusing errors if we need other schema fields (Required/Computed/etc).
		// - We could write a sdkv2schema.Resource -> fwschema.Schema converter helper function, only mapping the fields that we know are 1 to 1 (or that we need)
		//    - Not sure where a helper like this would leave, we don't want SDKv2 or framework to reference each other :)
		Schema: thingFrameworkSchema,
	}

	if !thingSDKv2Resource.GetProtoType(ctx).Equal(thingFrameworkSchema.Type().TerraformType(ctx)) {
		panic("just to solidify, that these two schemas are the exact same terraform-plugin-go type")
	}
}

func (r ThingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ThingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r ThingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ThingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r ThingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type ThingResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
