package dynamic6provider_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testprovider"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/providerserver"
	"github.com/hashicorp/terraform-provider-corner/internal/testing/testsdk/resource"
)

func TestDynamicLiterals_Collections_V6(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// Terraform only has literal expressions for creating Tuple and Object values.
		// With DynamicPseudoType as the schema for an attribute, Terraform will not perform any
		// type conversion that would otherwise treat the following config as a list of strings.
		//
		// - https://developer.hashicorp.com/terraform/language/expressions/type-constraints#complex-type-literals
		// - https://developer.hashicorp.com/terraform/language/expressions/types#more-about-complex-types
		//
		// Practitioners can force Terraform to perform explicit type conversion with built-in functions like `tolist()`
		// - https://developer.hashicorp.com/terraform/language/functions/tolist
		Steps: []r.TestStep{
			{
				Config: `resource "corner_dynamic_thing" "foo" {
					dyn = ["hey", "there", "tuple"]
				}`,
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"corner": providerserver.NewProviderServer(testprovider.Provider{
				Resources: map[string]testprovider.Resource{
					"corner_dynamic_thing": {
						SchemaResponse: simpleDynamicSchema(&tfprotov6.SchemaAttribute{
							Name:     "dyn",
							Type:     tftypes.DynamicPseudoType,
							Required: true,
						}),
						PlanChangeFunc: verifyDynamicTypeByPath(
							tftypes.NewAttributePath().WithAttributeName("dyn"),
							tftypes.Tuple{
								ElementTypes: []tftypes.Type{
									tftypes.String,
									tftypes.String,
									tftypes.String,
								},
							},
						),
					},
				},
			}),
		},
	})
}

func simpleDynamicSchema(attr *tfprotov6.SchemaAttribute) *resource.SchemaResponse {
	return &resource.SchemaResponse{
		Schema: &tfprotov6.Schema{
			Block: &tfprotov6.SchemaBlock{
				Attributes: []*tfprotov6.SchemaAttribute{attr},
			},
		},
	}
}

func verifyDynamicTypeByPath(path *tftypes.AttributePath, expectedTyp tftypes.Type) func(context.Context, resource.PlanChangeRequest, *resource.PlanChangeResponse) {
	return func(ctx context.Context, req resource.PlanChangeRequest, resp *resource.PlanChangeResponse) {
		if req.Config.IsNull() {
			return
		}

		val, _, err := tftypes.WalkAttributePath(req.Config, path)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Test Verify Failed",
				Detail:    fmt.Sprintf("error finding dynamic type path: %s", err.Error()),
				Attribute: path,
			})
		}
		tftypeVal, ok := val.(tftypes.Value)
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Test Verify Failed",
				Detail:    fmt.Sprintf("error reading dynamic value, expected tftypes.Value, got %T", val),
				Attribute: path,
			})
		}

		if !tftypeVal.Type().Equal(expectedTyp) {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
				Severity:  tfprotov6.DiagnosticSeverityError,
				Summary:   "Test Verify Failed",
				Detail:    fmt.Sprintf("expected: %s, got: %s", expectedTyp, tftypeVal.Type()),
				Attribute: path,
			})
		}
	}
}
