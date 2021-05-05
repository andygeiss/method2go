package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/andygeiss/method2go/services/project"
	"github.com/andygeiss/method2go/services/project/engines"
	"github.com/andygeiss/method2go/services/project/resources"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

//go:embed .configs
var content embed.FS

func ensureDefaults() {
	if len(os.Args) == 1 {
		fmt.Printf("\n%s %s (%s)\n", name, version, build)
		fmt.Printf(`
This tool generates a basic Go system design based on building-blocks from "The Method".

Usage:	  ` + os.Args[0] + ` <Project-Name>

Example:  ` + os.Args[0] + ` template

`)
		os.Exit(255)
	}
}

func main() {
	files := []string{
		filepath.FromSlash("clients/api/contracts/contracts.go"),
		filepath.FromSlash("clients/api/contracts/contracts.http"),
		filepath.FromSlash("clients/api/handlers/status.go"),
		filepath.FromSlash("clients/api/handlers/utils.go"),
		filepath.FromSlash("clients/api/main.go"),
		filepath.FromSlash("clients/cli/main.go"),
		filepath.FromSlash("clients/web/scripts/app.js"),
		filepath.FromSlash("clients/web/styles/app.scss"),
		filepath.FromSlash("clients/web/styles/colors.scss"),
		filepath.FromSlash("clients/web/index.html"),
		filepath.FromSlash("services/status/engines/lowerCase.go"),
		filepath.FromSlash("services/status/engines/lowerCase_test.go"),
		filepath.FromSlash("services/status/resources/inMemory.go"),
		filepath.FromSlash("services/status/resources/inMemory_test.go"),
		filepath.FromSlash("services/status/engine.go"),
		filepath.FromSlash("services/status/manager.go"),
		filepath.FromSlash("services/status/manager_test.go"),
		filepath.FromSlash("services/status/mockups_test.go"),
		filepath.FromSlash("services/status/resourceAccess.go"),
		"Makefile",
		"status.plantuml",
	}
	ensureDefaults()
	name := os.Args[1]
	ra := resources.NewFileSystem(name, files)
	te := engines.NewDefaultTemplateEngine(&content, ".configs", files)
	pm := project.NewManager(te, ra)
	_ = pm.CreateByName(context.Background(), name)
}
