// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"errors"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWriteOnce() *schema.Resource {
	return &schema.Resource{
		// Prevent any accidental data inconsistencies
		EnableLegacyTypeSystemPlanErrors:  true,
		EnableLegacyTypeSystemApplyErrors: true,

		CreateContext: resourceWriteOnceCreate,
		ReadContext:   resourceWriteOnceRead,
		UpdateContext: resourceWriteOnceUpdate,
		DeleteContext: resourceWriteOnceDelete,

		Schema: map[string]*schema.Schema{
			"trigger_attr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// The only place that validation can reference prior state (which is required to determine if the planned action is
			// a create, i.e, prior state is null) is during plan modification. So the customize diff implementation is responsible
			// for applying the "write-once" validation
			"writeonce_string": {
				Type:      schema.TypeString,
				Optional:  true,
				WriteOnly: true,
			},
		},
		CustomizeDiff: func(ctx context.Context, rd *schema.ResourceDiff, _ interface{}) error {
			valPath := cty.GetAttrPath("writeonce_string")

			// If there is a non-null state, we are destroying or updating so no validation is needed
			if !rd.GetRawState().IsNull() {
				return nil
			}

			configVal, diags := rd.GetRawConfigAt(valPath) // New method duplicated from (*schema.ResourceData).GetRawConfigAt
			if diags.HasError() {
				// This error shouldn't occur unless there is a schema change
				return errors.New("error retrieving config value for write-once attribute")
			}

			// We are creating, but the write-once attribute value is not present in config, return an error.
			if configVal.IsNull() {
				return errors.New(`"writeonce_string" is required when creating the resource, but no definition was found.`)
			}

			return nil
		},
	}
}

func resourceWriteOnceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("fakeid-123")

	// Write-only string is required on create, verify the data
	strVal, diags := d.GetRawConfigAt(cty.GetAttrPath("writeonce_string"))
	if diags.HasError() {
		return diags
	}

	expectedString := "fakepassword"
	if strVal.AsString() != expectedString {
		return diag.Errorf("expected `writeonce_string` to be: %q, got: %q", expectedString, strVal.AsString())
	}

	// Setting shouldn't result in anything sent back to Terraform, but we want to test that
	// our SDKv2 logic would revert these changes.
	err := d.Set("writeonce_string", "different value")
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceWriteOnceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Read, so can't verify write-only data
	return nil
}

func resourceWriteOnceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Once created, the only operation that can occur is replacement (delete/create)
	return nil
}

func resourceWriteOnceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Delete, so can't verify write-only data
	return nil
}
