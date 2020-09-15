package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"basic_data_source": dataSourceBasic(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"basic_resource": resourceBasic(),
		},
	}

	p.ConfigureContextFunc = configure(p)

	return p
}

type apiClient struct {
}

func configure(p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return &apiClient{}, nil
	}
}
