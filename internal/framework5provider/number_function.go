// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = NumberFunction{}

func NewNumberFunction() function.Function {
	return &NumberFunction{}
}

type NumberFunction struct{}

func (f NumberFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "number"
}

func (f NumberFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.NumberParameter{
				Name: "number_param",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (f NumberFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg *big.Float

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
