package resourceaccess

import (
	"{{ .Module }}/services/status"
)

// InMemoryStatus ...
type InMemoryStatus struct{}

func (a *InMemoryStatus) ID() string {
	return "resourceAccess.InMemoryStatus"
}

func (a *InMemoryStatus) ReadStatus() (text string, err error) {
	return "OK", nil
}

// NewInMemoryStatus ...
func NewInMemoryStatus() status.StatusResourceAccess {
	return &InMemoryStatus{}
}
