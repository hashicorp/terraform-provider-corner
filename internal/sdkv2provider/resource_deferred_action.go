// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceDeferredAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeferredActionCreate,
		ReadContext:   resourceDeferredActionRead,
		UpdateContext: resourceDeferredActionUpdate,
		DeleteContext: resourceDeferredActionDelete,

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

		CustomizeDiff: customdiff.All(
			customdiff.ValidateChange("age", func(ctx context.Context, oldVal, newVal, meta any) error {
				if newVal.(int) > 100 { //nolint
					return fmt.Errorf("age value must be less than 100")
				}
				return nil
			}),
		),
	}
}

func resourceDeferredActionPlanModification() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeferredActionCreate,
		ReadContext:   resourceDeferredActionRead,
		UpdateContext: resourceDeferredActionUpdate,
		DeleteContext: resourceDeferredActionDelete,

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

		CustomizeDiff: customdiff.All(
			customdiff.ValidateChange("age", func(ctx context.Context, oldVal, newVal, meta any) error {
				if newVal.(int) > 100 { //nolint
					return fmt.Errorf("age value must be less than 100")
				}
				return nil
			}),
		),

		// This allows CustomizeDiff to be called even if a deferral response will be returned.
		ResourceBehavior: schema.ResourceBehavior{
			ProviderDeferred: schema.ProviderDeferredBehavior{
				EnablePlanModification: true,
			},
		},
	}
}

func resourceDeferredActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceDeferredActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceDeferredActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceDeferredActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
