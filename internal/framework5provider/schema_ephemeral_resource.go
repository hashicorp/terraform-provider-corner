package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ephemeral.EphemeralResource = ThingEphemeralResource{}

func NewThingEphemeralResource() ephemeral.EphemeralResource {
	return &ThingEphemeralResource{}
}

// ThingEphemeralResource is for testing an ephemeral resource
type ThingEphemeralResource struct{}

func (t ThingEphemeralResource) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (t ThingEphemeralResource) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"token": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type ThingEphemeralResourceModel struct {
	Name  types.String `tfsdk:"name"`
	Token types.String `tfsdk:"token"`
}

func (t ThingEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data ThingEphemeralResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Token = types.StringValue("token-abc-123-do-re-mi")

	resp.Diagnostics.Append(resp.Result.Set(ctx, data)...)
}
