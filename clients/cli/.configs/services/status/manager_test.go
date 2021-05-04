package status_test

import (
	"context"
	"testing"

	"github.com/andygeiss/method2go_template/services/status"
	assert "github.com/andygeiss/utilities/testing"
)

func TestManager_GetStatus_Should_Return_Without_An_Error(t *testing.T) {
	manager := status.NewManager("status.Manager")
	text := manager.GetStatus(context.Background())
	err := manager.Error()
	assert.That("error should be nil", t, err, nil)
	assert.That("text should be OK", t, text, "OK")
}
