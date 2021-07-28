package resources_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/method2go/resources"
	"github.com/andygeiss/method2go/services/project"
	assert "github.com/andygeiss/utilities/testing"
)

func TestFileSystem_GenerateProjectStructure_Should_Return_Without_An_Error(t *testing.T) {
	os.RemoveAll("testdata")
	access := resources.NewProjectResourceAccessFileSystem("testdata", []string{"foo/bar/foo.txt"})
	project := &project.Project{
		Contents: map[string]string{
			"foo/bar/foo.txt": "bar",
		},
	}
	err := access.GenerateProjectStructure(project)
	exists := resources.HasFile(filepath.Join("testdata", "foo", "bar", "foo.txt"))
	assert.That("error should be nil", t, err, nil)
	assert.That("project should not be nil", t, project == nil, false)
	assert.That("foo/bar/main.go should exists", t, exists, true)
	os.RemoveAll("testdata")
}
