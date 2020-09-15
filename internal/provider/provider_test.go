package provider

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
	for _, c := range TestCases {
		resource.ParallelTest(t, c(t))
	}
}

// public map of test cases that can be imported by Core/SDK etc
var TestCases = map[string]func(*testing.T) resource.TestCase{
	"basic_resource": testAccResourceBasic,
	// "basic_data_source": testAccDataSourceBasic,
}
