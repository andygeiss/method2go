package resources

import "github.com/andygeiss/method2go/services/project"

// FileSystem ...
type FileSystem struct {
	path string
}

func (a *FileSystem) GenerateProjectStructure(p *project.Project) (err error) {
	return nil
}

// NewFileSystem ...
func NewFileSystem(path string) project.ResourceAccess {
	return &FileSystem{
		path: path,
	}
}
