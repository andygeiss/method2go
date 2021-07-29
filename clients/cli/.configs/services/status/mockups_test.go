package status_test

import "{{ .Module }}/resources"

type MockupRepository struct{}

func (a *MockupRepository) ID() string {
	return "MockupRepository"
}

func (a *MockupRepository) ReadStatus() (stat *resources.Status, err error) {
	return resources.NewStatus("OK"), nil
}

type MockupTransformationEngine struct{}

func (a *MockupTransformationEngine) ID() string {
	return "MockupTransformationEngine"
}

func (a *MockupTransformationEngine) Transform(in string) (out string) {
	return "ok"
}
