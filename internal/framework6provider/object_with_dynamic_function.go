// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = ObjectWithDynamicFunction{}

func NewObjectWithDynamicFunction() function.Function {
	return &ObjectWithDynamicFunction{}
}

type ObjectWithDynamicFunction struct{}

func (f ObjectWithDynamicFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "object_with_dynamic"
}

func (f ObjectWithDynamicFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.ObjectParameter{
				AttributeTypes: map[string]attr.Type{
					"dynamic_attr1": types.DynamicType,
					"dynamic_attr2": types.DynamicType,
				},
				Name: "object_param",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"dynamic_attr1": types.DynamicType,
				"dynamic_attr2": types.DynamicType,
			},
		},
	}
}

func (f ObjectWithDynamicFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg struct {
		DynamicAttr1 types.Dynamic `tfsdk:"dynamic_attr1"`
		DynamicAttr2 types.Dynamic `tfsdk:"dynamic_attr2"`
	}

	resp.Error = req.Arguments.Get(ctx, &arg)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arg))
}
