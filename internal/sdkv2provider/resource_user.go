// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

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

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)
	newUser := &backend.User{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.CreateUser(newUser)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	email := d.Get("email").(string)

	p, err := client.ReadUser(email)
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

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	user := &backend.User{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.UpdateUser(user)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	user := &backend.User{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.DeleteUser(user)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
