package main

import (
	"context"
	"log"
	"time"
	
	"{{ .Module }}/services/status/engines"
	"{{ .Module }}/services/status/repositories"
	"{{ .Module }}/services/status"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)
	// Handlers
	http.HandleFunc("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3080", nil)
}
