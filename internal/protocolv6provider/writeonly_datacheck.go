// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceWriteOnlyDataCheck struct {
	resourceRouter

	enableLegacyTypeSystem   bool
	applyDataError           bool
	planDataError            bool
	readDataError            bool
	importDataError          bool
	moveResourceDataError    bool
	upgradeResourceDataError bool
}

func (r resourceWriteOnlyDataCheck) schema() *tfprotov6.Schema {
	return &tfprotov6.Schema{
		Block: &tfprotov6.SchemaBlock{
			Attributes: []*tfprotov6.SchemaAttribute{
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
func (r resourceWriteOnlyDataCheck) nonNullWriteOnlyData() (tfprotov6.DynamicValue, error) {
	return tfprotov6.NewDynamicValue(
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
func (r resourceWriteOnlyDataCheck) nullWriteOnlyData() (tfprotov6.DynamicValue, error) {
	return tfprotov6.NewDynamicValue(
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

func (r resourceWriteOnlyDataCheck) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	plannedState, diag := dynamicValueToValue(r.schema(), req.PlannedState)
	if diag != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	// Destroy Op, just return planned state, which is null
	if plannedState.IsNull() {
		return &tfprotov6.ApplyResourceChangeResponse{
			NewState: req.PlannedState,
		}, nil
	}

	var newState tfprotov6.DynamicValue
	var err error
	if r.applyDataError {
		newState, err = r.nonNullWriteOnlyData()
	} else {
		newState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding new state",
					Detail:   fmt.Sprintf("Error encoding new state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.ApplyResourceChangeResponse{
		NewState:                    &newState,
		UnsafeToUseLegacyTypeSystem: r.enableLegacyTypeSystem,
	}, nil
}

func (r resourceWriteOnlyDataCheck) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	proposedNewState, diag := dynamicValueToValue(r.schema(), req.ProposedNewState)
	if diag != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	// Destroying the resource, just return proposed new state (which is null)
	if proposedNewState.IsNull() {
		return &tfprotov6.PlanResourceChangeResponse{
			PlannedState: req.ProposedNewState,
		}, nil
	}

	var plannedState tfprotov6.DynamicValue
	var err error
	if r.planDataError {
		plannedState, err = r.nonNullWriteOnlyData()
	} else {
		plannedState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding planned state",
					Detail:   fmt.Sprintf("Error encoding planned state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.PlanResourceChangeResponse{
		PlannedState:                &plannedState,
		UnsafeToUseLegacyTypeSystem: r.enableLegacyTypeSystem,
	}, nil
}

func (r resourceWriteOnlyDataCheck) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	var newState tfprotov6.DynamicValue
	var err error
	if r.readDataError {
		newState, err = r.nonNullWriteOnlyData()
	} else {
		newState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.ReadResourceResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding new state",
					Detail:   fmt.Sprintf("Error encoding new state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.ReadResourceResponse{
		NewState: &newState,
	}, nil
}

func (r resourceWriteOnlyDataCheck) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	return &tfprotov6.ValidateResourceConfigResponse{}, nil
}

func (r resourceWriteOnlyDataCheck) ImportResourceState(ctx context.Context, req *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {

	var importedState tfprotov6.DynamicValue
	var err error
	if r.importDataError {
		importedState, err = r.nonNullWriteOnlyData()
	} else {
		importedState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.ImportResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding import state",
					Detail:   fmt.Sprintf("Error encoding import state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.ImportResourceStateResponse{
		ImportedResources: []*tfprotov6.ImportedResource{
			{
				TypeName: req.TypeName,
				State:    &importedState,
			},
		},
	}, nil
}

func (r resourceWriteOnlyDataCheck) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {
	if req.Version != 0 {
		return &tfprotov6.UpgradeResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Unsupported UpgradeResourceState Operation",
					Detail:   fmt.Sprintf(`Unexpected version upgrade, there is only version 0 of the resource. Received upgrade request with version %d`, req.Version),
				},
			},
		}, nil
	}

	var upgradeResourceState tfprotov6.DynamicValue
	var err error
	if r.upgradeResourceDataError {
		upgradeResourceState, err = r.nonNullWriteOnlyData()
	} else {
		upgradeResourceState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.UpgradeResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding upgraded state",
					Detail:   fmt.Sprintf("Error encoding upgraded state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.UpgradeResourceStateResponse{
		UpgradedState: &upgradeResourceState,
	}, nil
}

func (r resourceWriteOnlyDataCheck) MoveResourceState(ctx context.Context, req *tfprotov6.MoveResourceStateRequest) (*tfprotov6.MoveResourceStateResponse, error) {
	if req.SourceProviderAddress != "terraform.io/builtin/terraform" || req.SourceTypeName != "terraform_data" {
		return &tfprotov6.MoveResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Unsupported MoveResourceState Operation",
					Detail:   `Move operations for this resource are only supported from the "terraform.io/builtin/terraform" provider and the "terraform_data" resource type.`,
				},
			},
		}, nil
	}

	var moveResourceState tfprotov6.DynamicValue
	var err error
	if r.moveResourceDataError {
		moveResourceState, err = r.nonNullWriteOnlyData()
	} else {
		moveResourceState, err = r.nullWriteOnlyData()
	}

	if err != nil {
		return &tfprotov6.MoveResourceStateResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding moved state",
					Detail:   fmt.Sprintf("Error encoding moved state: %s", err.Error()),
				},
			},
		}, nil
	}

	return &tfprotov6.MoveResourceStateResponse{
		TargetState: &moveResourceState,
	}, nil
}

func dynamicValueToValue(schema *tfprotov6.Schema, dynamicValue *tfprotov6.DynamicValue) (tftypes.Value, *tfprotov6.Diagnostic) {
	if schema == nil {
		diag := &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
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
		diag := &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to Convert DynamicValue",
			Detail:   "Converting the DynamicValue to Value returned an unexpected error: " + err.Error(),
		}

		return value, diag
	}

	return value, nil
}
