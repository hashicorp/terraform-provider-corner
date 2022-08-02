package framework

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

var stderr = os.Stderr

func New() provider.Provider {
	return &testProvider{}
}

type testProvider struct {
	client *backend.Client
}

func (p *testProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dummy": {
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
	p.client = client
}

func (p *testProvider) GetResources(_ context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"framework_user": resourceUserType{},
	}, nil
}

func (p *testProvider) GetDataSources(_ context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{}, nil
}
