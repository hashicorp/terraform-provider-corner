// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

type server struct {
	providerSchema     *tfprotov5.Schema
	providerMetaSchema *tfprotov5.Schema
	resourceSchemas    map[string]*tfprotov5.Schema
	dataSourceSchemas  map[string]*tfprotov5.Schema

	resourceRouter
	dataSourceRouter

	client *backend.Client
}

func (s *server) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider:          s.providerSchema,
		ProviderMeta:      s.providerMetaSchema,
		ResourceSchemas:   s.resourceSchemas,
		DataSourceSchemas: s.dataSourceSchemas,
	}, nil
}

func (s *server) PrepareProviderConfig(ctx context.Context, req *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	return &tfprotov5.PrepareProviderConfigResponse{
		PreparedConfig: req.Config,
	}, nil
}

func (s *server) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	var diags []*tfprotov5.Diagnostic
	client, err := backend.NewClient()
	if err != nil {
		diags = append(diags, &tfprotov5.Diagnostic{
			Summary:  "Error configuring provider",
			Detail:   fmt.Sprintf("Error instantiating new backend client: %s", err.Error()),
			Severity: tfprotov5.DiagnosticSeverityError,
		})
	}
	s.client = client
	return &tfprotov5.ConfigureProviderResponse{
		Diagnostics: diags,
	}, nil
}

func (s *server) StopProvider(ctx context.Context, req *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	return &tfprotov5.StopProviderResponse{}, nil
}

func Server() tfprotov5.ProviderServer {
	return &server{
		providerSchema: &tfprotov5.Schema{
			Block: &tfprotov5.SchemaBlock{},
		},
		dataSourceSchemas: map[string]*tfprotov5.Schema{
			"corner_time": {
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name:            "current",
							Type:            tftypes.String,
							Description:     "The current time in RFC3339 format.",
							DescriptionKind: tfprotov5.StringKindPlain,
							Computed:        true,
						},
						{
							Name:     "id",
							Type:     tftypes.String,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
		dataSourceRouter: dataSourceRouter{
			"corner_time": dataSourceTime{},
		},
	}
}
