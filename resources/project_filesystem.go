package resources

import (
	"os"
	"path/filepath"
)

// ProjectResourceAccessFileSystem ...
type ProjectResourceAccessFileSystem struct {
	files []string
	path  string
}

func (a *ProjectResourceAccessFileSystem) GenerateProjectStructure(p *Project) (err error) {
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

// NewProjectResourceAccessFileSystem ...
func NewProjectResourceAccessFileSystem(path string, files []string) ProjectResourceAccess {
	return &ProjectResourceAccessFileSystem{
		files: files,
		path:  path,
	}
}
