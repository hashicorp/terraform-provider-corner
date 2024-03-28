// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"fmt"
	"math/big"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestNumberFunction_known(t *testing.T) {
	// Using 9223372036854775808, the smallest number that can't be represented as an int64,
	// results in an Terraform error where [Planned value does not match config value for number].
	// A value of 9223372036854775809 is used for the meanwhile.
	//
	// [Planned value does not match config value for number]: https://github.com/hashicorp/terraform/issues/34866
	bf, _, err := big.ParseFloat("9223372036854775809", 10, 512, big.ToNearestEven)

	if err != nil {
		t.Errorf("%s", err)
	}

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				output "test" {
					value = provider::framework::number(%f)
				}`, bf),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.NumberExact(bf)),
				},
			},
		},
	})
}

func TestNumberFunction_null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::framework::number(null)
				}`,
				ExpectError: regexp.MustCompile("Invalid function argument"),
			},
		},
	})
}

func TestNumberFunction_unknown(t *testing.T) {
	// Using 9223372036854775808, the smallest number that can't be represented as an int64,
	// results in an Terraform error where [Planned value does not match config value for cty.NumberIntVal],
	// which is related to a bug in go.cty relating to [Large integer comparisons and msgpack encoding].
	// A value of 9223372036854775809 is used for the meanwhile.
	//
	// [Planned value does not match config value for cty.NumberIntVal]: https://github.com/hashicorp/terraform/issues/34589
	// [Large integer comparisons and msgpack encoding]: https://github.com/zclconf/go-cty/pull/176
	bf, _, err := big.ParseFloat("9223372036854775809", 10, 512, big.ToNearestEven)

	if err != nil {
		t.Errorf("%s", err)
	}

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"framework": providerserver.NewProtocol5WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "terraform_data" "test" {
					input = provider::framework::number(%f)
				}

				output "test" {
					value = terraform_data.test.output
				}`, bf),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownOutputValue("test"),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.NumberExact(bf)),
				},
			},
		},
	})
}
