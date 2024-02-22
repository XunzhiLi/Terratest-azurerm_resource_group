package test

import (
	"os/exec"
	"strings"
	"testing"

	//"os"
	// "path/filepath"
	// "fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCondition int

const (
	TestConditionEquals   TestCondition = 0
	TestConditionNotEmpty TestCondition = 1
	TestConditionContains TestCondition = 2
	apiVersion                          = "2021-11-01"
)

var (
	globalEnvVars  = make(map[string]string)
	subscriptionId string
)

const (
	fixtures = "../"
)

var (
	expected_name     = "rg-lxzdedemo-001"
	expected_location = "eastus"
)

func terraformOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: fixtures,
		Vars: map[string]interface{}{
			"resource_group_name": "rg-lxzdedemo-001",
			"location":            "eastus",
			"resource_name1":      "xiaolideprojectwokaozhendejiade ",
		},
		// VarFiles: []string{"test.tfvrs"},
	}

}

func TestTerraformAzureInfrastructure(t *testing.T) {
	t.Parallel()
	defer terraform.Destroy(t, terraformOptions())
	terraform.InitAndApply(t, terraformOptions())

	resourceGroupName := terraform.Output(t, terraformOptions(), "resource_group_name")

	require.Equal(t, "rg-lxzdedemo-001", resourceGroupName)
	assert.Equal(t, "rg-lxzdedemo-001", resourceGroupName)

	println("----------------------------------------------continue next validation")
	t.Run("Output Validation", OutputValidation)
	cmd := exec.Command("az", "group", "show", "--name", resourceGroupName, "--query", "location")
	actualLocation, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Error executing Azure CLI command: %v", err)
	}
	actualLocationString := strings.Trim(string(actualLocation), "\"\r\n")
	assert.Equal(t, expected_location, actualLocationString, "Location Mismatch")
	// avn_id := terraform.Output(t, terraformOptions(), "azurerm_virtual_network_id")
	// // fmt.Printf("avn id is %v", avn_id)

	// resp, err := getResourceFromRESTAPI(avn_id)

}

func OutputValidation(t *testing.T) {
	testCases := []struct {
		Name      string
		Got       string
		Want      string
		Condition TestCondition
	}{
		{"resource_group_name", terraform.Output(t, terraformOptions(), "resource_group_name"), expected_name, TestConditionEquals},
		{"location", terraform.Output(t, terraformOptions(), "azurerm_resource_group_location"), expected_location, TestConditionEquals}}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			switch tc.Condition {
			case TestConditionEquals:
				assert.Equal(t, tc.Got, tc.Want)
			case TestConditionNotEmpty:
				assert.NotEmpty(t, tc.Got)
			case TestConditionContains:
				assert.Contains(t, tc.Got, tc.Want)
			}

		})
	}
}
