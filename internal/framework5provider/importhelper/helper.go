package importhelper

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// This helper should eventually be in terraform-plugin-framework (and an equivalent implementation in SDKv2)
func ImportStatePassthrough(ctx context.Context, attrPath path.Path, identityPath path.Path, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if attrPath.Equal(path.Empty()) {
		resp.Diagnostics.AddError(
			"Resource Import Passthrough Missing Attribute Path",
			"This is always an error in the provider. Please report the following to the provider developer:\n\n"+
				"Resource ImportState method call to ImportStatePassthrough path must be set to a valid attribute path that can accept a string value.",
		)
	}
	if identityPath.Equal(path.Empty()) {
		resp.Diagnostics.AddError(
			"Resource Import Passthrough Missing Identity Attribute Path",
			"This is always an error in the provider. Please report the following to the provider developer:\n\n"+
				"Resource ImportState method call to ImportStatePassthrough path must be set to a valid identity attribute path that is a string value.",
		)
	}

	if req.ID != "" {
		// If the import is using the ID string identifier, (either via the "terraform import" CLI command, or a config block with the "id" attribute set)
		// pass through the ID to the designated state attribute.
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, attrPath, req.ID)...)
	} else {
		// If the import is using identity, (config block with the "identity" attribute set) set the provided identity attribute ID in the import stub state.
		var identityVal types.String
		resp.Diagnostics.Append(resp.Identity.GetAttribute(ctx, identityPath, &identityVal)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, attrPath, &identityVal)...)
	}
}
