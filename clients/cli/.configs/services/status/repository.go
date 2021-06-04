package status

// Repository ...
type Repository interface {
	ReadStatus() (text string, err error)
}
