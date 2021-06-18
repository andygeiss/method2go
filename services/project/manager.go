package project

import (
	"context"
	"os"
	"os/exec"
)

// Manager ...
type Manager struct {
	resourceAccess ProjectResourceAccess
	engine         TemplateEngine
	err            error
}

// Create ...
func (a *Manager) CreateByName(c context.Context, name string) (p *Project) {
	// Skip action on a previous error.
	if a.err != nil {
		return
	}
	// Create a new project by name.
	p = NewProject(name)
	// Init templates
	a.engine.InitTemplates(p)
	// Generate the project structure.
	if err := a.resourceAccess.GenerateProjectStructure(p); err != nil {
		a.err = err
	}
	// Initialize Go module
	os.Chdir(name)
	exec.Command("go", "mod", "init").Run()
	// Initialize Git repository
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "initial commit", ".").Run()
	return
}

// Error ...
func (a *Manager) Error() error {
	return a.err
}

// NewManager ...
func NewManager(e TemplateEngine, ra ProjectResourceAccess) *Manager {
	return &Manager{
		resourceAccess: ra,
		engine:         e,
	}
}
