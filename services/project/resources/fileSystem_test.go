package resources_test

import (
	"testing"

	"github.com/andygeiss/method2go/services/project"
	"github.com/andygeiss/method2go/services/project/resources"
	assert "github.com/andygeiss/utilities/testing"
)

func TestFileSystem_GenerateProjectStructure_Should_Return_Without_An_Error(t *testing.T) {
	resourceAccess := resources.NewFileSystem("testdata")
	project := &project.Project{}
	err := resourceAccess.GenerateProjectStructure(project)
	assert.That("error should be nil", t, err, nil)
	assert.That("project should not be nil", t, project == nil, false)
}
