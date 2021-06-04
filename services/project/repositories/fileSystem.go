package repositories

import (
	"os"
	"path/filepath"

	"github.com/andygeiss/method2go/services/project"
)

// FileSystem ...
type FileSystem struct {
	files []string
	path  string
}

func (a *FileSystem) GenerateProjectStructure(p *project.Project) (err error) {
	for _, file := range a.files {
		dst := filepath.Join(a.path, file)
		data := []byte(p.Contents[file])
		if err := CreateFoldersByFile(dst); err != nil {
			return err
		}
		if err := os.WriteFile(dst, data, 0644); err != nil {
			return err
		}
	}
	return nil
}

// NewFileSystem ...
func NewFileSystem(path string, files []string) project.Repository {
	return &FileSystem{
		files: files,
		path:  path,
	}
}
