// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type errUnsupportedAction string

func (e errUnsupportedAction) Error() string {
	return "unsupported action: " + string(e)
}

type actionRouter struct {
	ActionRoutes     map[string]tfprotov6.ActionServer
	contextCancels   map[string]context.CancelFunc
	contextCancelsMu sync.Mutex
	CallbackServers  map[string]tfprotov6.InvokeActionCallBackServer
}

func (a *actionRouter) registerContext(in context.Context, typeName string, callbackServer tfprotov6.InvokeActionCallBackServer) (context.Context, string) {
	ctx, cancel := context.WithCancel(in)
	a.contextCancelsMu.Lock()
	defer a.contextCancelsMu.Unlock()
	cancellationToken := typeName + randomString(32)
	a.contextCancels[cancellationToken] = cancel
	a.CallbackServers[cancellationToken] = callbackServer
	return ctx, cancellationToken
}

func (a *actionRouter) cancelActionContext(ctx context.Context, token string) *tfprotov6.Diagnostic {
	a.contextCancelsMu.Lock()
	defer a.contextCancelsMu.Unlock()
	//tflog.Debug(ctx, "Cancel Action Context")
	if cancel, ok := a.contextCancels[token]; ok {
		if cancel != nil {
			cancel()
			a.contextCancels[token] = nil
		}
		if callback, ok := a.CallbackServers[token]; ok {
			err := callback.Send(ctx, &tfprotov6.ProgressActionEvent{
				StdOut: []string{"Action got cancelled"},
			})
			if err != nil {
				return &tfprotov6.Diagnostic{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error cancelling action",
					Detail:   fmt.Sprintf("Error sending cancelled progress action event original error: %s", err),
				}
			}

			err = callback.Send(ctx, &tfprotov6.CancelledActionEvent{})
			if err != nil {
				return &tfprotov6.Diagnostic{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error cancelling action",
					Detail:   fmt.Sprintf("Error sending cancelled action event original error: %s", err),
				}
			}
		} else {
			return &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error cancelling action",
				Detail:   fmt.Sprintf("Cannot find callback server for token %s", token),
			}
		}
	} else {
		//tflog.Debug(ctx, "Cancel Action Context returns error")
		return &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error cancelling action",
			Detail:   fmt.Sprintf("Cannot find cancellation contect for token %s", token),
		}
	}
	//tflog.Debug(ctx, "Cancel Action Context returns")
	return nil
}

func (a *actionRouter) cancelRegisteredContexts(_ context.Context) {
	a.contextCancelsMu.Lock()
	defer a.contextCancelsMu.Unlock()
	for _, cancel := range a.contextCancels {
		cancel()
	}
	a.contextCancels = nil
}

func (a *actionRouter) PlanAction(ctx context.Context, req *tfprotov6.PlanActionRequest) (*tfprotov6.PlanActionResponse, error) {
	action, ok := a.ActionRoutes[req.TypeName]
	if !ok {
		return nil, errUnsupportedAction(req.TypeName)
	}
	return action.PlanAction(ctx, req)
}

func (a *actionRouter) InvokeAction(ctx context.Context, req *tfprotov6.InvokeActionRequest, resp *tfprotov6.InvokeActionResponse) error {
	action, ok := a.ActionRoutes[req.TypeName]
	if !ok {
		return errUnsupportedAction(req.TypeName)
	}

	ctx, token := a.registerContext(ctx, req.TypeName, resp.CallbackServer)
	err := resp.CallbackServer.Send(ctx, &tfprotov6.StartedActionEvent{
		CancellationToken: token,
	})
	if err != nil {
		return err
	}
	return action.InvokeAction(ctx, req, resp)
}

func (a *actionRouter) CancelAction(ctx context.Context, req *tfprotov6.CancelActionRequest) (*tfprotov6.CancelActionResponse, error) {
	diag := a.cancelActionContext(ctx, req.CancellationToken)
	if diag != nil {
		return &tfprotov6.CancelActionResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				diag,
			},
		}, nil
	}
	return &tfprotov6.CancelActionResponse{}, nil
}

func randomString(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
