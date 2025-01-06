package cornertesting

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func PreferWriteOnlyAttribute(oldAttribute, newAttribute path.Expression) resource.ConfigValidator {
	return preferWriteOnlyAttributeValidator{
		oldAttribute: oldAttribute,
		newAttribute: newAttribute,
	}
}

var _ resource.ConfigValidator = &preferWriteOnlyAttributeValidator{}

type preferWriteOnlyAttributeValidator struct {
	oldAttribute path.Expression
	newAttribute path.Expression
}

func (v preferWriteOnlyAttributeValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v preferWriteOnlyAttributeValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("If the client supports write-only attributes (Terraform v1.11+), attribute %s should be used in-place of %s", v.newAttribute, v.oldAttribute)
}

func (v preferWriteOnlyAttributeValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// Write-only attributes are not supported in the client, so no warning should be raised.
	if !req.ClientCapabilities.WriteOnlyAttributesAllowed {
		return
	}

	matchedOldPaths, matchedOldPathsDiags := req.Config.PathMatches(ctx, v.oldAttribute)
	resp.Diagnostics.Append(matchedOldPathsDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var diags diag.Diagnostics
	for _, matchedOldPath := range matchedOldPaths {
		var value attr.Value
		getAttributeDiags := req.Config.GetAttribute(ctx, matchedOldPath, &value)

		diags.Append(getAttributeDiags...)

		// Collect all errors
		if getAttributeDiags.HasError() {
			continue
		}

		// Value must not be null or unknown to trigger validation error
		if value.IsNull() || value.IsUnknown() {
			continue
		}

		diags.AddAttributeWarning(
			matchedOldPath,
			"Available Write-Only Attribute Alternative",
			fmt.Sprintf("The attribute %s has a WriteOnly version %s available. "+
				"Use the WriteOnly version of the attribute when possible.", matchedOldPath, v.newAttribute),
		)
	}

	resp.Diagnostics = diags
}
