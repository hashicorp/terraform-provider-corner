package protocol

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
)

type dataSourceTime struct{}

func (d dataSourceTime) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	log.Println("[TRACE] it's logging via println time")
	tflog.Trace(ctx, "paddyugh what time is it?", "time", "data source time!")
	tflog.Warn(ctx, "paddyugh what time is it?", "time", "data source warning time!")
	tfsdklog.Trace(ctx, "paddyugh This is an SDK-level trace log")
	tfsdklog.Warn(ctx, "paddyugh this is an SDK-level warn log")
	tfsdklog.SubsystemTrace(ctx, "proto", "paddyugh this is a protocol level trace log")
	tfsdklog.SubsystemWarn(ctx, "proto", "paddyugh this is a protocol level warn log")
	state, err := tfprotov5.NewDynamicValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"current": tftypes.String,
			"id":      tftypes.String,
		},
	}, tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"current": tftypes.String,
			"id":      tftypes.String,
		},
	}, map[string]tftypes.Value{
		"current": tftypes.NewValue(tftypes.String, time.Now().Format(time.RFC3339)),
		"id":      tftypes.NewValue(tftypes.String, "static_id"),
	}))
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error encoding state",
					Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
				},
			},
		}, nil
	}
	return &tfprotov5.ReadDataSourceResponse{
		State: &state,
	}, nil
}

func (d dataSourceTime) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	return &tfprotov5.ValidateDataSourceConfigResponse{}, nil
}
