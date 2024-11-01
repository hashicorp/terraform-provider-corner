package echoprovider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ tfprotov6.ResourceServer = echoResource{}

type echoResource struct {
	providerConfig *tfprotov6.DynamicValue
}

var echoSchemaType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"data": tftypes.DynamicPseudoType,
	},
}

func (e echoResource) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	plannedState, err := req.PlannedState.Unmarshal(echoSchemaType)
	if err != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Detail:   fmt.Sprintf("error unmarhsaling planned state data: %s", err.Error()),
				},
			},
		}, nil
	}

	// Destroy Op, return the null from planned state
	if plannedState.IsNull() {
		return &tfprotov6.ApplyResourceChangeResponse{
			NewState: req.PlannedState,
		}, nil
	}

	if !plannedState.IsFullyKnown() {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Detail:   "echo_resource encountered an unexpected unknown value, this resource is only meant to echo configuration from the provider config.",
				},
			},
		}, nil
	}
	// Take the provider config verbatim and put back into state. It shares the same schema
	// as the echo resource, so the data types/value should match up and there shouldn't be any
	// unknown values present
	return &tfprotov6.ApplyResourceChangeResponse{
		NewState: e.providerConfig,
	}, nil
}

func (e echoResource) ImportResourceState(ctx context.Context, req *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {
	return &tfprotov6.ImportResourceStateResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   "import not supported",
			},
		},
	}, nil
}

func (e echoResource) MoveResourceState(ctx context.Context, req *tfprotov6.MoveResourceStateRequest) (*tfprotov6.MoveResourceStateResponse, error) {
	return &tfprotov6.MoveResourceStateResponse{
		Diagnostics: []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   "move state not supported",
			},
		},
	}, nil
}

func (e echoResource) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	return &tfprotov6.PlanResourceChangeResponse{
		PlannedState: req.ProposedNewState,
	}, nil
}

func (e echoResource) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	return &tfprotov6.ReadResourceResponse{
		NewState: req.CurrentState,
	}, nil
}

func (e echoResource) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {

	rawStateValue, err := req.RawState.UnmarshalWithOpts(
		echoSchemaType,
		tfprotov6.UnmarshalOpts{
			ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
				IgnoreUndefinedAttributes: true,
			},
		},
	)

	if err != nil {
		return &tfprotov6.UpgradeResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Detail:   fmt.Sprintf("error unmarhsaling raw state data: %s", err.Error()),
				},
			},
		}, nil
	}

	rawDynamicValue, err := tfprotov6.NewDynamicValue(rawStateValue.Type(), rawStateValue)
	if err != nil {
		return &tfprotov6.UpgradeResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Detail:   fmt.Sprintf("error creating dynamic value from raw state data: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.UpgradeResourceStateResponse{
		UpgradedState: &rawDynamicValue,
		Diagnostics:   []*tfprotov6.Diagnostic{},
	}, nil
}

func (e echoResource) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	return &tfprotov6.ValidateResourceConfigResponse{}, nil
}
