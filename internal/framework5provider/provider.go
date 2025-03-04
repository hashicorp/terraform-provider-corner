// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

var (
	_ provider.ProviderWithFunctions          = (*testProvider)(nil)
	_ provider.ProviderWithEphemeralResources = (*testProvider)(nil)
)

func New() provider.Provider {
	return &testProvider{
		ephSpyClient: &EphemeralResourceSpyClient{},
	}
}

func NewWithEphemeralSpy(spy *EphemeralResourceSpyClient) provider.Provider {
	return &testProvider{
		ephSpyClient: spy,
	}
}

type testProvider struct {
	ephSpyClient *EphemeralResourceSpyClient
}

func (p *testProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "framework"
}

func (p *testProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dummy": schema.StringAttribute{
				Optional: true,
			},
			"deferral": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

func (p *testProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerConfig
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := backend.NewClient()
	if err != nil {
		resp.Diagnostics.AddError("Error initialising client", err.Error())
	}
	if req.ClientCapabilities.DeferralAllowed && config.Deferral.ValueBool() {
		resp.Deferred = &provider.Deferred{
			Reason: provider.DeferredReasonProviderConfigUnknown,
		}
	}
	resp.ResourceData = client
	resp.EphemeralResourceData = p.ephSpyClient
}

func (p *testProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSchemaResource,
		NewDeferredActionResource,
		NewDeferredActionPlanModificationResource,
		NewDynamicSchemaResource,
		NewDynamicComputedTypeChangeResource,
		NewTimeoutsResource,
		NewTimeTypesResource,
		NewUserResource,
		NewFloat32PrecisionResource,
		NewFloat64PrecisionResource,
		NewTFSDKReflectionResource,
		NewMoveStateResource,
		NewSetNestedBlockWithDefaultsResource,
		NewSetSemanticEqualityResource,
		NewCustomTypeResource,
	}
}

func (p *testProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *testProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewBoolFunction,
		NewDynamicFunction,
		NewFloat32Function,
		NewFloat64Function,
		NewInt32Function,
		NewInt64Function,
		NewListFunction,
		NewMapFunction,
		NewNumberFunction,
		NewObjectFunction,
		NewObjectWithDynamicFunction,
		NewSetFunction,
		NewStringFunction,
		NewVariadicFunction,
		NewDynamicVariadicFunction,
	}
}

func (p *testProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewSchemaEphemeralResource,
		NewEphemeralLifecycleResource,
	}
}

type providerConfig struct {
	Dummy    types.String `tfsdk:"dummy"`
	Deferral types.Bool   `tfsdk:"deferral"`
}
