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
	functions          map[string]*tfprotov5.Function

	resourceRouter
	dataSourceRouter
	functionRouter

	client *backend.Client
}

func (s *server) serverCapabilities() *tfprotov5.ServerCapabilities {
	return &tfprotov5.ServerCapabilities{
		GetProviderSchemaOptional: true,
	}
}

func (s *server) GetMetadata(ctx context.Context, req *tfprotov5.GetMetadataRequest) (*tfprotov5.GetMetadataResponse, error) {
	resp := &tfprotov5.GetMetadataResponse{
		DataSources:        make([]tfprotov5.DataSourceMetadata, 0, len(s.dataSourceSchemas)),
		Resources:          make([]tfprotov5.ResourceMetadata, 0, len(s.resourceSchemas)),
		ServerCapabilities: s.serverCapabilities(),
	}

	for typeName := range s.dataSourceSchemas {
		resp.DataSources = append(resp.DataSources, tfprotov5.DataSourceMetadata{
			TypeName: typeName,
		})
	}

	for typeName := range s.resourceSchemas {
		resp.Resources = append(resp.Resources, tfprotov5.ResourceMetadata{
			TypeName: typeName,
		})
	}

	return resp, nil
}

func (s *server) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider:           s.providerSchema,
		ProviderMeta:       s.providerMetaSchema,
		ResourceSchemas:    s.resourceSchemas,
		DataSourceSchemas:  s.dataSourceSchemas,
		ServerCapabilities: s.serverCapabilities(),
		Functions:          s.functions,
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

func (s *server) ValidateEphemeralResourceConfig(context.Context, *tfprotov5.ValidateEphemeralResourceConfigRequest) (*tfprotov5.ValidateEphemeralResourceConfigResponse, error) {
	return &tfprotov5.ValidateEphemeralResourceConfigResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Validate Config Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) OpenEphemeralResource(context.Context, *tfprotov5.OpenEphemeralResourceRequest) (*tfprotov5.OpenEphemeralResourceResponse, error) {
	return &tfprotov5.OpenEphemeralResourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Open Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) RenewEphemeralResource(context.Context, *tfprotov5.RenewEphemeralResourceRequest) (*tfprotov5.RenewEphemeralResourceResponse, error) {
	return &tfprotov5.RenewEphemeralResourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Renew Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func (s *server) CloseEphemeralResource(context.Context, *tfprotov5.CloseEphemeralResourceRequest) (*tfprotov5.CloseEphemeralResourceResponse, error) {
	return &tfprotov5.CloseEphemeralResourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unsupported Ephemeral Resource Close Operation",
				Detail:   "Ephemeral resources are not supported by this provider.",
			},
		},
	}, nil
}

func Server() tfprotov5.ProviderServer {
	return &server{
		providerSchema: &tfprotov5.Schema{
			Block: &tfprotov5.SchemaBlock{
				Attributes: []*tfprotov5.SchemaAttribute{
					{
						Name:     "deferral",
						Type:     tftypes.Bool,
						Optional: true,
					},
				},
			},
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
					},
				},
			},
			"corner_deferred_action": {
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name:            "current",
							Type:            tftypes.String,
							Description:     "The current time in RFC3339 format.",
							DescriptionKind: tfprotov5.StringKindPlain,
							Computed:        true,
						},
					},
				},
			},
		},
		dataSourceRouter: dataSourceRouter{
			"corner_time":            dataSourceTime{},
			"corner_deferred_action": dataDeferredAction{},
		},
		functions: map[string]*tfprotov5.Function{
			"bool": {
				Parameters: []*tfprotov5.FunctionParameter{
					{
						Name: "param",
						Type: tftypes.Bool,
					},
				},
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
		},
		functionRouter: functionRouter{
			"bool": functionBool{},
		},
	}
}
