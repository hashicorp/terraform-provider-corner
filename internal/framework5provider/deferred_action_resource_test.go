// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// TODO: test provider deferral with modify plan enabled
func TestDeferredActionResource_ProviderDeferral(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan:  resource.PlanOptions{AllowDeferral: true},
			Apply: resource.ApplyOptions{AllowDeferral: true},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `provider "framework" {
					deferral = true
				}

				resource "framework_deferred_action" "test" {
					id = "test_id"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.test", plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
			},
			{
				Config: `provider "framework" {
					deferral = true
				}

				resource "framework_deferred_action" "test" {
					id = "test_id"
					modify_plan_deferral = true
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.test", plancheck.DeferredReasonResourceConfigUnknown),
					},
				},
			},
			{
				Config: `resource "framework_deferred_action" "test" {
					id = "test_id"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_deferred_action.test", tfjsonpath.New("id"), knownvalue.StringExact("test_id")),
				},
			},
		},
	})
}

func TestDeferredActionResource_ModifyPlanDeferral(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan:  resource.PlanOptions{AllowDeferral: true},
			Apply: resource.ApplyOptions{AllowDeferral: true},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_deferred_action" "test" {
					id = "test_id"
					modify_plan_deferral = true
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.test", plancheck.DeferredReasonResourceConfigUnknown),
					},
				},
			},
			{
				Config: `resource "framework_deferred_action" "test" {
					id = "test_id"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_deferred_action.test", tfjsonpath.New("id"), knownvalue.StringExact("test_id")),
				},
			},
		},
	})
}

func TestDeferredActionResource_ReadDeferral(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan:  resource.PlanOptions{AllowDeferral: true},
			Apply: resource.ApplyOptions{AllowDeferral: true},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_deferred_action" "test" {
					id = "test_id"
					read_deferral = true
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.test", plancheck.DeferredReasonResourceConfigUnknown),
					},
				},
			},
		},
	})
}

func TestDeferredActionResource_ImportStateDeferral(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan:  resource.PlanOptions{AllowDeferral: true},
			Apply: resource.ApplyOptions{AllowDeferral: true},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `import {
					to = framework_deferred_action.import_test
					id = "test-id"
				}
				resource "framework_deferred_action" "import_test" {}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.import_test", plancheck.DeferredReasonResourceConfigUnknown),
					},
				},
			},
		},
	})
}
