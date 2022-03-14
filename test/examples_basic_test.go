package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// You must set these environment variables for this test
const (
	vsphere_user     = "VSPHERE_USER"
	vsphere_password = "VSPHERE_PASSWORD"
	vsphere_server   = "VSPHERE_SERVER"
	// vsphere_allow_unverified_ssl = "VSPHERE_ALLOW_UNVERIFIED_SSL"
)

func GetEnvOrExit(env_var string) {
	if os.Getenv(env_var) == "" {
		fmt.Println("Environment variable: '" + env_var + "' is not set.")
		os.Exit(1)
	}
}

func TestExamplesNewCategoryNewTags(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		// Vars: map[string]interface{}{
		// 	"myvar":     "test",
		// 	"mylistvar": []string{"list_item_1"},
		// },
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportCategoryImportTag(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_tag_category_name":   "terraform",
			"vsphere_tags":                []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			"create_vsphere_tag_category": false,
			"create_vsphere_tags":         false,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportCategoryNewTag(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_tag_category_name":   "terraform",
			"vsphere_tags":                []interface{}{map[string]interface{}{"name": "project", "description": "terraform-vsphere-tags"}},
			"create_vsphere_tag_category": false,
			"create_vsphere_tags":         true,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}
