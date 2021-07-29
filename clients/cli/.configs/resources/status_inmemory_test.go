package resources_test

import (
	"testing"

	"{{ .Module }}/resources"
	assert "github.com/andygeiss/utilities/testing"
)

func TestInMemoryStatusResourceAccess_ReadStatus_Should_Return_Without_An_Error(t *testing.T) {
	access := resources.NewInMemoryStatusResourceAccess()
	_, err := access.ReadStatus()
	assert.That("error should be nil", t, err, nil)
}

func TestInMemoryStatusResourceAccess_ReadStatus_Should_Return_Status_Text_OK(t *testing.T) {
	access := resources.NewInMemoryStatusResourceAccess()
	stat, _ := access.ReadStatus()
	assert.That("status text should be ok", t, stat.Text, "OK")
}
