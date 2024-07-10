// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = Int32Function{}

func NewInt32Function() function.Function {
	return &Int32Function{}
}

type Int32Function struct{}

func (f Int32Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "int32"
}

func (f Int32Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.Int32Parameter{
				Name: "int32_param",
			},
		},
		Return: function.Int32Return{},
	}
}

func (f Int32Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg int32

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
