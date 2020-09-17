package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"corner_data_source": dataSourceBasic(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"corner_person": resourcePerson(),
		},
	}

	p.ConfigureContextFunc = configure(p)

	return p
}

func configure(p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		client, err := NewClient()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, nil
	}
}

// each test has its own db
// db address set on configure
