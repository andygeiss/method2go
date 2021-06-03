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
	te := engines.NewDefaultTemplateEngine(&content, "testdata/templates", []string{"main.go"})
	project := project.NewProject("test")
	err := te.InitTemplates(project)
	mainGoTemplate, mainGoTemplateDefined := project.Templates["main.go"]
	mainGoContent, mainGoContentDefined := project.Contents["main.go"]
	assert.That("error should be nil", t, err, nil)
	assert.That("main.go template should be defined", t, mainGoTemplateDefined, true)
	assert.That("main.go template should not be empty", t, len(mainGoTemplate) > 0, true)
	assert.That("main.go content should be defined", t, mainGoContentDefined, true)
	assert.That("main.go content should not be empty", t, len(mainGoContent) > 0, true)
}
