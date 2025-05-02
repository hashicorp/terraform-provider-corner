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
			Version: 2,
			SchemaFunc: func() map[string]*schema.Schema {
				return map[string]*schema.Schema{
					// previous version of the identity (version 1)
					// "email": {
					// 	Type:              schema.TypeString,
					// 	RequiredForImport: true,
					// },
					// The second version of the identity splits the email into local part and domain
					// (for no good reason, just for one of testing of upgraders)
					"local_part": {
						Type:              schema.TypeString,
						RequiredForImport: true,
					},
					"domain": {
						Type:              schema.TypeString,
						RequiredForImport: true,
					},
				}
			},
			IdentityUpgraders: []schema.IdentityUpgrader{
				{
					Version: 1,
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

		Importer: &schema.ResourceImporter{
			// for state version 1, this could have been used:
			// StateContext: schema.ImportStatePassthroughWithIdentity("email"),
			StateContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				if rd.Id() != "" {
					return []*schema.ResourceData{rd}, nil // just return the resource data, since the string id is used
				}

				identity, err := rd.Identity()
				if err != nil {
					return nil, err
				}

				localPartRaw, ok := identity.GetOk("local_part")
				if !ok {
					return nil, fmt.Errorf("error getting local_part from identity")
				}

				localPart, ok := localPartRaw.(string)
				if !ok {
					return nil, fmt.Errorf("error converting local_part to string")
				}

				if localPart == "" {
					return nil, fmt.Errorf("local_part cannot be empty")
				}

				domainRaw, ok := identity.GetOk("domain")
				if !ok {
					return nil, fmt.Errorf("error getting domain from identity: domain value is missing or invalid")
				}
				domain, ok := domainRaw.(string)
				if !ok {
					return nil, fmt.Errorf("error converting domain to string")
				}
				if domain == "" {
					return nil, fmt.Errorf("domain cannot be empty")
				}

				email := fmt.Sprintf("%s@%s", localPart, domain)

				rd.SetId(email)
				err = rd.Set("email", email) // required for import because else it requires a replace
				if err != nil {
					return nil, fmt.Errorf("error setting email in resource data: %w", err)
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
