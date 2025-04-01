package protocolv6

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ tfprotov6.ActionServer = actionFile{}

type actionFile struct {
}

func (a actionFile) schema() *tfprotov6.ActionSchema {
	return &tfprotov6.ActionSchema{
		LinkedResources: map[string]*tfprotov6.LinkedResource{
			"cty.GetAttrPath(\"object\")": {
				TypeName: "corner_v6_file",
			},
		},
		Block: &tfprotov6.SchemaBlock{
			Attributes: []*tfprotov6.SchemaAttribute{
				{
					Name: "object",
					Type: tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					},
					Required: true,
				},
				{
					Name:     "content",
					Type:     tftypes.String,
					Required: true,
				},
			},
		},
	}
}

func (a actionFile) PlanAction(ctx context.Context, req *tfprotov6.PlanActionRequest) (*tfprotov6.PlanActionResponse, error) {
	return &tfprotov6.PlanActionResponse{}, nil
}

func (a actionFile) InvokeAction(ctx context.Context, req *tfprotov6.InvokeActionRequest, resp *tfprotov6.InvokeActionResponse) error {
	config, diag := actionDynamicValueToValue(a.schema(), req.Config)
	if diag != nil {
		resp.Diagnostics = []*tfprotov6.Diagnostic{diag}
		return nil
	}

	if !config.IsFullyKnown() {
		resp.Diagnostics = []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "config is not wholly known",
				Detail:   "currently not supported, might cause a 'deferred action'-action",
			},
		}
		return nil
	}

	actionConfig := make(map[string]tftypes.Value)

	err := config.As(&actionConfig)
	if err != nil {
		return err
	}

	obj := actionConfig["object"]
	if obj.IsNull() {
		resp.Diagnostics = []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "object not found",
				Detail:   "object is in fact required",
			},
		}
		return nil
	}

	object := make(map[string]tftypes.Value)
	err = obj.As(&object)
	if err != nil {
		return err
	}

	contentAttr := actionConfig["content"]
	if obj.IsNull() {
		resp.Diagnostics = []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "object not found",
				Detail:   "content is in fact required",
			},
		}
		return nil
	}

	if !contentAttr.Type().Is(tftypes.String) {
		resp.Diagnostics = []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "content should be a string",
				Detail:   "and its not",
			},
		}
		return nil
	}
	var content string
	contentAttr.As(&content)

	err = resp.CallbackServer.Send(ctx, &tfprotov6.ProgressActionEvent{
		StdOut: []string{"Wohooo, we got an update"},
		StdErr: []string{"And an error"},
	})
	if err != nil {
		return err
	}

	filenameAttr := object["name"]
	if filenameAttr.IsNull() {
		resp.Diagnostics = []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "filename not found",
				Detail:   "filename is in fact required",
			},
		}
		return nil
	}

	var filename string
	if filenameAttr.Type().Is(tftypes.String) {
		filenameAttr.As(&filename)
	}

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		err = resp.CallbackServer.Send(ctx, &tfprotov6.DiagnosticsActionEvent{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "filename not found",
					Detail:   "filename is in fact required",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	}

	newFileContent := string(fileContent) + "\n" + content

	err = os.WriteFile(filename, []byte(newFileContent), 0644)
	if err != nil {

		err = resp.CallbackServer.Send(ctx, &tfprotov6.DiagnosticsActionEvent{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Could not write file for backup action",
					Detail:   fmt.Sprintf("Could not write file '%s': %s", filename, err),
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	}

	err = resp.CallbackServer.Send(ctx, &tfprotov6.ProgressActionEvent{
		StdOut: []string{"Sleeping for 2 minutes..."},
	})
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Minute)
	err = resp.CallbackServer.Send(ctx, &tfprotov6.ProgressActionEvent{
		StdOut: []string{"Aaaaaand done"},
	})
	if err != nil {
		return err
	}

	err = resp.CallbackServer.Send(ctx, &tfprotov6.FinishedActionEvent{})
	if err != nil {
		return err
	}

	return nil
}

func (a actionFile) CancelAction(ctx context.Context, request *tfprotov6.CancelActionRequest) (*tfprotov6.CancelActionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func actionDynamicValueToValue(schema *tfprotov6.ActionSchema, dynamicValue *tfprotov6.DynamicValue) (tftypes.Value, *tfprotov6.Diagnostic) {
	if schema == nil {
		diag := &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Unable to Convert DynamicValue",
			Detail:   "Converting the DynamicValue to Value returned an unexpected error: missing schema",
		}

		return tftypes.NewValue(tftypes.Object{}, nil), diag
	}

	if dynamicValue == nil {
		return tftypes.NewValue(schema.Block.ValueType(), nil), nil
	}

	value, err := dynamicValue.Unmarshal(schema.Block.ValueType())

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
