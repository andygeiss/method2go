package main

import (
	"context"
	"embed"
	"log"

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

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)
	files := []string{
		"clients/api/main.go",
	}
	ra := resources.NewFileSystem("testdata", files)
	te := engines.NewDefaultTemplateEngine(&content, ".configs", files)
	pm := project.NewManager(te, ra)

	_ = pm.CreateByName(context.Background(), "testdata")
}
