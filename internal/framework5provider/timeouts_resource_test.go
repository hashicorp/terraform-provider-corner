// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTimeoutsResource_unconfigured(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_timeouts" "test" {}`,
				//nolint:staticcheck //Deprecated functions
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_timeouts.test", "id", "test"),
					resource.TestCheckNoResourceAttr("framework_timeouts.test", "timeouts"),
				),
			},
		},
	})
}

func TestTimeoutsResource_configured(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_timeouts" "test" {
					timeouts {
						create = "120s"
					}
				}`,
				//nolint:staticcheck //Deprecated functions
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_timeouts.test", "id", "test"),
					resource.TestCheckResourceAttr("framework_timeouts.test", "timeouts.create", "120s"),
				),
			},
		},
	})
}
