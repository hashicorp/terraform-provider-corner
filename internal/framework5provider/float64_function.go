// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = Float64Function{}

func NewFloat64Function() function.Function {
	return &Float64Function{}
}

type Float64Function struct{}

func (f Float64Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "float64"
}

func (f Float64Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.Float64Parameter{},
		},
		Return: function.Float64Return{},
	}
}

func (f Float64Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg float64

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
