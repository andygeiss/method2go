package status_test

import (
	"context"
	"testing"

	"{{ .Module }}/services/status"
	assert "github.com/andygeiss/utilities/testing"
)

func TestManager_GetStatus_Should_Return_Without_An_Error(t *testing.T) {
	repository := &MockupRepository{}
	engine := &MockupTransformationEngine{}
	manager := status.NewManager(engine, repository)
	text := manager.GetStatus(context.Background())
	err := manager.Error()
	assert.That("error should be nil", t, err, nil)
	assert.That("text should be ok", t, text, "ok")
}
