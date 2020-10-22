package warningprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"corner_warning_only": dataSourceWarningOnly(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"corner_warning_only": resourceWarningOnly(),
		},
	}

	p.ConfigureContextFunc = configure(p)

	return p
}

func warning(summary string) diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  summary,
		},
	}
}

func configure(p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return nil, warning("Warning from ConfigureContextFunc!")
	}
}

func warningResourceFunc(summary string) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
		return warning(summary)
	}
}

func dataSourceWarningOnly() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			d.SetId("static_id_value")
			return warning("Warning from Data Source ReadContext!")
		},

		Schema: map[string]*schema.Schema{},
	}
}

func resourceWarningOnly() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			d.SetId(d.Get("set_id").(string))
			return warning("Warning from Resource CreateContext!")
		},
		ReadContext: warningResourceFunc("Warning from Resource ReadContext!"),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			d.SetId(d.Get("set_id").(string))
			return warning("Warning from Resource UpdateContext!")
		},
		DeleteContext: warningResourceFunc("Warning from Resource DeleteContext!"),
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.Set("set_id", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"set_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
