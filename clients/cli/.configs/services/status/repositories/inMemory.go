package repositories

import (
	"{{ .Module }}/services/status"
)

// InMemoryStatus ...
type InMemoryStatus struct{}

func (a *InMemoryStatus) ReadStatus() (text string, err error) {
	return "OK", nil
}

// NewInMemoryStatus ...
func NewInMemoryStatus() status.ResourceAccess {
	return &InMemoryStatus{}
}
