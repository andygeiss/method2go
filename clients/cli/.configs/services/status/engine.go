package status

// TransformationEngine ...
type TransformationEngine interface {
	Transform(in string) (out string)
}
