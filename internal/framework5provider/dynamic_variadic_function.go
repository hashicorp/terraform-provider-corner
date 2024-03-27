// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = DynamicVariadicFunction{}

func NewDynamicVariadicFunction() function.Function {
	return &DynamicVariadicFunction{}
}

type DynamicVariadicFunction struct{}

func (f DynamicVariadicFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "dynamic_variadic"
}

func (f DynamicVariadicFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		VariadicParameter: function.DynamicParameter{
			Name: "dynamic_variadic_param",
		},
		Return: function.DynamicReturn{},
	}
}

func (f DynamicVariadicFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var varg types.Tuple
	resp.Error = req.Arguments.Get(ctx, &varg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, types.DynamicValue(varg)))
}
