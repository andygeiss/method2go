package project

import (
	"encoding/hex"

	"github.com/andygeiss/utilities/security"
)

// Project ...
type Project struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Contents  map[string]string `json:"contents"`
	Templates map[string]string `json:"templates"`
}

// NewProject ...
func NewProject(name string) *Project {
	id := security.NewKey256()
	return &Project{
		ID:        hex.EncodeToString(id[:]),
		Name:      name,
		Contents:  make(map[string]string),
		Templates: make(map[string]string),
	}
}
