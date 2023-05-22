// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func dataSourceRegions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRegionsRead,

		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceRegionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	regions, err := client.ReadRegions()
	if err != nil {
		return diag.FromErr(err)
	}

	names := []string{}
	for _, r := range regions {
		names = append(names, r.Name)
	}

	d.SetId("regions")

	err = d.Set("names", names)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
