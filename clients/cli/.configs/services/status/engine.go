package status

// TransformationEngine ...
type TransformationEngine interface {
	ID() string
	Transform(in string) (out string)
}
