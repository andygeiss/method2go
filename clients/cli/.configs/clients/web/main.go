package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

//go:embed static
var content embed.FS

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)
	// Use embedding mode
	fsys, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}
	// Handlers
	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":3080", nil)
}
