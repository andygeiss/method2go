package project_test

import "github.com/andygeiss/method2go/resources"

type MockupResourceAccess struct{}

func (a *MockupResourceAccess) GenerateProjectStructure(p *resources.Project) (err error) {
	return nil
}

type MockupTemplateEngine struct{}

func (a *MockupTemplateEngine) InitTemplates(p *resources.Project) error {
	return nil
}
