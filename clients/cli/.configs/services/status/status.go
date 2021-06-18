package status

// Status ...
type Status struct {
	Text string
}

// NewStatus ...
func NewStatus(text string) *Status {
	return &Status{
		Text: text,
	}
}

// StatusResourceAccess ...
type StatusResourceAccess interface {
	ID() string
	ReadStatus() (status *Status, err error)
}
