package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = Int64Function{}

func NewInt64Function() function.Function {
	return &Int64Function{}
}

type Int64Function struct{}

func (f Int64Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "int64"
}

func (f Int64Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.Int64Parameter{},
		},
		Return: function.Int64Return{},
	}
}

func (f Int64Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var arg int64

	resp.Diagnostics.Append(req.Arguments.Get(ctx, &arg)...)

	resp.Diagnostics.Append(resp.Result.Set(ctx, arg)...)
}
