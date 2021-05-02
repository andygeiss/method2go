package engines_test

import (
	"embed"
	"testing"

	"github.com/andygeiss/method2go/services/project"
	"github.com/andygeiss/method2go/services/project/engines"
	assert "github.com/andygeiss/utilities/testing"
)

//go:embed testdata/templates
var content embed.FS

func TestDefaultTemplateEngine_InitTemplates_Should_Return_Without_An_Error(t *testing.T) {
	te := engines.NewDefaultTemplateEngine(&content, "testdata/templates")
	project := project.NewProject("test")
	err := te.InitTemplates(project)
	mainGoContent, mainGoDefined := project.Templates["main.go"]
	assert.That("error should be nil", t, err, nil)
	assert.That("main.go should be defined", t, mainGoDefined, true)
	assert.That("main.go content should not be empty", t, len(mainGoContent) > 0, true)
}
