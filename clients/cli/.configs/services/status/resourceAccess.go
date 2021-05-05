package status

// ResourceAccess ...
type ResourceAccess interface {
	ReadStatus() (text string, err error)
}
