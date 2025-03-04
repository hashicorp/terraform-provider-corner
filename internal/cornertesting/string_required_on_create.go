// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package cornertesting

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func RequiredOnCreate() planmodifier.String {
	return requiredOnCreateModifier{}
}

type requiredOnCreateModifier struct{}

func (m requiredOnCreateModifier) Description(_ context.Context) string {
	return "This attribute is required only when creating the resource."
}

func (m requiredOnCreateModifier) MarkdownDescription(_ context.Context) string {
	return "This attribute is required only when creating the resource."
}

func (m requiredOnCreateModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// If there is a non-null state, we are destroying or updating so no validation is needed
	if !req.State.Raw.IsNull() {
		return
	}

	// We are creating, but the attribute value is not present in config, return an error.
	if req.ConfigValue.IsNull() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Attribute Required when Creating",
			fmt.Sprintf("Must set a configuration value for the %s attribute when creating.", req.Path.String()),
		)
		return
	}
}
