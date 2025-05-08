// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceUserIdentityUpgrade(version int) *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserIdentityUpgradeCreate(version),
		ReadContext:   resourceUserIdentityUpgradeRead(version),
		UpdateContext: resourceUserIdentityUpgradeUpdate(version),
		DeleteContext: resourceUserIdentityUpgradeDelete(version),

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
			Version: int64(version),
			SchemaFunc: func() map[string]*schema.Schema {
				if version == 0 {
					return map[string]*schema.Schema{
						"email": {
							Type:              schema.TypeString,
							RequiredForImport: true,
						},
					}
				}
				if version == 1 {
					return map[string]*schema.Schema{
						"local_part": {
							Type:              schema.TypeString,
							RequiredForImport: true,
						},
						"domain": {
							Type:              schema.TypeString,
							RequiredForImport: true,
						},
					}
				}
				panic(fmt.Sprintf("unknown version %d", version))
			},
			IdentityUpgraders: []schema.IdentityUpgrader{
				{
					Version: 0,
					Type: tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"email": tftypes.String,
						},
					},
					Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
						email := rawState["email"].(string)
						parts := strings.Split(email, "@")
						if len(parts) != 2 {
							return nil, fmt.Errorf("invalid email format: %s", email)
						}
						return map[string]interface{}{
							"local_part": parts[0],
							"domain":     parts[1],
						}, nil
					},
				},
			},
		},
	}
}

func resourceUserIdentityUpgradeCreate(version int) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

		return resourceUserIdentityUpgradeRead(version)(ctx, d, meta)
	}
}

func resourceUserIdentityUpgradeRead(version int) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

		switch version {
		case 0:
			err = identity.Set("email", email)
			if err != nil {
				return diag.FromErr(err)
			}
		case 1:
			parts := strings.Split(email, "@")
			if len(parts) != 2 {
				return diag.FromErr(fmt.Errorf("invalid email format: %s", email))
			}

			err = identity.Set("local_part", parts[0])
			if err != nil {
				return diag.FromErr(err)
			}
			err = identity.Set("domain", parts[1])
			if err != nil {
				return diag.FromErr(err)
			}
		default:
			return diag.FromErr(fmt.Errorf("unknown version: %d", version))
		}

		return nil
	}
}

func resourceUserIdentityUpgradeUpdate(version int) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
}

func resourceUserIdentityUpgradeDelete(version int) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
}
