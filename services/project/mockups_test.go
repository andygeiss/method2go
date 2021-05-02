package project_test

import "github.com/andygeiss/method2go/services/project"

type MockupResourceAccess struct{}

func (a *MockupResourceAccess) GenerateProjectStructure(p *project.Project) (err error) {
	return nil
}

type MockupTemplateEngine struct{}

func (a *MockupTemplateEngine) InitTemplates(p *project.Project) error {
	return nil
}
