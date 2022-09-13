package provider2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func New() provider.Provider {
	return &testProvider{}
}

type testProvider struct{}

func (p *testProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tf6muxprovider"
}

func (p *testProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"example": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
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
		NewUserResource,
	}
}

func (p *testProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
