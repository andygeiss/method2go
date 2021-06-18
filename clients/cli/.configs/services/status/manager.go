package status

import (
	"context"

	"github.com/andygeiss/utilities/tracing"
)

// Manager ...
type Manager struct {
	statusResourceAccess StatusResourceAccess
	transformationEngine TransformationEngine
	err                  error
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
	// Add actors ...
	trace := tracing.FromContext(ctx)
	trace.Register(manager.ID())
	trace.Register(a.statusResourceAccess.ID())
	trace.Register(a.transformationEngine.ID())
	ctx = trace.ToContext(ctx)
	// Step 1 ...
	ctx = tracing.Call(ctx, a.ID(), a.statusResourceAccess.ID(), "ReadStatus", func() {
		status, err := a.statusResourceAccess.ReadStatus()
		if err != nil {
			a.err = err
			return
		}
		text = status.Text
	})
	// Step 2 ...
	ctx = tracing.Call(ctx, a.ID(), a.transformationEngine.ID(), "Transform", func() {
		text = a.transformationEngine.Transform(text)
	})
	return
}

func (a *Manager) ID() string {
	return "status.Manager"
}

// NewManager ...
func NewManager(transformationEngine TransformationEngine, statusResourceAccess StatusResourceAccess) *Manager {
	return &Manager{
		statusResourceAccess: statusResourceAccess,
		transformationEngine: transformationEngine,
	}
}
