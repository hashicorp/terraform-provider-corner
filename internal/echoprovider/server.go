// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package echoprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
	// Quick hack to save the ephemeral data from the config for later
	s.resourceRouter["echo_resource"] = echoResource{
		providerConfig: req.Config,
	}

	return &tfprotov6.ConfigureProviderResponse{}, nil
}

func (s *server) StopProvider(ctx context.Context, req *tfprotov6.StopProviderRequest) (*tfprotov6.StopProviderResponse, error) {
	return &tfprotov6.StopProviderResponse{}, nil
}

func NewServer() func() (tfprotov6.ProviderServer, error) {
	return func() (tfprotov6.ProviderServer, error) {
		return &server{
			// Both provider config + echo_resource have the same schema
			providerSchema: &tfprotov6.Schema{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name:            "data",
							Type:            tftypes.DynamicPseudoType,
							DescriptionKind: tfprotov6.StringKindPlain,
							Required:        true,
						},
					},
				},
			},
			resourceSchemas: map[string]*tfprotov6.Schema{
				// Both provider config + echo_resource have the same schema
				"echo_resource": {
					Block: &tfprotov6.SchemaBlock{
						Attributes: []*tfprotov6.SchemaAttribute{
							{
								Name:            "data",
								Type:            tftypes.DynamicPseudoType,
								DescriptionKind: tfprotov6.StringKindPlain,
								Computed:        true,
							},
						},
					},
				},
			},
			resourceRouter: resourceRouter{
				"echo_resource": echoResource{},
			},
		}, nil
	}
}
