package main

import (
	"context"
	"embed"
	"fmt"
	"os"

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
  This tool generates a basic Go project structure based on "The Method".

Usage:	  ` + os.Args[0] + ` <Project-Name>

Example:  ` + os.Args[0] + ` template
`)
		os.Exit(255)
	}
}

func main() {
	files := []string{
		"clients/api/contracts/contracts.go",
		"clients/api/contracts/contracts.http",
		"clients/api/handlers/status.go",
		"clients/api/handlers/utils.go",
		"clients/api/main.go",
		"clients/cli/main.go",
		"clients/web/scripts/app.js",
		"clients/web/styles/app.scss",
		"clients/web/styles/colors.scss",
		"clients/web/index.html",
		"services/status/manager.go",
		"services/status/manager_test.go",
		"Makefile",
	}
	ensureDefaults()
	name := os.Args[1]
	ra := resources.NewFileSystem(name, files)
	te := engines.NewDefaultTemplateEngine(&content, ".configs", files)
	pm := project.NewManager(te, ra)
	_ = pm.CreateByName(context.Background(), name)
}
