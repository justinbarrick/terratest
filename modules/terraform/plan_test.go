package terraform

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/stretchr/testify/assert"
)

func TestPlanNoError(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-plan-no-error", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	options := &Options{
		TerraformDir: testFolder,
	}

	plan := InitAndPlan(t, options)

	resource := plan.GetResource("null_resource.test[1]")
	assert.NotNil(t, resource)

	attr := resource.GetAttribute("triggers.abc")
	assert.NotNil(t, attr)
	assert.Equal(t, "def", attr.NewValue)
}

func TestPlanWithError(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-plan-with-error", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	options := &Options{
		TerraformDir: testFolder,
	}

	_, err = InitAndPlanE(t, options)
	assert.Error(t, err)
}
