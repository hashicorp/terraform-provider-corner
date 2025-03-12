// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWriteOnlyValidations() *schema.Resource {
	return &schema.Resource{
		// Prevent any accidental data inconsistencies
		EnableLegacyTypeSystemPlanErrors:  true,
		EnableLegacyTypeSystemApplyErrors: true,

		CreateContext: resourceWriteOnlyValidationsCreate,
		ReadContext:   resourceWriteOnlyValidationsRead,
		UpdateContext: resourceWriteOnlyValidationsUpdate,
		DeleteContext: resourceWriteOnlyValidationsDelete,

		// TODO: The testing framework can't verify warning diagnostics currently
		// https://github.com/hashicorp/terraform-plugin-testing/issues/69
		ValidateRawResourceConfigFuncs: []schema.ValidateRawResourceConfigFunc{
			validation.PreferWriteOnlyAttribute(
				cty.GetAttrPath("old_password_attr"),
				cty.GetAttrPath("writeonly_password"),
			),
		},

		Schema: map[string]*schema.Schema{
			"old_password_attr": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ExactlyOneOf:  []string{"writeonly_password"},
				ConflictsWith: []string{"password_version"},
			},
			"password_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"writeonly_password"},
			},
			"writeonly_password": {
				Type:         schema.TypeString,
				Optional:     true,
				WriteOnly:    true,
				RequiredWith: []string{"password_version"},
			},
		},
	}
}

func resourceWriteOnlyValidationsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("fakeid-123")

	passwordVal, diags := d.GetRawConfigAt(cty.GetAttrPath("writeonly_password"))
	if diags.HasError() {
		return diags
	}

	if !passwordVal.IsNull() {
		if passwordVal.AsString() != "newpassword" && passwordVal.AsString() != "newpassword2" {
			return diag.Errorf("expected `writeonly_password` to be `newpassword` or `newpassword2`, got: %q", passwordVal.AsString())
		}

		// Setting shouldn't result in anything sent back to Terraform, but we want to test that
		// our SDKv2 logic would revert these changes.
		err := d.Set("writeonly_password", "different value")
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		passwordVal := d.Get("old_password_attr").(string)
		if passwordVal != "oldpassword" && passwordVal != "oldpassword2" {
			return diag.Errorf("expected `old_password_attr` to be `oldpassword` or `oldpassword2`, got: %q", passwordVal)
		}
	}

	return nil
}

func resourceWriteOnlyValidationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Read, so can't verify write-only data
	return nil
}

func resourceWriteOnlyValidationsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Once created, the only operation that can occur is replacement (delete/create)
	return nil
}

func resourceWriteOnlyValidationsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Delete, so can't verify write-only data
	return nil
}
