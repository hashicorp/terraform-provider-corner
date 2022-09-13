package provider1

import (
	"context"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &resourceNested{}
)

func NewNestedResource() resource.Resource {
	return &resourceNested{}
}

type resourceNested struct{}

func (r resourceNested) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nested"
}

func (r resourceNested) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]tfsdk.Block{
			"set": {
				Attributes: map[string]tfsdk.Attribute{
					"id": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
				Blocks: map[string]tfsdk.Block{
					"list": {
						Attributes: map[string]tfsdk.Attribute{
							"id": {
								Type:     types.StringType,
								Computed: true,
								PlanModifiers: []tfsdk.AttributePlanModifier{
									resource.UseStateForUnknown(),
								},
							},
						},
						NestingMode: tfsdk.BlockNestingModeList,
					},
				},
				NestingMode: tfsdk.BlockNestingModeList,
			},
		},
	}, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func id(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type nested struct {
	Id         types.String  `tfsdk:"id"`
	NestedItem []*NestedItem `tfsdk:"set"`
}

type NestedItem struct {
	Id               types.String        `tfsdk:"id"`
	NestedNestedItem []*NestedNestedItem `tfsdk:"list"`
}

type NestedNestedItem struct {
	Id types.String `tfsdk:"id"`
}

func (r resourceNested) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	nested := nested{}
	diags := req.Plan.Get(ctx, &nested)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if nested.Id.IsUnknown() {
		nested.Id = types.String{Value: id(8)}
	}

	for _, nst := range nested.NestedItem {
		if nst.Id.IsUnknown() {
			nst.Id = types.String{Value: id(8)}
		}

		for _, nstnst := range nst.NestedNestedItem {
			nstnst.Id = types.String{Value: id(8)}
		}
	}

	diags = resp.State.Set(ctx, nested)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
}

func (r resourceNested) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.State.Set(ctx, req.State.Raw)
}

func (r resourceNested) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.State.Set(ctx, req.Plan.Raw)
}

func (r resourceNested) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r resourceNested) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
