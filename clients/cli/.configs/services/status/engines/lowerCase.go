package engines

import (
	"strings"
	
	"{{ .Module }}/services/status"
)

// TransformationEngine ...
type TransformationEngine struct{}

func (a *TransformationEngine) Transform(in string) (out string) {
	return strings.ToLower(in)
}

// NewTransformationEngine ...
func NewTransformationEngine() status.TransformationEngine {
	return &TransformationEngine{}
}
