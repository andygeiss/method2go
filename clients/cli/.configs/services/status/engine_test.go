package status_test

import (
	"testing"

	"{{ .Module }}/services/status"
	assert "github.com/andygeiss/utilities/testing"
)

func TestTransformationEngine_Transform_Should_Return_Without_A_Correct_Text(t *testing.T) {
	engine := status.NewLowerCaseTransformationEngine()
	text := engine.Transform("OK")
	assert.That("text should be ok", t, text, "ok")
}
