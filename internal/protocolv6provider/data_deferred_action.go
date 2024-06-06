// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package protocolv6

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type dataSourceDeferredAction struct{}

func (d dataSourceDeferredAction) ReadDataSource(ctx context.Context, req *tfprotov6.ReadDataSourceRequest) (*tfprotov6.ReadDataSourceResponse, error) {
	state, err := tfprotov6.NewDynamicValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"current": tftypes.String,
		},
	}, tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"current": tftypes.String,
		},
	}, map[string]tftypes.Value{
		"current": tftypes.NewValue(tftypes.String, time.Now().Format(time.RFC3339)),
	}))
	if err != nil {
		return &tfprotov6.ReadDataSourceResponse{
			Diagnostics: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Error encoding state",
					Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
				},
			},
		}, nil
	}
	return &tfprotov6.ReadDataSourceResponse{
		State: &state,
		Deferred: &tfprotov6.Deferred{
			Reason: tfprotov6.DeferredReasonResourceConfigUnknown,
		},
	}, nil
}

func (d dataSourceDeferredAction) ValidateDataResourceConfig(ctx context.Context, req *tfprotov6.ValidateDataResourceConfigRequest) (*tfprotov6.ValidateDataResourceConfigResponse, error) {
	return &tfprotov6.ValidateDataResourceConfigResponse{}, nil
}
