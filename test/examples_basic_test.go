package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// You must set these environment variables for this test
const (
	env_var_not_set_error_message = "Environment variable: '%s' is not set."
	vsphere_user                  = "VSPHERE_USER"
	vsphere_password              = "VSPHERE_PASSWORD"
	vsphere_server                = "VSPHERE_SERVER"
	// vsphere_allow_unverified_ssl = "VSPHERE_ALLOW_UNVERIFIED_SSL"
)

func TestExamplesBasic(t *testing.T) {
	if os.Getenv(vsphere_user) == "" {
		fmt.Println("Environment variable: '" + vsphere_user + "' is not set.")
		os.Exit(1)
	}

	if os.Getenv(vsphere_password) == "" {
		fmt.Println("Environment variable: '" + vsphere_password + "' is not set.")
		os.Exit(1)
	}

	if os.Getenv(vsphere_server) == "" {
		fmt.Println("Environment variable: '" + vsphere_server + "' is not set.")
		os.Exit(1)
	}

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
