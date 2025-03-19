// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type errUnsupportedAction string

func (e errUnsupportedAction) Error() string {
	return "unsupported action: " + string(e)
}

type actionRouter map[string]tfprotov6.ActionServer

func (a actionRouter) PlanAction(ctx context.Context, req *tfprotov6.PlanActionRequest) (*tfprotov6.PlanActionResponse, error) {
	action, ok := a[req.TypeName]
	if !ok {
		return nil, errUnsupportedAction(req.TypeName)
	}
	return action.PlanAction(ctx, req)
}

func (a actionRouter) InvokeAction(ctx context.Context, req *tfprotov6.InvokeActionRequest, resp *tfprotov6.InvokeActionResponse) error {
	action, ok := a[req.TypeName]
	if !ok {
		return errUnsupportedAction(req.TypeName)
	}
	return action.InvokeAction(ctx, req, resp)
}

func (a actionRouter) CancelAction(ctx context.Context, req *tfprotov6.CancelActionRequest) (*tfprotov6.CancelActionResponse, error) {
	return nil, fmt.Errorf("unimplemented")
	//action, ok := a[req.TypeName]
	//if !ok {
	//	return errUnsupportedAction(req.TypeName)
	//}
	//return action.CancelAction(ctx, req)
}
