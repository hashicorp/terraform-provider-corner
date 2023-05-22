// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package sdkv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
)

func resourceUserCty() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCtyCreate,
		ReadContext:   resourceUserCtyRead,
		UpdateContext: resourceUserCtyUpdate,
		DeleteContext: resourceUserCtyDelete,
		CustomizeDiff: resourceUserCtyCustomizeDiff,

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

func resourceUserCtyCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if d.Get("email").(string) != "" && d.GetRawConfig().IsNull() {
		return errors.New("raw config not set in plan")
	}
	if d.Id() != "" && d.GetRawState().IsNull() {
		return errors.New("raw state not set in plan")
	}
	if d.GetRawPlan().IsNull() {
		return errors.New("raw plan not set in plan")
	}
	return nil
}

func resourceUserCtyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)
	config := d.GetRawConfig().AsValueMap()
	age, _ := config["age"].AsBigFloat().Int64()
	newUser := &backend.User{
		Email: config["email"].AsString(),
		Name:  config["name"].AsString(),
		Age:   int(age),
	}

	// plan should be set
	if d.GetRawPlan().IsNull() {
		return diag.FromErr(errors.New("plan wasn't set"))
	}

	// state should not be set
	if !d.GetRawState().IsNull() {
		return diag.FromErr(fmt.Errorf("state was %s, not null", d.GetRawState().GoString()))
	}

	err := client.CreateUser(newUser)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newUser.Email)

	err = d.Set("name", newUser.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("age", newUser.Age)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceUserCtyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	// state should be set
	state := d.GetRawState().AsValueMap()
	email := state["email"].AsString()

	// plan should not be set
	if !d.GetRawPlan().IsNull() {
		return diag.FromErr(fmt.Errorf("plan was %s, not null", d.GetRawPlan().GoString()))
	}

	// config should not be set
	if !d.GetRawConfig().IsNull() {
		return diag.FromErr(fmt.Errorf("config was %s, not null", d.GetRawConfig().GoString()))
	}

	p, err := client.ReadUser(email)
	if err != nil {
		return diag.FromErr(err)
	}

	if p == nil {
		d.SetId("")
		return nil
	}

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

func resourceUserCtyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	config := d.GetRawConfig().AsValueMap()
	age, _ := config["age"].AsBigFloat().Int64()
	user := &backend.User{
		Email: config["email"].AsString(),
		Name:  config["name"].AsString(),
		Age:   int(age),
	}

	// plan should be set
	if d.GetRawPlan().IsNull() {
		return diag.FromErr(errors.New("plan wasn't set"))
	}

	// state should be set
	if d.GetRawState().IsNull() {
		return diag.FromErr(errors.New("state wasn't set"))
	}

	err := client.UpdateUser(user)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceUserCtyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*backend.Client)

	state := d.GetRawState().AsValueMap()
	age, _ := state["age"].AsBigFloat().Int64()
	user := &backend.User{
		Email: state["email"].AsString(),
		Name:  state["name"].AsString(),
		Age:   int(age),
	}

	// plan should be null
	if !d.GetRawPlan().IsNull() {
		return diag.FromErr(fmt.Errorf("plan was set to %s", d.GetRawPlan().GoString()))
	}

	// config should be null
	if !d.GetRawConfig().IsNull() {
		return diag.FromErr(fmt.Errorf("config was set to %s", d.GetRawConfig().GoString()))
	}

	err := client.DeleteUser(user)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
