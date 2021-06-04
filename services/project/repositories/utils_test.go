package repositories_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/method2go/services/project/repositories"
	assert "github.com/andygeiss/utilities/testing"
)

func TestCreateFolders_Should_Create_One_Folder(t *testing.T) {
	os.RemoveAll("foo")
	err := repositories.CreateFolders("foo")
	assert.That("err should be nil", t, err, nil)
	assert.That("project folder exists", t, repositories.HasFolder("foo"), true)
}

func TestCreateFolders_Should_Create_Two_Folders(t *testing.T) {
	err := repositories.CreateFolders("foo", filepath.Join("foo", "bar"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, repositories.HasFolder("foo"), true)
	assert.That("folder [foo/bar] exists", t, repositories.HasFolder(filepath.Join("foo", "bar")), true)
	os.RemoveAll("foo")
}

func TestCreateFoldersByFile_Should_Create_One_Folder(t *testing.T) {
	err := repositories.CreateFoldersByFile(filepath.Join("foo", "bar.go"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, repositories.HasFolder("foo"), true)
	os.RemoveAll("foo")
}

func TestCreateFoldersByFile_Should_Create_Two_Folders(t *testing.T) {
	err := repositories.CreateFoldersByFile(filepath.Join("foo", "bar.go"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, repositories.HasFolder("foo"), true)
	assert.That("folder [foo/bar] does not exists", t, repositories.HasFolder(filepath.Join("foo", "bar.go")), false)
	os.RemoveAll("foo")
}
