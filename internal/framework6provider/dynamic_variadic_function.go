// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	var varg []types.Dynamic

	resp.Error = req.Arguments.Get(ctx, &varg)

	dynVals := make([]attr.Value, 0)

	for _, arg := range varg {
		dynVals = append(dynVals, types.DynamicValue(arg.UnderlyingValue()))
	}

	// Despite types.List not fully supporting dynamic types, in this restricted scenario it will work fine
	// as long as all the dynamic types coming in are the same.
	//
	// TODO: Switch this to a tuple once `terraform-plugin-testing` bug has been fixed with Tuple output:
	// 	- https://github.com/hashicorp/terraform-plugin-testing/issues/310
	listReturn, diags := types.ListValue(types.DynamicType, dynVals)
	if diags.HasError() {
		resp.Error = function.FuncErrorFromDiags(ctx, diags)
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, types.DynamicValue(listReturn)))
}
