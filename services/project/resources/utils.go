package resources

import (
	"os"
	"path/filepath"
)

// CreateFolders ...
func CreateFolders(folders ...string) error {
	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}
	return nil
}

// CreateFoldersByFile ...
func CreateFoldersByFile(files ...string) error {
	for _, file := range files {
		path := filepath.Dir(file)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

// HasFile ...
func HasFile(filename string) bool {
	stat, _ := os.Stat(filename)
	fileExists := false
	if stat != nil {
		fileExists = !stat.IsDir()
	}
	return fileExists
}

// HasFolder ...
func HasFolder(name string) bool {
	stat, _ := os.Stat(name)
	dirExists := false
	if stat != nil {
		dirExists = stat.IsDir()
	}
	return dirExists
}
