package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBasic() *schema.Resource {
	return &schema.Resource{
		Create: resourceBasicCreate,
		Read:   resourceBasicRead,
		Update: resourceBasicUpdate,
		Delete: resourceBasicDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBasicCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceBasicRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceBasicUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceBasicDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
