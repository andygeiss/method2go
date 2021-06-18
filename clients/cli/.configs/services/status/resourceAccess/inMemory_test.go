package resourceaccess_test

import (
	"testing"

	"{{ .Module }}/services/status/resourceaccess"
	assert "github.com/andygeiss/utilities/testing"
)

func TestInMemoryStatus_ReadStatus_Should_Return_Without_An_Error(t *testing.T) {
	statusResourceAccess := resourceaccess.NewInMemoryStatus()
	text, err := statusResourceAccess.ReadStatus()
	assert.That("error should be nil", t, err, nil)
	assert.That("text should not be nil", t, text == "", false)
}
