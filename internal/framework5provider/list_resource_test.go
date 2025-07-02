// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	frameworktypes "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	framework "github.com/hashicorp/terraform-provider-corner/internal/framework5provider"
)

// TODO: Model off this test ~

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

	s, ok := rawS.(tfprotov5.ProviderServerWithListResource) //nolint:staticcheck
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

	stream, err := s.ListResource(ctx, listRequest) // TODO: try to invoke this
	if err != nil {
		t.Fatalf("Failed to list resources: %v", err)
	}

	got := []string{}
	wanted := []string{"plateau", "platinum", "platypus"}
	for result := range stream.Results {
		got = append(got, result.DisplayName)

		if len(result.Diagnostics) > 0 {
			t.Errorf("expected 0 diagnostics; got: %v", result.Diagnostics)
		}
	}

	opts := cmp.Options{
		cmpopts.SortSlices(func(x, y string) bool { return x < y }),
	}
	if diff := cmp.Diff(got, wanted, opts); diff != "" {
		t.Errorf("ListResource results mismatch (-got +wanted):\n%s", diff)
	}
}
