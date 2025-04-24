package protocolv6

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceFile struct {
}

func (r resourceFile) ImportResourceState(ctx context.Context, request *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r resourceFile) MoveResourceState(ctx context.Context, request *tfprotov6.MoveResourceStateRequest) (*tfprotov6.MoveResourceStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r resourceFile) schema() *tfprotov6.Schema {
	return &tfprotov6.Schema{
		Block: &tfprotov6.SchemaBlock{
			Attributes: []*tfprotov6.SchemaAttribute{
				{
					Name:     "name",
					Type:     tftypes.String,
					Required: true,
				},
				{
					Name:     "content",
					Type:     tftypes.String,
					Computed: true,
				},
			},
		},
	}
}

func (r resourceFile) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	return &tfprotov6.ValidateResourceConfigResponse{}, nil
}

func (r resourceFile) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {
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

	return &tfprotov6.UpgradeResourceStateResponse{
		UpgradedState: &upgradeResourceState,
	}, nil
}

func (r resourceFile) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	currentState, diag := dynamicValueToValue(r.schema(), req.CurrentState)
	if diag != nil {
		return &tfprotov6.ReadResourceResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	if currentState.IsFullyKnown() && !currentState.IsNull() {
		// We can read the file content
		m := make(map[string]tftypes.Value)

		err := currentState.As(&m)
		if err != nil {
			return nil, err
		}
		filenameAttr := m["name"]
		var filename string
		if filenameAttr.Type().Is(tftypes.String) {
			filenameAttr.As(&filename)
		}

		content, err := os.ReadFile(filename)
		if err != nil {
			return &tfprotov6.ReadResourceResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error reading content",
						Detail:   fmt.Sprintf("Error reading content: %s", err.Error()),
					},
				},
			}, nil
		}

		newStateVal, err := tfprotov6.NewDynamicValue(
			r.schema().ValueType(),
			tftypes.NewValue(
				r.schema().ValueType(),
				map[string]tftypes.Value{
					"name":    tftypes.NewValue(tftypes.String, filename),
					"content": tftypes.NewValue(tftypes.String, string(content)),
				},
			),
		)
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
			NewState: &newStateVal,
		}, nil

	}

	newStateVal, _ := tfprotov6.NewDynamicValue(
		r.schema().ValueType(),
		tftypes.NewValue(
			r.schema().ValueType(),
			map[string]tftypes.Value{
				"name":    tftypes.NewValue(tftypes.String, "file.txt"),
				"content": tftypes.NewValue(tftypes.String, "test content updatedt"),
			},
		),
	)

	return &tfprotov6.ReadResourceResponse{
		NewState: &newStateVal,
	}, nil
}

func (r resourceFile) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	priorState, diag := dynamicValueToValue(r.schema(), req.PriorState)
	if diag != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	proposedNewState, diag := dynamicValueToValue(r.schema(), req.ProposedNewState)
	if diag != nil {
		return &tfprotov6.PlanResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	if proposedNewState.IsNull() {
		// destroy op
		return &tfprotov6.PlanResourceChangeResponse{
			PlannedState: req.ProposedNewState,
		}, nil
	}

	if priorState.IsNull() {
		// create op
		return &tfprotov6.PlanResourceChangeResponse{
			PlannedState: req.ProposedNewState,
		}, nil
	}

	if priorState.IsFullyKnown() && !priorState.IsNull() {
		// We can read the file content
		m := make(map[string]tftypes.Value)

		err := priorState.As(&m)
		if err != nil {
			return nil, err
		}
		filenameAttr := m["name"]
		var filename string
		if filenameAttr.Type().Is(tftypes.String) {
			filenameAttr.As(&filename)
		}

		content, err := os.ReadFile(filename)
		if err != nil {
			return &tfprotov6.PlanResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error reading content",
						Detail:   fmt.Sprintf("Error reading content: %s", err.Error()),
					},
				},
			}, nil
		}

		newStateVal, err := tfprotov6.NewDynamicValue(
			r.schema().ValueType(),
			tftypes.NewValue(
				r.schema().ValueType(),
				map[string]tftypes.Value{
					"name":    tftypes.NewValue(tftypes.String, filename),
					"content": tftypes.NewValue(tftypes.String, string(content)),
				},
			),
		)
		if err != nil {
			return &tfprotov6.PlanResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error encoding new state",
						Detail:   fmt.Sprintf("Error encoding new state: %s", err.Error()),
					},
				},
			}, nil
		}

		return &tfprotov6.PlanResourceChangeResponse{
			PlannedState: &newStateVal,
		}, nil

	}

	return &tfprotov6.PlanResourceChangeResponse{
		PlannedState: req.PriorState,
	}, nil
}

func (r resourceFile) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	priorState, diag := dynamicValueToValue(r.schema(), req.PriorState)
	if diag != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	plannedState, diag := dynamicValueToValue(r.schema(), req.PlannedState)
	if diag != nil {
		return &tfprotov6.ApplyResourceChangeResponse{
			Diagnostics: []*tfprotov6.Diagnostic{diag},
		}, nil
	}

	if plannedState.IsNull() {
		// destroy op
		if !priorState.IsNull() {
			m := make(map[string]tftypes.Value)

			err := priorState.As(&m)
			if err != nil {
				return nil, err
			}
			filenameAttr := m["name"]
			var filename string
			if filenameAttr.Type().Is(tftypes.String) {
				filenameAttr.As(&filename)
			}
			err = os.Remove(filename)
			if err != nil {
				return &tfprotov6.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov6.Diagnostic{
						{
							Severity: tfprotov6.DiagnosticSeverityError,
							Summary:  "Error deleting content",
							Detail:   fmt.Sprintf("Error deleting file: %s", err.Error()),
						},
					},
				}, nil
			}

			return &tfprotov6.ApplyResourceChangeResponse{
				NewState: req.PlannedState,
			}, nil
		}
	}

	if priorState.IsNull() {
		// create op
		m := make(map[string]tftypes.Value)

		err := plannedState.As(&m)
		if err != nil {
			return nil, err
		}
		filenameAttr := m["name"]
		var filename string
		if filenameAttr.Type().Is(tftypes.String) {
			filenameAttr.As(&filename)
		}

		contentAttr := m["content"]
		var content string
		if contentAttr.Type().Is(tftypes.String) {
			contentAttr.As(&content)
		}

		err = os.WriteFile(filename, []byte(content), 0666)
		if err != nil {
			return &tfprotov6.ApplyResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error reading content",
						Detail:   fmt.Sprintf("Error reading content: %s", err.Error()),
					},
				},
			}, nil
		}

		newStateVal, err := tfprotov6.NewDynamicValue(
			r.schema().ValueType(),
			tftypes.NewValue(
				r.schema().ValueType(),
				map[string]tftypes.Value{
					"name":    tftypes.NewValue(tftypes.String, filename),
					"content": tftypes.NewValue(tftypes.String, string(content)),
				},
			),
		)

		return &tfprotov6.ApplyResourceChangeResponse{
			NewState: &newStateVal,
		}, nil

	} else {
		// update op
		prior := make(map[string]tftypes.Value)
		planned := make(map[string]tftypes.Value)

		err := priorState.As(&prior)
		if err != nil {
			return nil, err
		}
		prev_filename := prior["name"]
		var prevFilename string
		if prev_filename.Type().Is(tftypes.String) {
			prev_filename.As(&prevFilename)
		}

		err = plannedState.As(&planned)
		if err != nil {
			return nil, err
		}
		new_filename := planned["name"]
		var newFilename string
		if new_filename.Type().Is(tftypes.String) {
			new_filename.As(&newFilename)
		}

		if newFilename != prevFilename {
			err := os.Rename(prevFilename, newFilename)
			if err != nil {
				return &tfprotov6.ApplyResourceChangeResponse{
					Diagnostics: []*tfprotov6.Diagnostic{
						{
							Severity: tfprotov6.DiagnosticSeverityError,
							Summary:  "Error renaming file",
							Detail:   fmt.Sprintf("Error renaming file: %s", err.Error()),
						},
					},
				}, nil
			}
		}

		contentAttr := planned["content"]
		var content string
		if contentAttr.Type().Is(tftypes.String) {
			contentAttr.As(&content)
		}

		err = os.WriteFile(newFilename, []byte(content), 0666)
		if err != nil {
			return &tfprotov6.ApplyResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					{
						Severity: tfprotov6.DiagnosticSeverityError,
						Summary:  "Error reading content",
						Detail:   fmt.Sprintf("Error reading content: %s", err.Error()),
					},
				},
			}, nil
		}

		newStateVal, err := tfprotov6.NewDynamicValue(
			r.schema().ValueType(),
			tftypes.NewValue(
				r.schema().ValueType(),
				map[string]tftypes.Value{
					"name":    tftypes.NewValue(tftypes.String, newFilename),
					"content": tftypes.NewValue(tftypes.String, string(content)),
				},
			),
		)

		return &tfprotov6.ApplyResourceChangeResponse{
			NewState: &newStateVal,
		}, nil
	}
}
