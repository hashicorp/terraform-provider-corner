package protocolv6

import (
	"context"

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

func (a actionFile) InvokeAction(ctx context.Context, request *tfprotov6.InvokeActionRequest, response *tfprotov6.InvokeActionResponse) error {
	events := response.Events

	events <- &tfprotov6.ProgressActionEvent{
		StdOut: []string{"Wohooo, we got an update"},
		StdErr: []string{"And an error"},
	}

	events <- &tfprotov6.FinishedActionEvent{}

	return nil
}

func (a actionFile) CancelAction(ctx context.Context, request *tfprotov6.CancelActionRequest) (*tfprotov6.CancelActionResponse, error) {
	//TODO implement me
	panic("implement me")
}
