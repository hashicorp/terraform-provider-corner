// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"iter"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceUserList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserListCreate,
		ReadContext:   resourceUserListRead,
		UpdateContext: resourceUserListUpdate,
		DeleteContext: resourceUserListDelete,
		ListContext:   resourceUserListList,

		ListSchemaFunc: func() map[string]*schema.Schema {

			// TF-235: this is "a hypothetical list resource schema defined in
			// SDKv2" for a `list "user"` block.
			return map[string]*schema.Schema{
				"namePrefix": {
					Type:     schema.TypeString,
					Required: true,
				},
			}

		},

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
			StateContext: schema.ImportStatePassthroughWithIdentity("email"),
		},
	}
}

func resourceUserListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)
	email := d.Get("email").(string)
	newUser := &backend.User{
		Email: email,
		Name:  d.Get("name").(string),
		Age:   d.Get("age").(int),
	}

	err := client.CreateUser(newUser)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(email)

	return resourceUserListRead(ctx, d, meta)
}

func resourceUserListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	email := d.Id()

	p, err := client.ReadUser(email)
	if err != nil {
		return diag.FromErr(err)
	}

	if p == nil {
		return nil
	}

	err = d.Set("email", email)
	if err != nil {
		return diag.FromErr(err)
	}
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

func resourceUserListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceUserListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceUserListList(ctx context.Context, d *schema.ResourceData, meta interface{}) iter.Seq2[*schema.ResourceData, diag.Diagnostics] {
	client := meta.(*backend.Client)

	// TF-235: Step 1. Decodes the list resource config, using a hypothetical list
	// resource schema defined in SDKv2.
	namePrefix := d.Get("namePrefix").(string)

	return func(push func(*schema.ResourceData, diag.Diagnostics) bool) {

		// TF-235: Step 3. Performs one or more remote API requests to retrieve
		// resources, using the configured API client.
		users, err := client.ListUsersByNamePrefix(namePrefix)
		if err != nil {
			push(nil, diag.FromErr(err))
			return
		}

		var diags diag.Diagnostics
		for _, user := range users {

			// TF-235: Step 5. Decodes each API resource into
			// ResourceData.State() and ResourceData.Identity() values,
			// using the resource schema and resource identity schema
			// defined in SDKv2.

			d.SetId(user.Email)
			if err := d.Set("email", user.Email); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}

			if err := d.Set("name", user.Name); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}

			if err := d.Set("age", user.Age); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}

			// Where to set DisplayName? Let's not worry about that right now.
			d.SetListDisplayName(user.Name + " (" + user.Email + ")")

			// Respect the return value of push by stopping iteration on false
			if !push(d, diags) {
				break
			}

		}
	}
}
