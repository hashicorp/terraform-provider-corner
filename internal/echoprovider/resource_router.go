// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package echoprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type errUnsupportedResource string

func (e errUnsupportedResource) Error() string {
	return "unsupported resource: " + string(e)
}

type resourceRouter map[string]tfprotov6.ResourceServer

func (r resourceRouter) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.ValidateResourceConfig(ctx, req)
}

func (r resourceRouter) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.UpgradeResourceState(ctx, req)
}

func (r resourceRouter) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.ReadResource(ctx, req)
}

func (r resourceRouter) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.PlanResourceChange(ctx, req)
}

func (r resourceRouter) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.ApplyResourceChange(ctx, req)
}

func (r resourceRouter) ImportResourceState(ctx context.Context, req *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TypeName)
	}
	return res.ImportResourceState(ctx, req)
}

func (r resourceRouter) MoveResourceState(ctx context.Context, req *tfprotov6.MoveResourceStateRequest) (*tfprotov6.MoveResourceStateResponse, error) {
	_, ok := r[req.TargetTypeName]
	if !ok {
		return nil, errUnsupportedResource(req.TargetTypeName)
	}
	// If this support ever needs to be added, this can follow the existing
	// pattern of calling res.MoveResourceState(ctx, req).
	return &tfprotov6.MoveResourceStateResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unsupported Resource Operation",
				Detail:   "MoveResourceState is not supported by this provider.",
			},
		},
	}, nil
}
