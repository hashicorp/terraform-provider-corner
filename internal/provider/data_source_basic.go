package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBasic() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBasicRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBasicRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
