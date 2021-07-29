package resources

// InMemoryStatusResourceAccess ...
type InMemoryStatusResourceAccess struct{}

func (a *InMemoryStatusResourceAccess) ID() string {
	return "resources.InMemoryStatusResourceAccess"
}

func (a *InMemoryStatusResourceAccess) ReadStatus() (stat *Status, err error) {
	return NewStatus("OK"), nil
}

// NewInMemoryStatusResourceAccess ...
func NewInMemoryStatusResourceAccess() StatusResourceAccess {
	return &InMemoryStatusResourceAccess{}
}
