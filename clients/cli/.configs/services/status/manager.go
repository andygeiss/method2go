package status

import "context"

// Manager ...
type Manager struct {
	ID     string
	access ResourceAccess
	engine TransformationEngine
	err    error
}

// Error ...
func (a *Manager) Error() error {
	return nil
}

// GetStatus ...
func (a *Manager) GetStatus(ctx context.Context) (text string) {
	if a.err != nil {
		return
	}
	text, err := a.access.ReadStatus()
	if err != nil {
		a.err = err
		return
	}
	transformed := a.engine.Transform(text)
	return transformed
}

// NewManager ...
func NewManager(id string, engine TransformationEngine, access ResourceAccess) *Manager {
	return &Manager{
		ID:     id,
		access: access,
		engine: engine,
	}
}
