// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider1

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func New() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"tf5muxprovider_user1": resourceUser(),
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
