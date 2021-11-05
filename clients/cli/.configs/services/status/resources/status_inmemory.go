package resources

import (
	"{{ .Module }}/services/status"
)


// InMemoryStatusResourceAccess ...
type InMemoryStatusResourceAccess struct {
	status *status.Status
}

func (a *InMemoryStatusResourceAccess) ID() string {
	return "status.InMemoryStatusResourceAccess"
}

func (a *InMemoryStatusResourceAccess) ReadStatus() (status *status.Status, err error) {
	return a.status, nil
}

// NewInMemoryStatusResourceAccess ...
func NewInMemoryStatusResourceAccess() status.StatusResourceAccess {
	return &InMemoryStatusResourceAccess{
		status: &status.Status{Text: "OK"},
	}
}
