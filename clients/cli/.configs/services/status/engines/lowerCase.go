package engines

import (
	"{{ .Module }}/services/status"
)

// TransformationEngine ...
type TransformationEngine struct{}

func (a *TransformationEngine) Transform(in string) (out string) {
	return ""
}

// NewTransformationEngine ...
func NewTransformationEngine() status.TransformationEngine {
	return &TransformationEngine{}
}
