// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = ObjectFunction{}

func NewObjectFunction() function.Function {
	return &ObjectFunction{}
}

type ObjectFunction struct{}

func (f ObjectFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "object"
}

func (f ObjectFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.ObjectParameter{
				AttributeTypes: map[string]attr.Type{
					"attr1": types.StringType,
					"attr2": types.Int64Type,
				},
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"attr1": types.StringType,
				"attr2": types.Int64Type,
			},
		},
	}
}

func (f ObjectFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg struct {
		Attr1 *string `tfsdk:"attr1"`
		Attr2 *int64  `tfsdk:"attr2"`
	}

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
