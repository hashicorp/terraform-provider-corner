// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBigint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBigintCreate,
		ReadContext:   resourceBigintRead,
		UpdateContext: resourceBigintUpdate,
		DeleteContext: resourceBigintDelete,
		UseJSONNumber: true,

		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"int64": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceBigintCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	number := d.Get("number").(int)
	d.SetId(strconv.Itoa(number))

	if err := d.Set("int64", number); err != nil {
		return diag.FromErr(err)
	}
	return resourceBigintRead(ctx, d, meta)
}

func resourceBigintRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	number, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("int64", number); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceBigintUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	number := d.Get("number").(int)
	d.SetId(strconv.Itoa(number))
	if err := d.Set("int64", number); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceBigintDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
