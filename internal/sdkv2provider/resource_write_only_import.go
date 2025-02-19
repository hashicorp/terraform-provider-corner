// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWriteOnlyImport() *schema.Resource {
	return &schema.Resource{
		// Prevent any accidental data inconsistencies
		EnableLegacyTypeSystemPlanErrors:  true,
		EnableLegacyTypeSystemApplyErrors: true,

		CreateContext: resourceWriteOnlyImportCreate,
		ReadContext:   resourceWriteOnlyImportRead,
		UpdateContext: resourceWriteOnlyImportUpdate,
		DeleteContext: resourceWriteOnlyImportDelete,

		Importer: &schema.ResourceImporter{
			StateContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
				// Setting shouldn't result in anything sent back to Terraform, but we want to test that
				// our SDKv2 logic would revert these changes.
				err := d.Set("writeonly_string", "different value")
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"string_attr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"writeonly_string": {
				Type:      schema.TypeString,
				Optional:  true,
				WriteOnly: true,
			},
		},
	}
}

func resourceWriteOnlyImportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("fakeid-123")

	return nil
}

func resourceWriteOnlyImportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := d.Set("string_attr", "hello world!")
	if err != nil {
		return diag.FromErr(err)
	}

	// Setting shouldn't result in anything sent back to Terraform, but we want to test that
	// our SDKv2 logic would revert these changes.
	err = d.Set("writeonly_string", "different value")
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceWriteOnlyImportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceWriteOnlyImportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
