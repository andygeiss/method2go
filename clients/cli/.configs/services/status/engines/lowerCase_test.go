package engines_test

import (
	"testing"

	"{{ .Module }}/services/status/engines"
	assert "github.com/andygeiss/utilities/testing"
)

func TestTransformationEngine_Transform_Should_Return_Without_A_Correct_Text(t *testing.T) {
	engine := engines.NewLowerCaseTransformationEngine()
	text := engine.Transform("OK")
	assert.That("text should be ok", t, text, "ok")
}
