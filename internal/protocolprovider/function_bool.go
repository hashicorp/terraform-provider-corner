// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type functionBool struct{}

func (f functionBool) CallFunction(context.Context, *tfprotov5.CallFunctionRequest) (*tfprotov5.CallFunctionResponse, error) {
	value, err := tfprotov5.NewDynamicValue(tftypes.Bool, tftypes.NewValue(tftypes.Bool, true))

	if err != nil {
		return nil, err
	}

	return &tfprotov5.CallFunctionResponse{
		Result: &value,
	}, nil
}

func (f functionBool) GetFunctions(context.Context, *tfprotov5.GetFunctionsRequest) (*tfprotov5.GetFunctionsResponse, error) {
	panic("implement me")
}
