// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes/refinement"
)

type resourceRefinements struct{}

func (r resourceRefinements) schemaType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"str_value": tftypes.String,
		},
	}
}

func (r resourceRefinements) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	plannedState, err := tftypes.ValueFromMsgPack(req.PlannedState.MsgPack, r.schemaType()) //nolint
	if err != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error decoding prior state",
					Detail:   fmt.Sprintf("Error decoding prior state: %s", err.Error()),
				},
			},
		}, nil
	}

	// Destroy op, state should be null
	if plannedState.IsNull() {
		return &tfprotov6.ApplyResourceChangeResponse{
			NewState: req.PlannedState,
		}, nil
	}

	objVal := map[string]tftypes.Value{}

	plannedState.As(&objVal) //nolint

	newStrValue := tftypes.NewValue(tftypes.String, "hello world!")

	// If the value exists in config, use it.
	if strValue, ok := objVal["str_value"]; ok && strValue.IsKnown() {
		newStrValue = strValue
	}

	newState, err := tfprotov6.NewDynamicValue(
		r.schemaType(),
		tftypes.NewValue(
			r.schemaType(),
			map[string]tftypes.Value{
				"str_value": newStrValue,
			},
		),
	)
	if err != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding state",
					Detail:   fmt.Sprintf("Error encoding planned state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.ApplyResourceChangeResponse{
		NewState: &newState,
	}, nil
}

func (r resourceRefinements) ImportResourceState(ctx context.Context, req *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {
	return &tfprotov6.ImportResourceStateResponse{}, nil
}

func (r resourceRefinements) MoveResourceState(ctx context.Context, req *tfprotov6.MoveResourceStateRequest) (*tfprotov6.MoveResourceStateResponse, error) {
	return &tfprotov6.MoveResourceStateResponse{}, nil
}

func (r resourceRefinements) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	if req.PriorState.MsgPack != nil {
		priorState, err := tftypes.ValueFromMsgPack(req.PriorState.MsgPack, r.schemaType()) //nolint
		if err != nil {
			return &tfprotov6.PlanResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error decoding prior state",
						Detail:   fmt.Sprintf("Error decoding prior state: %s", err.Error()),
					},
				},
			}, nil
		}

		// Update op, keep prior state
		if !priorState.IsNull() {
			return &tfprotov6.PlanResourceChangeResponse{
				PlannedState: req.PriorState,
			}, nil
		}
	}

	proposedNewState, err := tftypes.ValueFromMsgPack(req.ProposedNewState.MsgPack, r.schemaType()) //nolint
	if err != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error decoding proposed state",
					Detail:   fmt.Sprintf("Error decoding proposed state: %s", err.Error()),
				},
			},
		}, nil
	}

	objVal := map[string]tftypes.Value{}

	proposedNewState.As(&objVal) //nolint

	newStrValue := tftypes.NewValue(tftypes.String, tftypes.UnknownValue).
		Refine(refinement.Refinements{
			// str_value will never be null
			refinement.KeyNullness: refinement.NewNullness(false),
		})

	// If the value exists in config (unknown or known), keep it.
	if strValue, ok := objVal["str_value"]; ok && !strValue.IsNull() {
		newStrValue = strValue
	}

	plannedState, err := tfprotov6.NewDynamicValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"str_value": tftypes.String,
		},
	}, tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"str_value": tftypes.String,
		},
	}, map[string]tftypes.Value{
		"str_value": newStrValue,
	}))
	if err != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding state",
					Detail:   fmt.Sprintf("Error encoding planned state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.PlanResourceChangeResponse{
		PlannedState: &plannedState,
	}, nil
}

func (r resourceRefinements) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	return &tfprotov6.ReadResourceResponse{
		NewState: req.CurrentState,
	}, nil
}

func (r resourceRefinements) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {
	rawStateValue, _ := req.RawState.UnmarshalWithOpts(r.schemaType(), tfprotov6.UnmarshalOpts{
		ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
			IgnoreUndefinedAttributes: true,
		},
	})

	upgradedState, _ := tfprotov6.NewDynamicValue(r.schemaType(), rawStateValue)

	return &tfprotov6.UpgradeResourceStateResponse{
		UpgradedState: &upgradedState,
		Diagnostics:   []*tfprotov6.Diagnostic{},
	}, nil
}

func (r resourceRefinements) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	return &tfprotov6.ValidateResourceConfigResponse{}, nil
}
