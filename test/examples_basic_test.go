package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

const (
	user     = "VSPHERE_USER"
	password = "VSPHERE_PASSWORD"
	server   = "VSPHERE_SERVER"
	// allow_unverified_ssl = "VSPHERE_ALLOW_UNVERIFIED_SSL"

	existing_category_name = "VSPHERE_EXISTING_CATEGORY_NAME"
	existing_tag_name      = "VSPHERE_EXISTING_TAG_NAME"

	input_validation_test_failed_message = "Invalid '%s' value input validation test failed."
)

func GetEnvOrExit(env_var string) {
	if os.Getenv(env_var) == "" {
		fmt.Println("Environment variable: '" + env_var + "' is not set.")
		os.Exit(1)
	}
}

func GetEnvsOrExit() {
	GetEnvOrExit(user)
	GetEnvOrExit(password)
	GetEnvOrExit(server)
}

func TestExamplesNewCategoryNewTags(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		// Vars: map[string]interface{}{
		// 	"tag_category_name":        "example-category",
		// 	"tag_category_description": "",
		// 	"tag_category_cardinality": "MULTIPLE",
		// 	"create_tag_category":      true,
		// 	"create_tags":              true,
		// 	"tag_category_associable_types": []string{
		// 		"Folder",
		// 		"ClusterComputeResource",
		// 		"Datacenter",
		// 		"Datastore",
		// 		"StoragePod",
		// 		"DistributedVirtualPortgroup",
		// 		"DistributedVirtualSwitch",
		// 		"VmwareDistributedVirtualSwitch",
		// 		"HostSystem",
		// 		"com.vmware.content.Library",
		// 		"com.vmware.content.library.Item",
		// 		"HostNetwork",
		// 		"Network",
		// 		"OpaqueNetwork",
		// 		"ResourcePool",
		// 		"VirtualApp",
		// 		"VirtualMachine",
		// 	},
		// 	"tags": map[string]interface{}{
		// 		"terraform": "",
		// 		"project": "terraform-vsphere-tags",
		// 	},
		// },
		Vars: map[string]interface{}{
			"create_tag_category": true,
			"create_tags":         true,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportCategoryImportTag(t *testing.T) {
	GetEnvsOrExit()
	GetEnvOrExit(existing_category_name)
	GetEnvOrExit(existing_tag_name)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_name":   os.Getenv(existing_category_name),
			"create_tag_category": false,
			"create_tags":         false,
			"tags": map[string]interface{}{
				os.Getenv(existing_tag_name): "",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportCategoryNewTag(t *testing.T) {
	GetEnvsOrExit()
	GetEnvOrExit(existing_category_name)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_name":   os.Getenv(existing_category_name),
			"create_tag_category": false,
			"create_tags":         true,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesNewCategoryImportTag(t *testing.T) {
	// Create category and skip importing tags since they don't exist. Intentionally passes.
	GetEnvsOrExit()
	GetEnvOrExit(existing_tag_name)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"create_tag_category": true,
			"create_tags":         false,
			"tags": map[string]interface{}{
				os.Getenv(existing_tag_name): "",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesInvalidCategoryName(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "tag_category_name"))
	}
}

func TestExamplesInvalidCategoryCardinality(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_cardinality": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "tag_category_cardinality"))
	}
}

func TestExamplesInvalidCategoryAssociableTypesLength(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_associable_types": []string{},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "tag_category_associable_types"))
	}
}

func TestExamplesInvalidCategoryAssociableTypesValue(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tag_category_associable_types": []string{"test"},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "tag_category_associable_types"))
	}
}

func TestExamplesInvalidTagLength(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"tags": map[string]interface{}{},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "tags"))
	}
}
