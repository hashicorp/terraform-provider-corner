// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2testingprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func New() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"corner_regions":     dataSourceRegions(),
			"corner_bigint":      dataSourceBigint(),
			"corner_regions_cty": dataSourceRegionsCty(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"corner_user":     resourceUser(),
			"corner_bigint":   resourceBigint(),
			"corner_user_cty": resourceUserCty(),
		},
	}

	p.ConfigureContextFunc = configure(p)

	return p
}

func configure(p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		client, err := backend.NewClient()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, nil
	}
}
