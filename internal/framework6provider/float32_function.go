// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = Float32Function{}

func NewFloat32Function() function.Function {
	return &Float32Function{}
}

type Float32Function struct{}

func (f Float32Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "float32"
}

func (f Float32Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.Float32Parameter{
				Name: "float32_param",
			},
		},
		Return: function.Float32Return{},
	}
}

func (f Float32Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg float32

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
