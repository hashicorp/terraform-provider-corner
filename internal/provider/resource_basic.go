package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePerson() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePersonCreate,
		ReadContext:   resourcePersonRead,
		UpdateContext: resourcePersonUpdate,
		DeleteContext: resourcePersonDelete,

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"age": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourcePersonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	newPerson := &Person{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.CreatePerson(newPerson)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourcePersonRead(ctx, d, meta)
}

func resourcePersonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	email := d.Get("email").(string)

	p, err := client.ReadPerson(email)
	if err != nil {
		return diag.FromErr(err)
	}

	if p == nil {
		return nil
	}

	d.SetId(email)

	err = d.Set("name", p.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("age", p.Age)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourcePersonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	person := &Person{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.UpdatePerson(person)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*Client)

	person := &Person{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.DeletePerson(person)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
