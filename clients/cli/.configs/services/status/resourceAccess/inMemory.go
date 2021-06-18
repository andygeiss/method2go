package resourceaccess

import (
	"{{ .Module }}/services/status"
)


// InMemoryStatus ...
type InMemoryStatus struct {
	status *status.Status
}

func (a *InMemoryStatus) ID() string {
	return "resourceAccess.InMemoryStatus"
}

func (a *InMemoryStatus) ReadStatus() (status *status.Status, err error) {
	return a.status, nil
}

// NewInMemoryStatus ...
func NewInMemoryStatus() status.StatusResourceAccess {
	return &InMemoryStatus{
		status: &status.Status{Text: "OK"},
	}
}
