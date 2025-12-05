// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestModifyFileAction(t *testing.T) {
	f := filepath.Join(t.TempDir(), "local_file")
	f = strings.ReplaceAll(f, `\`, `\\`)

	content := "test data"
	updatedContent := "updated test data"

	resource.UnitTest(t, resource.TestCase{

		// Unlinked Actions are only available in 1.14.0+
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.14.0"))),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"framework": providerserver.NewProtocol6WithError(New()),
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "terraform_data" "test" {
					input = "fake-string"

					lifecycle {
						action_trigger {
						  events  = [before_create]
						  actions = [action.framework_modify_file_action.file]
						}
					}
				}

				action "framework_modify_file_action" "file" {
					config {
						filename = %[1]q
						content  = %[2]q
					}
				  
				}`, f, content),
				Check: func(state *terraform.State) error {
					resultContent, err := os.ReadFile(f)
					if err != nil {
						return fmt.Errorf("Error occurred while reading file at path: %s\n, error: %s\n", f, err)
					}

					if string(resultContent) != content {
						return fmt.Errorf("Expected file content %s\n, got: %s\n", content, resultContent)
					}
					return nil
				},
			},
			// Test that changing the action configuration by itself doesn't invoke the action
			{
				Config: fmt.Sprintf(`
				resource "terraform_data" "test" {
					input = "fake-string"

					lifecycle {
						action_trigger {
						  events  = [before_create]
						  actions = [action.framework_modify_file_action.file]
						}
					}
				}

				action "framework_modify_file_action" "file" {
					config {
						filename = %[1]q
						content  = "updated test data"
					}
				  
				}`, f),
				Check: func(state *terraform.State) error {
					resultContent, err := os.ReadFile(f)
					if err != nil {
						return fmt.Errorf("Error occurred while reading file at path: %s\n, error: %s\n", f, err)
					}

					if string(resultContent) != content {
						return fmt.Errorf("Expected file content %s\n, got: %s\n", content, resultContent)
					}
					return nil
				},
			},
			// test an 'after_update' event
			{
				Config: fmt.Sprintf(`
				resource "terraform_data" "test" {
					input = "updated-fake-string" # trigger an update
			
					lifecycle {
						action_trigger {
						  events  = [after_update] # action triggers after update
						  actions = [action.framework_modify_file_action.file]
						}
					}
				}

				action "framework_modify_file_action" "file" {
					config {
						filename = %[1]q
						content  = "updated test data"
					}
				  
				}`, f),
				Check: func(state *terraform.State) error {
					resultContent, err := os.ReadFile(f)
					if err != nil {
						return fmt.Errorf("Error occurred while reading file at path: %s\n, error: %s\n", f, err)
					}

					if string(resultContent) != updatedContent {
						return fmt.Errorf("Expected file content %s\n, got: %s\n", updatedContent, resultContent)
					}
					return nil
				},
			},
			// Test Plan Modification
			{
				Config: fmt.Sprintf(`
				resource "terraform_data" "test" {
					input = "fake-strings" # trigger an update
			
					lifecycle {
						action_trigger {
						  events  = [after_update]
						  actions = [action.framework_modify_file_action.file]
						}
					}
				}
			
				action "framework_modify_file_action" "file" {
					config {
						filename = %[1]q
						content  = %[2]q
						plan_error = true
					}
			
				}`, f, content),
				ExpectError: regexp.MustCompile("ModifyPlan error"),
				// Assert that the file remains unchanged
				Check: func(state *terraform.State) error {
					resultContent, err := os.ReadFile(f)
					if err != nil {
						return fmt.Errorf("Error occurred while reading file at path: %s\n, error: %s\n", f, err)
					}

					if string(resultContent) != updatedContent {
						return fmt.Errorf("Expected file content %s\n, got: %s\n", updatedContent, resultContent)
					}
					return nil
				},
			},
		},
	})
}
