// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//nolint:forcetypeassert // Test SDK provider
package provider1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-corner/internal/backend"
	"strings"
)

func (r UserResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
			"name": identityschema.StringAttribute{
				OptionalForImport: true,
			},
		},
	}
}

func (u UserListResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "tf5muxprovider_user1"
}

func (u UserListResource) ListResourceConfigSchema(ctx context.Context, request list.ListResourceSchemaRequest, response *list.ListResourceSchemaResponse) {
	response.Schema = listschema.Schema{
		Attributes: map[string]listschema.Attribute{
			"filter": listschema.StringAttribute{
				Required: true,
			},
		},
	}
}

type UserListResourceModel struct {
	Filter types.String `tfsdk:"filter"`
}

type UserResource struct {
	UserListResourceIdentity
	Name types.String `tfsdk:"name"`
}

type UserListResourceIdentity struct {
	ID types.String `tfsdk:"id"`
}

type UserListResource struct {
	Filter types.String `tfsdk:"filter"`
}

var identities = map[string]UserListResourceIdentity{
	"plateau":   {ID: types.StringValue("id-001")},
	"platinum":  {ID: types.StringValue("id-002")},
	"platypus":  {ID: types.StringValue("id-003")},
	"bookworm":  {ID: types.StringValue("id-004")},
	"bookshelf": {ID: types.StringValue("id-005")},
	"bookmark":  {ID: types.StringValue("id-006")},
}

var items = map[string]UserResource{
	"plateau":   {UserListResourceIdentity: identities["plateau"], Name: types.StringValue("plateau")},
	"platinum":  {UserListResourceIdentity: identities["platinum"], Name: types.StringValue("platinum")},
	"platypus":  {UserListResourceIdentity: identities["platypus"], Name: types.StringValue("platypus")},
	"bookworm":  {UserListResourceIdentity: identities["bookworm"], Name: types.StringValue("bookworm")},
	"bookshelf": {UserListResourceIdentity: identities["bookshelf"], Name: types.StringValue("bookshelf")},
	"bookmark":  {UserListResourceIdentity: identities["bookmark"], Name: types.StringValue("bookmark")},
}

func (u UserListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data UserListResourceModel

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	stream.Results = func(push func(list.ListResult) bool) {
		for name, item := range items {
			if !strings.HasPrefix(name, data.Filter.ValueString()) {
				continue
			}

			result := req.NewListResult()
			result.DisplayName = item.Name.ValueString()

			if diags := result.Resource.Set(ctx, item); diags.HasError() {
				result.Diagnostics.Append(diags...)
			}

			if diags := result.Identity.Set(ctx, identities[name]); diags.HasError() {
				result.Diagnostics.Append(diags...)
			}

			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
			}

			if !push(result) {
				return
			}
		}
	}
}

var _ list.ListResource = UserListResource{}

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

		// TODO add a resource identity for this
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
