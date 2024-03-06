// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type errUnsupportedFunction string

func (e errUnsupportedFunction) Error() string {
	return "unsupported function: " + string(e)
}

type functionRouter map[string]tfprotov5.FunctionServer

func (f functionRouter) CallFunction(ctx context.Context, req *tfprotov5.CallFunctionRequest) (*tfprotov5.CallFunctionResponse, error) {
	fu, ok := f[req.Name]

	if !ok {
		return nil, errUnsupportedFunction(req.Name)
	}

	return fu.CallFunction(ctx, req)
}

func (f functionRouter) GetFunctions(ctx context.Context, req *tfprotov5.GetFunctionsRequest) (*tfprotov5.GetFunctionsResponse, error) {
	panic("not implemented")
}
