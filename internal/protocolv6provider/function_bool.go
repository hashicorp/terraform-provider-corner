// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type functionBool struct{}

func (f functionBool) CallFunction(context.Context, *tfprotov6.CallFunctionRequest) (*tfprotov6.CallFunctionResponse, error) {
	value, err := tfprotov6.NewDynamicValue(tftypes.Bool, tftypes.NewValue(tftypes.Bool, true))

	if err != nil {
		return nil, err
	}

	return &tfprotov6.CallFunctionResponse{
		Result: &value,
	}, nil
}

func (f functionBool) GetFunctions(context.Context, *tfprotov6.GetFunctionsRequest) (*tfprotov6.GetFunctionsResponse, error) {
	panic("implement me")
}
