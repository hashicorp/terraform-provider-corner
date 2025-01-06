// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceWriteOnlyDataCheck struct {
	enableLegacyTypeSystem bool
	applyDataError         bool
	planDataError          bool
	readDataError          bool
	importDataError        bool
}

func (r resourceWriteOnlyDataCheck) schema() *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Block: &tfprotov5.SchemaBlock{
			Attributes: []*tfprotov5.SchemaAttribute{
				{
					// Only used for import testing
					Name:     "id",
					Type:     tftypes.String,
					Computed: true,
				},
				{
					Name:      "writeonly_attr",
					Type:      tftypes.String,
					Required:  true,
					WriteOnly: true,
				},
			},
		},
	}
}

// nonNullWriteOnlyData is used to produce data which will raise an error diagnostic in Terraform core.
func (r resourceWriteOnlyDataCheck) nonNullWriteOnlyData() (tfprotov5.DynamicValue, error) {
	return tfprotov5.NewDynamicValue(
		r.schema().ValueType(),
		tftypes.NewValue(
			r.schema().ValueType(),
			map[string]tftypes.Value{
				"id":             tftypes.NewValue(tftypes.String, "test-123"),
				"writeonly_attr": tftypes.NewValue(tftypes.String, "this should cause an error!"),
			},
		),
	)
}

// nullWriteOnlyData is used to produce valid data.
func (r resourceWriteOnlyDataCheck) nullWriteOnlyData() (tfprotov5.DynamicValue, error) {
	return tfprotov5.NewDynamicValue(
		r.schema().ValueType(),
		tftypes.NewValue(
			r.schema().ValueType(),
			map[string]tftypes.Value{
				"id":             tftypes.NewValue(tftypes.String, "test-123"),
				"writeonly_attr": tftypes.NewValue(tftypes.String, nil),
			},
		),
	)
}

func (r resourceWriteOnlyDataCheck) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	plannedState, diag := dynamicValueToValue(r.schema(), req.PlannedState)
	if diag != nil {
		return &tfprotov5.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{diag},
		}, nil
	}

	// Destroy Op, just return planned state, which is null
	if plannedState.IsNull() {
		return &tfprotov5.ApplyResourceChangeResponse{
			NewState: req.PlannedState,
		}, nil
	}

	var newState tfprotov5.DynamicValue
	var err error
	if r.applyDataError {
		newState, err = r.nonNullWriteOnlyData()
	} else {
		newState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov5.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding new state",
					Detail:   fmt.Sprintf("Error encoding new state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov5.ApplyResourceChangeResponse{
		NewState:                    &newState,
		UnsafeToUseLegacyTypeSystem: r.enableLegacyTypeSystem,
	}, nil
}

func (r resourceWriteOnlyDataCheck) PlanResourceChange(ctx context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	proposedNewState, diag := dynamicValueToValue(r.schema(), req.ProposedNewState)
	if diag != nil {
		return &tfprotov5.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{diag},
		}, nil
	}

	// Destroying the resource, just return proposed new state (which is null)
	if proposedNewState.IsNull() {
		return &tfprotov5.PlanResourceChangeResponse{
			PlannedState: req.ProposedNewState,
		}, nil
	}

	var plannedState tfprotov5.DynamicValue
	var err error
	if r.planDataError {
		plannedState, err = r.nonNullWriteOnlyData()
	} else {
		plannedState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov5.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding planned state",
					Detail:   fmt.Sprintf("Error encoding planned state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov5.PlanResourceChangeResponse{
		PlannedState:                &plannedState,
		UnsafeToUseLegacyTypeSystem: r.enableLegacyTypeSystem,
	}, nil
}

func (r resourceWriteOnlyDataCheck) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	var newState tfprotov5.DynamicValue
	var err error
	if r.readDataError {
		newState, err = r.nonNullWriteOnlyData()
	} else {
		newState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov5.ReadResourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding new state",
					Detail:   fmt.Sprintf("Error encoding new state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov5.ReadResourceResponse{
		NewState: &newState,
	}, nil
}

func (r resourceWriteOnlyDataCheck) ValidateResourceTypeConfig(ctx context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
}

func (r resourceWriteOnlyDataCheck) ImportResourceState(ctx context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {

	var importedState tfprotov5.DynamicValue
	var err error
	if r.importDataError {
		importedState, err = r.nonNullWriteOnlyData()
	} else {
		importedState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov5.ImportResourceStateResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding import state",
					Detail:   fmt.Sprintf("Error encoding import state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov5.ImportResourceStateResponse{
		ImportedResources: []*tfprotov5.ImportedResource{
			{
				TypeName: req.TypeName,
				State:    &importedState,
			},
		},
	}, nil
}

func (r resourceWriteOnlyDataCheck) UpgradeResourceState(ctx context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	resp := &tfprotov5.UpgradeResourceStateResponse{}
	// Define options to be used when unmarshalling raw state.
	// IgnoreUndefinedAttributes will silently skip over fields in the JSON
	// that do not have a matching entry in the schema.
	unmarshalOpts := tfprotov5.UnmarshalOpts{
		ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
			IgnoreUndefinedAttributes: true,
		},
	}

	// Terraform CLI can call UpgradeResourceState even if the stored state
	// version matches the current schema. Presumably this is to account for
	// the previous terraform-plugin-sdk implementation, which handled some
	// state fixups on behalf of Terraform CLI. This will attempt to roundtrip
	// the prior RawState to a state matching the current schema.
	rawStateValue, err := req.RawState.UnmarshalWithOpts(r.schema().ValueType(), unmarshalOpts)

	if err != nil {
		diag := &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to Read Previously Saved State for UpgradeResourceState",
			Detail:   "There was an error reading the saved resource state using the current resource schema: " + err.Error(),
		}

		resp.Diagnostics = append(resp.Diagnostics, diag)

		return resp, nil //nolint:nilerr // error via diagnostic, not gRPC
	}

	upgradedState, diag := valuetoDynamicValue(r.schema(), rawStateValue)

	if diag != nil {
		resp.Diagnostics = append(resp.Diagnostics, diag)

		return resp, nil
	}

	resp.UpgradedState = upgradedState

	return resp, nil
}

func (r resourceWriteOnlyDataCheck) MoveResourceState(ctx context.Context, req *tfprotov5.MoveResourceStateRequest) (*tfprotov5.MoveResourceStateResponse, error) {
	return &tfprotov5.MoveResourceStateResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unsupported Resource Operation",
				Detail:   "MoveResourceState is not supported by this resource.",
			},
		},
	}, nil
}

func valuetoDynamicValue(schema *tfprotov5.Schema, value tftypes.Value) (*tfprotov5.DynamicValue, *tfprotov5.Diagnostic) {
	if schema == nil {
		diag := &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to Convert Value",
			Detail:   "Converting the Value to DynamicValue returned an unexpected error: missing schema",
		}

		return nil, diag
	}

	dynamicValue, err := tfprotov5.NewDynamicValue(schema.ValueType(), value)
	if err != nil {
		diag := &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to Convert Value",
			Detail:   "Converting the Value to DynamicValue returned an unexpected error: " + err.Error(),
		}

		return &dynamicValue, diag
	}

	return &dynamicValue, nil
}

func dynamicValueToValue(schema *tfprotov5.Schema, dynamicValue *tfprotov5.DynamicValue) (tftypes.Value, *tfprotov5.Diagnostic) {
	if schema == nil {
		diag := &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to Convert DynamicValue",
			Detail:   "Converting the DynamicValue to Value returned an unexpected error: missing schema",
		}

		return tftypes.NewValue(tftypes.Object{}, nil), diag
	}

	if dynamicValue == nil {
		return tftypes.NewValue(schema.ValueType(), nil), nil
	}

	value, err := dynamicValue.Unmarshal(schema.ValueType())

	if err != nil {
		diag := &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to Convert DynamicValue",
			Detail:   "Converting the DynamicValue to Value returned an unexpected error: " + err.Error(),
		}

		return value, diag
	}

	return value, nil
}
