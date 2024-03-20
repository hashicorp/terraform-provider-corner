// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = MapFunction{}

func NewMapFunction() function.Function {
	return &MapFunction{}
}

type MapFunction struct{}

func (f MapFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "map"
}

func (f MapFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map_param",
			},
		},
		Return: function.MapReturn{
			ElementType: types.StringType,
		},
	}
}

func (f MapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg map[string]*string

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
