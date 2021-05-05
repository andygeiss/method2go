package project_test

import (
	"context"
	"os"
	"testing"

	"github.com/andygeiss/method2go/services/project"
	assert "github.com/andygeiss/utilities/testing"
)

func Test_Manager_CreateProjectByName_Should_Return_Without_An_Error(t *testing.T) {
	ctx := context.Background()
	ra := &MockupResourceAccess{}
	te := &MockupTemplateEngine{}
	projectManager := project.NewManager(te, ra)
	project := projectManager.CreateByName(ctx, "foo")
	assert.That("error should be nil", t, projectManager.Error(), nil)
	assert.That("project should not be nil", t, project == nil, false)
	os.RemoveAll("go.mod")
}
