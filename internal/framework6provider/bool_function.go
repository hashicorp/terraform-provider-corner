// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = BoolFunction{}

func NewBoolFunction() function.Function {
	return &BoolFunction{}
}

type BoolFunction struct{}

func (f BoolFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "bool"
}

func (f BoolFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.BoolParameter{
				Name: "bool_param",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f BoolFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg bool

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
