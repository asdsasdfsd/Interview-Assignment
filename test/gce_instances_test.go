package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestComputeInstance(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "./gce_instances.tf",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	instanceName := terraform.Output(t, terraformOptions, "instance_name")
	zone := terraform.Output(t, terraformOptions, "zone")
	project := terraform.Output(t, terraformOptions, "project")

	computeInstance := gcp.GetComputeInstance(t, instanceName, zone, project)
	if computeInstance == nil {
		t.Fatalf("Compute instance %s does not exist in zone %s of project %s", instanceName, zone, project)
	}

	if computeInstance.Status != "RUNNING" {
		t.Fatalf("Compute instance %s in zone %s of project %s is not running", instanceName, zone, project)
	}
}
