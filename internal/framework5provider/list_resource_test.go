package framework_test

import (
	"fmt"
	"testing"

	frameworktypes "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	framework "github.com/hashicorp/terraform-provider-corner/internal/framework5provider"
)

func TestListResource(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	rawProvider := framework.New()
	provider, ok := rawProvider.(frameworktypes.ProviderWithListResources)
	if !ok {
		t.Fatalf("Expected provider to implement ProviderWithListResources, got %T", rawProvider)
	}

	rawS, err := providerserver.NewProtocol5WithError(provider)()
	if err != nil {
		t.Fatalf("Failed to create provider server: %v", err)
	}

	s, ok := rawS.(tfprotov5.ProviderServerWithListResource)
	if !ok {
		t.Fatalf("Expected server to implement ProviderServerWithListResource, got %T", rawS)
	}

	config, err := tfprotov5.NewDynamicValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"filter": tftypes.String,
			},
		},
		tftypes.NewValue(
			tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"filter": tftypes.String,
				},
			},
			map[string]tftypes.Value{
				"filter": tftypes.NewValue(tftypes.String, "plat"),
			},
		),
	)
	if err != nil {
		t.Fatalf("Failed to create dynamic value: %v", err)
	}
	listRequest := &tfprotov5.ListResourceRequest{
		TypeName: "framework_list_resource",
		Config:   &config,
	}
	stream, err := s.ListResource(ctx, listRequest)
	if err != nil {
		t.Fatalf("Failed to list resources: %v", err)
	}
	for result := range stream.Results {
		fmt.Printf("ListResource result: %v\n", result.DisplayName)
		for _, diag := range result.Diagnostics {
			fmt.Printf("Diagnostic: %s %s\n", diag.Summary, diag.Detail)
		}
	}
}
