package engines

import (
	"strings"
	
	"{{ .Module }}/services/status"
)

// LowerCaseTransformationEngine ...
type LowerCaseTransformationEngine struct{}

func (a *LowerCaseTransformationEngine) ID() string {
	return "engines.LowerCaseTransformationEngine"
}

func (a *LowerCaseTransformationEngine) Transform(in string) (out string) {
	return strings.ToLower(in)
}

// NewLowerCaseTransformationEngine ...
func NewLowerCaseTransformationEngine() status.TransformationEngine {
	return &LowerCaseTransformationEngine{}
}
