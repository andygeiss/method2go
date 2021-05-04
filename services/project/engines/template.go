package engines

import (
	"io"
	"io/fs"
	"os"

	"github.com/andygeiss/method2go/services/project"
)

// DefaultTemplateEngine ...
type DefaultTemplateEngine struct {
	access fs.FS
	files  []string
	path   string
}

func (a *DefaultTemplateEngine) InitTemplates(p *project.Project) error {
	for _, file := range a.files {
		p.Templates[file] = safeReadAll(a.path + "/" + file)
	}
	return nil
}

// NewDefaultTemplateEngine ...
func NewDefaultTemplateEngine(access fs.FS, path string, files []string) project.TemplateEngine {
	return &DefaultTemplateEngine{
		access: access,
		files:  files,
		path:   path,
	}
}

func safeReadAll(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(content)
}
