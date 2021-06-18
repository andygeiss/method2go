package status

// Repository ...
type Repository interface {
	ID() string
	ReadStatus() (text string, err error)
}
