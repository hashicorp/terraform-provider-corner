// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = StringFunction{}

func NewDynamicFunction() function.Function {
	return &DynamicFunction{}
}

type DynamicFunction struct{}

func (f DynamicFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "dynamic"
}

func (f DynamicFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name: "dynamic_param",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (f DynamicFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg types.Dynamic

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
