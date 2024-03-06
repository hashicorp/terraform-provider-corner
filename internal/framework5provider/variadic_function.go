// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = VariadicFunction{}

func NewVariadicFunction() function.Function {
	return &VariadicFunction{}
}

type VariadicFunction struct{}

func (f VariadicFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "variadic"
}

func (f VariadicFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
		VariadicParameter: function.StringParameter{},
	}
}

func (f VariadicFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var varg []string

	resp.Error = req.Arguments.Get(ctx, &varg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, varg))
}
