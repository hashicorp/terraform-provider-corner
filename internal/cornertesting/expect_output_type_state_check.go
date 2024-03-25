// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cornertesting

import (
	"context"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/zclconf/go-cty/cty"
)

var _ statecheck.StateCheck = expectOutputType{}

type expectOutputType struct {
	outputAddress string
	expectedType  cty.Type
}

func (e expectOutputType) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
	var output *tfjson.StateOutput

	if req.State == nil {
		resp.Error = fmt.Errorf("state is nil")

		return
	}

	if req.State.Values == nil {
		resp.Error = fmt.Errorf("state does not contain any state values")

		return
	}

	for address, oc := range req.State.Values.Outputs {
		if e.outputAddress == address {
			output = oc

			break
		}
	}

	if output == nil {
		resp.Error = fmt.Errorf("%s - Output not found in state", e.outputAddress)

		return
	}

	if !output.Type.Equals(e.expectedType) {
		resp.Error = fmt.Errorf("expected %q output type to be %q, got %q", e.outputAddress, e.expectedType.FriendlyName(), output.Type.FriendlyName())
	}
}

// ExpectOutputType returns a state check that asserts that the specified output has a cty.Type that matches `expectedType`
//
// NOTE: This check is only useful for output values that are populated by dynamic attributes or dynamic function returns.
func ExpectOutputType(outputAddress string, expectedType cty.Type) statecheck.StateCheck {
	return expectOutputType{
		outputAddress: outputAddress,
		expectedType:  expectedType,
	}
}
