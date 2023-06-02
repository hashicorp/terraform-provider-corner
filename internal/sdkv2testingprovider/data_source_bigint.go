// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2testingprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBigint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBigintRead,

		Schema: map[string]*schema.Schema{
			"int64": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceBigintRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("7227701560655103598")

	if err := d.Set("int64", 7227701560655103598); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
