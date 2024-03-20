// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = ListFunction{}

func NewListFunction() function.Function {
	return &ListFunction{}
}

type ListFunction struct{}

func (f ListFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "list"
}

func (f ListFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list_param",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f ListFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg []*string

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
