// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

var (
	_ provider.ProviderWithFunctions = (*testProvider)(nil)
)

func New() provider.Provider {
	return &testProvider{}
}

type testProvider struct{}

func (p *testProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "framework"
}

func (p *testProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dummy": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *testProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	client, err := backend.NewClient()
	if err != nil {
		resp.Diagnostics.AddError("Error initialising client", err.Error())
	}
	resp.ResourceData = client
}

func (p *testProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSchemaResource,
		NewTimeoutsResource,
		NewUserResource,
		NewFloat64PrecisionResource,
	}
}

func (p *testProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *testProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewBoolFunction,
		NewDynamicFunction,
		NewFloat64Function,
		NewInt64Function,
		NewListFunction,
		NewMapFunction,
		NewNumberFunction,
		NewObjectFunction,
		NewSetFunction,
		NewStringFunction,
		NewVariadicFunction,
		NewDynamicVariadicFunction,
	}
}
