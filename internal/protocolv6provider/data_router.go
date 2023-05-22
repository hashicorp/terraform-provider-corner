// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type errUnsupportedDataSource string

func (e errUnsupportedDataSource) Error() string {
	return "unsupported data source: " + string(e)
}

type dataSourceRouter map[string]tfprotov6.DataSourceServer

func (d dataSourceRouter) ValidateDataResourceConfig(ctx context.Context, req *tfprotov6.ValidateDataResourceConfigRequest) (*tfprotov6.ValidateDataResourceConfigResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, errUnsupportedDataSource(req.TypeName)
	}
	return ds.ValidateDataResourceConfig(ctx, req)
}

func (d dataSourceRouter) ReadDataSource(ctx context.Context, req *tfprotov6.ReadDataSourceRequest) (*tfprotov6.ReadDataSourceResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, errUnsupportedDataSource(req.TypeName)
	}
	return ds.ReadDataSource(ctx, req)
}
