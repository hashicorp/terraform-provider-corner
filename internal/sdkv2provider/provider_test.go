package sdkv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = New()
	testAccProviders = map[string]*schema.Provider{
		"corner": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := New().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = New()
}

func testAccPreCheck(t *testing.T) {
}

func TestAccTests(t *testing.T) {
	for name, c := range TestCases {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			resource.Test(t, c(t))
		})
	}
}

// public map of test cases that can be imported by Core/SDK etc
var TestCases = map[string]func(*testing.T) resource.TestCase{
	"corner_user":        testAccResourceUser,
	"corner_regions":     testAccDataSourceRegions,
	"corner_bigint_data": testAccDataSourceBigint,
	"corner_bigint":      testAccResourceBigint,
	"corner_user_cty":    testAccResourceUserCty,
	"corner_regions_cty": testAccDataSourceRegionsCty,
}
