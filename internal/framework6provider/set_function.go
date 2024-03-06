// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = SetFunction{}

func NewSetFunction() function.Function {
	return &SetFunction{}
}

type SetFunction struct{}

func (f SetFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "set"
}

func (f SetFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.SetParameter{
				ElementType: types.StringType,
			},
		},
		Return: function.SetReturn{
			ElementType: types.StringType,
		},
	}
}

func (f SetFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg []*string

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
