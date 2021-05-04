package status

import "context"

// Manager ...
type Manager struct {
	ID  string
	err error
}

// Error ...
func (a *Manager) Error() error {
	return nil
}

// GetStatus ...
func (a *Manager) GetStatus(ctx context.Context) (text string) {
	if a.err != nil {
		return a.err.Error()
	}
	return "OK"
}

// NewManager ...
func NewManager(id string) *Manager {
	return &Manager{
		ID: id,
	}
}
