// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWriteOnlyUpgrade(version int) *schema.Resource {
	rSchema := &schema.Resource{
		// Prevent any accidental data inconsistencies
		EnableLegacyTypeSystemPlanErrors:  true,
		EnableLegacyTypeSystemApplyErrors: true,

		SchemaVersion: version,

		CreateContext: resourceWriteOnlyUpgradeCreate,
		ReadContext:   resourceWriteOnlyUpgradeRead,
		UpdateContext: resourceWriteOnlyUpgradeUpdate,
		DeleteContext: resourceWriteOnlyUpgradeDelete,

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

	// Avoid the internal validate error for defining a state upgrade that is equal to the schema version
	if version > 0 {
		rSchema.StateUpgraders = []schema.StateUpgrader{
			{
				Version: 0,
				Type:    rSchema.CoreConfigSchema().ImpliedType(),
				Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
					upgradedState := map[string]any{
						"id":               "fakeid-123",
						"string_attr":      "hello world!",
						"writeonly_string": "this shouldn't cause an error",
					}

					return upgradedState, nil
				},
			},
		}
	}
	return rSchema
}

func resourceWriteOnlyUpgradeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("fakeid-123")

	return nil
}

func resourceWriteOnlyUpgradeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceWriteOnlyUpgradeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceWriteOnlyUpgradeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
