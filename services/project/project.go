package project

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"

	"github.com/andygeiss/utilities/security"
)

// Project ...
type Project struct {
	ID        string            `json:"id"`
	Contents  map[string]string `json:"contents"`
	Module    string            `json:"module"`
	Name      string            `json:"name"`
	Templates map[string]string `json:"templates"`
}

// NewProject ...
func NewProject(name string) *Project {
	wd, _ := os.Getwd()
	id := security.NewKey256()
	return &Project{
		ID:        hex.EncodeToString(id[:]),
		Module:    getModulePathFromAbsolute(filepath.Join(wd, name)),
		Name:      name,
		Contents:  make(map[string]string),
		Templates: make(map[string]string),
	}
}

func getModulePathFromAbsolute(path string) string {
	parts := strings.Split(path, "src"+string(filepath.Separator))
	if len(parts) > 1 {
		return parts[1]
	}
	return parts[0]
}
