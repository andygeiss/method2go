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
	path   string
}

func (a *DefaultTemplateEngine) InitTemplates(p *project.Project) error {
	p.Templates["main.go"] = safeReadAll(a.path + "/" + "main.go")
	return nil
}

// NewDefaultTemplateEngine ...
func NewDefaultTemplateEngine(access fs.FS, path string) project.TemplateEngine {
	return &DefaultTemplateEngine{
		access: access,
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
