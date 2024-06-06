// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdkv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"deferral": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"corner_regions":     dataSourceRegions(),
			"corner_bigint":      dataSourceBigint(),
			"corner_regions_cty": dataSourceRegionsCty(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"corner_user":            resourceUser(),
			"corner_bigint":          resourceBigint(),
			"corner_user_cty":        resourceUserCty(),
			"corner_deferred_action": resourceDeferredAction(),
			"corner_deferred_action_plan_modification": resourceDeferredActionPlanModification(),
		},
	}

	p.ConfigureProvider = func(ctx context.Context, req schema.ConfigureProviderRequest, resp *schema.ConfigureProviderResponse) {
		client, err := backend.NewClient()
		if err != nil {
			resp.Diagnostics = diag.FromErr(err)
		}

		if req.ResourceData.Get("deferral") == true && req.DeferralAllowed {
			resp.Deferred = &schema.Deferred{
				Reason: schema.DeferredReasonProviderConfigUnknown,
			}
		}
		resp.Meta = client
	}

	return p
}
