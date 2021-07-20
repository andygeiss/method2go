package resources

import (
	"{{ .Module }}/services/status"
)

// InMemoryStatusResourceAccess ...
type InMemoryStatusResourceAccess struct {}

func (a *InMemoryStatusResourceAccess) ID() string {
	return "resources.InMemoryStatusResourceAccess"
}

func (a *InMemoryStatusResourceAccess) ReadStatus() (stat *status.Status, err error) {
	return status.NewStatus("OK"), nil 
}

// NewInMemoryStatusResourceAccess ...
func NewInMemoryStatusResourceAccess() status.StatusResourceAccess {
	return &InMemoryStatusResourceAccess{}
}
