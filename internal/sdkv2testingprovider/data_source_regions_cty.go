// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2testingprovider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func dataSourceRegionsCty() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRegionsCtyRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceRegionsCtyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	config := d.GetRawConfig().AsValueMap()
	filter := config["filter"].AsString()
	if filter == "" {
		return diag.FromErr(errors.New("filter wasn't set"))
	}

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
