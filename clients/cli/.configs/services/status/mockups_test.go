package status_test

type MockupRepository struct{}

func (a *MockupRepository) ID() string {
	return "MockupRepository"
}

func (a *MockupRepository) ReadStatus() (text string, err error) {
	return "OK", nil
}

type MockupTransformationEngine struct{}

func (a *MockupTransformationEngine) ID() string {
	return "MockupTransformationEngine"
}

func (a *MockupTransformationEngine) Transform(in string) (out string) {
	return "ok"
}
