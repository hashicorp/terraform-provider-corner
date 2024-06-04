// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

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
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `provider "framework" {
					deferral = true
				}

				resource "framework_deferred_action" "test" {
					id = "test"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action.test",
							plancheck.DeferredReasonProviderConfigUnknown),
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
						plancheck.ExpectDeferredChange("framework_deferred_action.test",
							plancheck.DeferredReasonProviderConfigUnknown),
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
					statecheck.ExpectKnownValue("framework_deferred_action.test",
						tfjsonpath.New("id"), knownvalue.StringExact("test_id")),
				},
			},
		},
	})
}

func TestDeferredActionPlanModificationResource_ProviderDeferral(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan:  resource.PlanOptions{AllowDeferral: true},
			Apply: resource.ApplyOptions{AllowDeferral: true},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `provider "framework" {
					deferral = true
				}

				resource "framework_deferred_action_plan_modification" "test" {
					id = "test"
				}`,
				ExpectError: regexp.MustCompile("Error: invalid id value"),
			},
			{
				Config: `provider "framework" {
					deferral = true
				}
			
				resource "framework_deferred_action_plan_modification"  "test" {
					id = "test_id"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action_plan_modification.test",
							plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
			},
			{
				Config: `provider "framework" {
					deferral = true
				}
			
				resource "framework_deferred_action_plan_modification"  "test" {
					id = "test_id"
					modify_plan_deferral = true
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredChange("framework_deferred_action_plan_modification.test",
							plancheck.DeferredReasonResourceConfigUnknown),
					},
				},
			},
			{
				Config: `resource "framework_deferred_action_plan_modification" "test" {
					id = "test_id"
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNoDeferredChanges(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("framework_deferred_action_plan_modification.test",
						tfjsonpath.New("id"), knownvalue.StringExact("test_id")),
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
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
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
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
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
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
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
