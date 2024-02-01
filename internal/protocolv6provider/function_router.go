// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type errUnsupportedFunction string

func (e errUnsupportedFunction) Error() string {
	return "unsupported function: " + string(e)
}

type functionRouter map[string]tfprotov6.FunctionServer

func (f functionRouter) CallFunction(ctx context.Context, req *tfprotov6.CallFunctionRequest) (*tfprotov6.CallFunctionResponse, error) {
	fu, ok := f[req.Name]

	if !ok {
		return nil, errUnsupportedFunction(req.Name)
	}

	return fu.CallFunction(ctx, req)
}

func (f functionRouter) GetFunctions(ctx context.Context, req *tfprotov6.GetFunctionsRequest) (*tfprotov6.GetFunctionsResponse, error) {
	panic("not implemented")
}
