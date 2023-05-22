// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

type server struct {
	providerSchema     *tfprotov6.Schema
	providerMetaSchema *tfprotov6.Schema
	resourceSchemas    map[string]*tfprotov6.Schema
	dataSourceSchemas  map[string]*tfprotov6.Schema

	resourceRouter
	dataSourceRouter

	client *backend.Client
}

func (s *server) GetProviderSchema(ctx context.Context, req *tfprotov6.GetProviderSchemaRequest) (*tfprotov6.GetProviderSchemaResponse, error) {
	return &tfprotov6.GetProviderSchemaResponse{
		Provider:          s.providerSchema,
		ProviderMeta:      s.providerMetaSchema,
		ResourceSchemas:   s.resourceSchemas,
		DataSourceSchemas: s.dataSourceSchemas,
	}, nil
}

func (s *server) ValidateProviderConfig(ctx context.Context, req *tfprotov6.ValidateProviderConfigRequest) (*tfprotov6.ValidateProviderConfigResponse, error) {
	return &tfprotov6.ValidateProviderConfigResponse{
		PreparedConfig: req.Config,
	}, nil
}

func (s *server) ConfigureProvider(ctx context.Context, req *tfprotov6.ConfigureProviderRequest) (*tfprotov6.ConfigureProviderResponse, error) {
	var diags []*tfprotov6.Diagnostic
	client, err := backend.NewClient()
	if err != nil {
		diags = append(diags, &tfprotov6.Diagnostic{
			Summary:  "Error configuring provider",
			Detail:   fmt.Sprintf("Error instantiating new backend client: %s", err.Error()),
			Severity: tfprotov6.DiagnosticSeverityError,
		})
	}
	s.client = client
	return &tfprotov6.ConfigureProviderResponse{
		Diagnostics: diags,
	}, nil
}

func (s *server) StopProvider(ctx context.Context, req *tfprotov6.StopProviderRequest) (*tfprotov6.StopProviderResponse, error) {
	return &tfprotov6.StopProviderResponse{}, nil
}

func Server() tfprotov6.ProviderServer {
	return &server{
		providerSchema: &tfprotov6.Schema{
			Block: &tfprotov6.SchemaBlock{},
		},
		dataSourceSchemas: map[string]*tfprotov6.Schema{
			"corner_v6_time": {
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name:            "current",
							Type:            tftypes.String,
							Description:     "The current time in RFC3339 format.",
							DescriptionKind: tfprotov6.StringKindPlain,
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
			"corner_v6_time": dataSourceTime{},
		},
	}
}
