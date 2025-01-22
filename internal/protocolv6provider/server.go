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
	functions          map[string]*tfprotov6.Function

	resourceRouter
	dataSourceRouter
	functionRouter

	client *backend.Client
}

func (s *server) serverCapabilities() *tfprotov6.ServerCapabilities {
	return &tfprotov6.ServerCapabilities{
		GetProviderSchemaOptional: true,
	}
}

func (s *server) GetMetadata(ctx context.Context, req *tfprotov6.GetMetadataRequest) (*tfprotov6.GetMetadataResponse, error) {
	resp := &tfprotov6.GetMetadataResponse{
		DataSources:        make([]tfprotov6.DataSourceMetadata, 0, len(s.dataSourceSchemas)),
		Resources:          make([]tfprotov6.ResourceMetadata, 0, len(s.resourceSchemas)),
		ServerCapabilities: s.serverCapabilities(),
	}

	for typeName := range s.dataSourceSchemas {
		resp.DataSources = append(resp.DataSources, tfprotov6.DataSourceMetadata{
			TypeName: typeName,
		})
	}

	for typeName := range s.resourceSchemas {
		resp.Resources = append(resp.Resources, tfprotov6.ResourceMetadata{
			TypeName: typeName,
		})
	}

	return resp, nil
}

func (s *server) GetProviderSchema(ctx context.Context, req *tfprotov6.GetProviderSchemaRequest) (*tfprotov6.GetProviderSchemaResponse, error) {
	return &tfprotov6.GetProviderSchemaResponse{
		Provider:           s.providerSchema,
		ProviderMeta:       s.providerMetaSchema,
		ResourceSchemas:    s.resourceSchemas,
		DataSourceSchemas:  s.dataSourceSchemas,
		ServerCapabilities: s.serverCapabilities(),
		Functions:          s.functions,
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

func (s *server) ValidateEphemeralResourceConfig(context.Context, *tfprotov6.ValidateEphemeralResourceConfigRequest) (*tfprotov6.ValidateEphemeralResourceConfigResponse, error) {
	return &tfprotov6.ValidateEphemeralResourceConfigResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Validate Config Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) OpenEphemeralResource(context.Context, *tfprotov6.OpenEphemeralResourceRequest) (*tfprotov6.OpenEphemeralResourceResponse, error) {
	return &tfprotov6.OpenEphemeralResourceResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Open Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) RenewEphemeralResource(context.Context, *tfprotov6.RenewEphemeralResourceRequest) (*tfprotov6.RenewEphemeralResourceResponse, error) {
	return &tfprotov6.RenewEphemeralResourceResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Renew Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) CloseEphemeralResource(context.Context, *tfprotov6.CloseEphemeralResourceRequest) (*tfprotov6.CloseEphemeralResourceResponse, error) {
	return &tfprotov6.CloseEphemeralResourceResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Close Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
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
					},
				},
			},
			"corner_v6_deferred_action": {
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name:            "current",
							Type:            tftypes.String,
							Description:     "The current time in RFC3339 format.",
							DescriptionKind: tfprotov6.StringKindPlain,
							Computed:        true,
						},
					},
				},
			},
		},
		dataSourceRouter: dataSourceRouter{
			"corner_v6_time":            dataSourceTime{},
			"corner_v6_deferred_action": dataSourceDeferredAction{},
		},
		functions: map[string]*tfprotov6.Function{
			"bool": {
				Parameters: []*tfprotov6.FunctionParameter{
					{
						Name: "param",
						Type: tftypes.Bool,
					},
				},
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
		},
		functionRouter: functionRouter{
			"bool": functionBool{},
		},
	}
}
