package status_test

type MockupResourceAccess struct{}

func (a *MockupResourceAccess) ReadStatus() (text string, err error) {
	return "OK", nil
}

type MockupTransformationEngine struct{}

func (a *MockupTransformationEngine) Transform(in string) (out string) {
	return "ok"
}
