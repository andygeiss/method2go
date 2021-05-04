package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/method2go/services/project"
	"github.com/andygeiss/method2go/services/project/resources"
	assert "github.com/andygeiss/utilities/testing"
)

func TestFileSystem_GenerateProjectStructure_Should_Return_Without_An_Error(t *testing.T) {
	os.RemoveAll("testdata")
	path := filepath.Join("testdata", "foo", "bar")
	resourceAccess := resources.NewFileSystem(path, []string{"main.go"})
	project := &project.Project{}
	err := resourceAccess.GenerateProjectStructure(project)
	exists := resources.HasFile(filepath.Join(path, "main.go"))
	assert.That("error should be nil", t, err, nil)
	assert.That("project should not be nil", t, project == nil, false)
	assert.That("foo/bar/main.go should exists", t, exists, true)
	os.RemoveAll("testdata")
}
