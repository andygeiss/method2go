package main

import (
	"embed"
	"log"

	"github.com/andygeiss/method2go/services/project/engines"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

//go:embed configs
var content embed.FS

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)

	te := engines.NewDefaultTemplateEngine(&content, "configs/method2go_template", []string{"main.go"})

}
