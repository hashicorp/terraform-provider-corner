// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = StringFunction{}

func NewStringFunction() function.Function {
	return &StringFunction{}
}

type StringFunction struct{}

func (f StringFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "string2"
}

func (f StringFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{},
		},
		Return: function.StringReturn{},
	}
}

func (f StringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg string

	resp.Diagnostics.Append(req.Arguments.Get(ctx, &arg)...)

	resp.Diagnostics.Append(resp.Result.Set(ctx, arg)...)
}
