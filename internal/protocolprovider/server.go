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
		MoveResourceState:         true,
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
		resourceSchemas: map[string]*tfprotov5.Schema{
			"corner_writeonly_datacheck":                      resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_planerror":            resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_applyerror":           resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_readerror":            resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_importerror":          resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_moveresourceerror":    resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_datacheck_upgraderesourceerror": resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_legacy_datacheck":               resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_legacy_datacheck_planerror":     resourceWriteOnlyDataCheck{}.schema(),
			"corner_writeonly_legacy_datacheck_applyerror":    resourceWriteOnlyDataCheck{}.schema(),
		},
		resourceRouter: resourceRouter{
			"corner_writeonly_datacheck": resourceWriteOnlyDataCheck{},
			"corner_writeonly_datacheck_planerror": resourceWriteOnlyDataCheck{
				planDataError: true,
			},
			"corner_writeonly_datacheck_applyerror": resourceWriteOnlyDataCheck{
				applyDataError: true,
			},
			"corner_writeonly_datacheck_readerror": resourceWriteOnlyDataCheck{
				readDataError: true,
			},
			"corner_writeonly_datacheck_importerror": resourceWriteOnlyDataCheck{
				importDataError: true,
			},
			"corner_writeonly_datacheck_moveresourceerror": resourceWriteOnlyDataCheck{
				moveResourceDataError: true,
			},
			"corner_writeonly_datacheck_upgraderesourceerror": resourceWriteOnlyDataCheck{
				upgradeResourceDataError: true,
			},
			"corner_writeonly_legacy_datacheck": resourceWriteOnlyDataCheck{
				enableLegacyTypeSystem: true,
			},
			// MAINTAINER NOTE: The only RPCs that have legacy type system flags are plan/apply
			"corner_writeonly_legacy_datacheck_planerror": resourceWriteOnlyDataCheck{
				enableLegacyTypeSystem: true,
				planDataError:          true,
			},
			"corner_writeonly_legacy_datacheck_applyerror": resourceWriteOnlyDataCheck{
				enableLegacyTypeSystem: true,
				applyDataError:         true,
			},
		},
	}
}
