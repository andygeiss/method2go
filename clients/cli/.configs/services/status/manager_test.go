package status_test

import (
	"context"
	"testing"

	"{{ .Module }}/services/status"
	assert "github.com/andygeiss/utilities/testing"
)

func TestManager_GetStatus_Should_Return_Without_An_Error(t *testing.T) {
	access := &MockupResourceAccess{}
	engine := &MockupTransformationEngine{}
	manager := status.NewManager("status.Manager", engine, access)
	text := manager.GetStatus(context.Background())
	err := manager.Error()
	assert.That("error should be nil", t, err, nil)
	assert.That("text should be ok", t, text, "ok")
}
