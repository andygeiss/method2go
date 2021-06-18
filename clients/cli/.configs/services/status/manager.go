package status

import (
	"context"

	"github.com/andygeiss/utilities/tracing"
)

// Manager ...
type Manager struct {
	repository Repository
	engine     TransformationEngine
	err        error
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
	// Step 1 ...
	ctx = tracing.Call(ctx, a.ID(), a.repository.ID(), "ReadStatus", func() {
		status, err := a.repository.ReadStatus()
		if err != nil {
			a.err = err
			return
		}
		text = status
	})
	// Step 2 ...
	ctx = tracing.Call(ctx, a.ID(), a.engine.ID(), "Transform", func() {
		text = a.engine.Transform(text)
	})
	return
}

func (a *Manager) ID() string {
	return "status.Manager"
}

// NewManager ...
func NewManager(engine TransformationEngine, repository Repository) *Manager {
	return &Manager{
		repository: repository,
		engine:     engine,
	}
}
