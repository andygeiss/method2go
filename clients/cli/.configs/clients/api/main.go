package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"{{ .Module }}/clients/api/handlers"
	"{{ .Module }}/services/status/engines"
	"{{ .Module }}/services/status/resources"
	"{{ .Module }}/services/status"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)
	// ResourceAccess
	access := resources.NewInMemoryStatus()
	// Engines
	engine := engines.NewTransformationEngine()
	// Managers
	statusManager := status.NewManager("status.Manager", engine, access)
	// Handlers
	http.HandleFunc("/status", handlers.NewStatusHandler(statusManager))
	http.ListenAndServe(":3000", nil)
}
