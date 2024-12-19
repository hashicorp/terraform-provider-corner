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

func resourceWriteOnly() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWriteOnlyCreate,
		ReadContext:   resourceWriteOnlyRead,
		UpdateContext: resourceWriteOnlyUpdate,
		DeleteContext: resourceWriteOnlyDelete,

		Schema: map[string]*schema.Schema{
			"string_attr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"writeonly_bool": {
				Type:      schema.TypeBool,
				Required:  true,
				WriteOnly: true,
			},
			"writeonly_string": {
				Type:      schema.TypeString,
				Required:  true,
				WriteOnly: true,
			},
			"writeonly_int": {
				Type:      schema.TypeInt,
				Required:  true,
				WriteOnly: true,
			},
			"nested_list_block": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"string_attr": {
							Type:     schema.TypeString,
							Required: true,
						},
						"opt_or_computed_string_attr": {
							Type:     schema.TypeString,
							Default:  "computed value!",
							Optional: true,
						},
						"writeonly_string": {
							Type:      schema.TypeString,
							Required:  true,
							WriteOnly: true,
						},
					},
				},
			},
		},
	}
}

func resourceWriteOnlyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("fakeid-123")

	return verifyWriteOnlyData(d)
}

func resourceWriteOnlyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Read, so can't verify write-only data
	return nil
}

func resourceWriteOnlyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return verifyWriteOnlyData(d)
}

func resourceWriteOnlyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Config isn't set for Delete, so can't verify write-only data
	return nil
}

// verifyWriteOnlyData compares the hardcoded test data for the write-only attributes in this resource, raising
// error diagnostics if the data differs from expectations.
func verifyWriteOnlyData(d *schema.ResourceData) diag.Diagnostics {
	// Write-only string assert
	strVal, diags := d.GetRawConfigAt(cty.GetAttrPath("writeonly_string"))
	if diags.HasError() {
		return diags
	}
	if strVal.IsNull() {
		return diag.FromErr(errors.New("expected `writeonly_string` write-only attribute to be set, got: <null>"))
	}
	expectedString := "fakepassword"
	if strVal.AsString() != expectedString {
		return diag.Errorf("expected `writeonly_string` to be: %q, got: %q", expectedString, strVal.AsString())
	}
	// Setting shouldn't result in anything sent back to Terraform, but we want to test that
	// our SDKv2 logic would revert these changes.
	err := d.Set("writeonly_string", "different value")
	if err != nil {
		return diag.FromErr(err)
	}

	// Write-only bool assert
	boolVal, diags := d.GetRawConfigAt(cty.GetAttrPath("writeonly_bool"))
	if diags.HasError() {
		return diags
	}
	if boolVal.IsNull() {
		return diag.FromErr(errors.New("expected `writeonly_bool` write-only attribute to be set, got: <null>"))
	}
	if boolVal.False() {
		return diag.FromErr(errors.New("expected `writeonly_bool` to be: true, got: false"))
	}
	// Setting shouldn't result in anything sent back to Terraform, but we want to test that
	// our SDKv2 logic would revert these changes.
	err = d.Set("writeonly_bool", false)
	if err != nil {
		return diag.FromErr(err)
	}

	// Write-only int assert
	intVal, diags := d.GetRawConfigAt(cty.GetAttrPath("writeonly_int"))
	if diags.HasError() {
		return diags
	}
	if intVal.IsNull() {
		return diag.FromErr(errors.New("expected `writeonly_int` write-only attribute to be set, got: <null>"))
	}
	expectedInt := int64(1234)
	gotInt, _ := intVal.AsBigFloat().Int64()
	if gotInt != expectedInt {
		return diag.Errorf("expected `writeonly_int` to be: %d, got: %d", expectedInt, gotInt)
	}
	// Setting shouldn't result in anything sent back to Terraform, but we want to test that
	// our SDKv2 logic would revert these changes.
	err = d.Set("writeonly_int", 999)
	if err != nil {
		return diag.FromErr(err)
	}

	// Nested list block with write-only attribute
	listBlockVal, diags := d.GetRawConfigAt(cty.GetAttrPath("nested_list_block"))
	if diags.HasError() {
		return diags
	}

	if listBlockVal.IsNull() || listBlockVal.LengthInt() != 1 {
		return diag.Errorf("expected `nested_list_block` to have length of 1, got: %s", listBlockVal.GoString())
	}

	nestedWriteOnlyStr, diags := d.GetRawConfigAt(cty.GetAttrPath("nested_list_block").IndexInt(0).GetAttr("writeonly_string"))
	if diags.HasError() {
		return diags
	}
	expectedNestedWriteOnlyStr := "fakepassword"
	if nestedWriteOnlyStr.AsString() != expectedNestedWriteOnlyStr {
		return diag.Errorf("expected `nested_list_block.0.writeonly_string` to be: %s, got: %s", expectedNestedWriteOnlyStr, nestedWriteOnlyStr.AsString())
	}
	err = d.Set("nested_list_block", []map[string]any{
		{
			"string_attr":                 d.Get("nested_list_block.0.string_attr"),
			"opt_or_computed_string_attr": d.Get("nested_list_block.0.opt_or_computed_string_attr"),
			// Setting shouldn't result in anything sent back to Terraform, but we want to test that
			// our SDKv2 logic would revert these changes.
			"writeonly_string": "different value!",
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
