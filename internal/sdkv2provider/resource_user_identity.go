// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceUserIdentity() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserIdentityCreate,
		ReadContext:   resourceUserIdentityRead,
		UpdateContext: resourceUserIdentityUpdate,
		DeleteContext: resourceUserIdentityDelete,

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

		Identity: &schema.ResourceIdentity{
			Version: 1,
			SchemaFunc: func() map[string]*schema.Schema {
				return map[string]*schema.Schema{
					"email": {
						Type:              schema.TypeString,
						RequiredForImport: true,
					},
				}
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				if rd.Id() != "" {
					return []*schema.ResourceData{rd}, nil // just return the resource data, since the string id is used
				}

				identity, err := rd.Identity()
				if err != nil {
					return nil, err
				}

				emailRaw, ok := identity.GetOk("email")
				if !ok {
					return nil, fmt.Errorf("error getting email from identity: %w", err)
				}

				email, ok := emailRaw.(string)
				if !ok {
					return nil, fmt.Errorf("error converting email to string")
				}

				if email == "" {
					return nil, fmt.Errorf("email cannot be empty")
				}

				err = rd.Set("email", email)
				rd.SetId(email) // TODO: document that this is still require with resource identity
				if err != nil {
					return nil, fmt.Errorf("error setting email: %w", err)
				}

				return []*schema.ResourceData{rd}, nil
			},
		},
	}
}

func resourceUserIdentityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return resourceUserIdentityRead(ctx, d, meta)
}

func resourceUserIdentityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	identity, err := d.Identity()
	if err != nil {
		return diag.FromErr(err)
	}
	err = identity.Set("email", email)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceUserIdentityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceUserIdentityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
