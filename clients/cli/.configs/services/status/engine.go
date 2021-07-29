package status

import (
	"strings"
)

// TransformationEngine ...
type TransformationEngine interface {
	ID() string
	Transform(in string) (out string)
}

// LowerCaseTransformationEngine ...
type LowerCaseTransformationEngine struct{}

func (a *LowerCaseTransformationEngine) ID() string {
	return "status.LowerCaseTransformationEngine"
}

func (a *LowerCaseTransformationEngine) Transform(in string) (out string) {
	return strings.ToLower(in)
}

// NewLowerCaseTransformationEngine ...
func NewLowerCaseTransformationEngine() TransformationEngine {
	return &LowerCaseTransformationEngine{}
}
