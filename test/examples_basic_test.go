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
		Vars: map[string]interface{}{
			"vsphere_tag_category_name":        "example-category",
			"vsphere_tag_category_description": "",
			"vsphere_tag_category_cardinality": "MULTIPLE",
			"create_vsphere_tag_category":      true,
			"create_vsphere_tags":              true,
			"vsphere_tag_category_associable_types": []string{
				"Folder",
				"ClusterComputeResource",
				"Datacenter",
				"Datastore",
				"StoragePod",
				"DistributedVirtualPortgroup",
				"DistributedVirtualSwitch",
				"VmwareDistributedVirtualSwitch",
				"HostSystem",
				"com.vmware.content.Library",
				"com.vmware.content.library.Item",
				"HostNetwork",
				"Network",
				"OpaqueNetwork",
				"ResourcePool",
				"VirtualApp",
				"VirtualMachine",
			},
			"vsphere_tags": []interface{}{
				map[string]interface{}{
					"name":        "terraform",
					"description": "",
				},
				map[string]interface{}{
					"name":        "project",
					"description": "terraform-vsphere-tags",
				},
			},
		},
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
			"vsphere_tag_category_name":             "terraform",
			"vsphere_tag_category_description":      "",
			"vsphere_tag_category_cardinality":      "SINGLE",
			"create_vsphere_tag_category":           false,
			"create_vsphere_tags":                   false,
			"vsphere_tag_category_associable_types": []string{},
			"vsphere_tags": []interface{}{
				map[string]interface{}{
					"name":        "terraform",
					"description": "",
				},
			},
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
			"vsphere_tag_category_name":             "terraform",
			"vsphere_tag_category_description":      "",
			"vsphere_tag_category_cardinality":      "SINGLE",
			"create_vsphere_tag_category":           false,
			"create_vsphere_tags":                   true,
			"vsphere_tag_category_associable_types": []string{},
			"vsphere_tags": []interface{}{
				map[string]interface{}{
					"name":        "project",
					"description": "terraform-vsphere-tags",
				},
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}
