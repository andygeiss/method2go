package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"{{ .Module }}/clients/api/handlers"
	"{{ .Module }}/services/status/engines"
	"{{ .Module }}/services/status/resourceaccess"
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
	statusResourceAccess := resourceaccess.NewInMemoryStatusResourceAccess()
	// Engines
	transformationEngine := engines.NewLowerCaseTransformationEngine()
	// Managers
	statusManager := status.NewManager(transformationEngine, statusResourceAccess)
	// Handlers
	http.HandleFunc("/status", handlers.NewStatusHandler(statusManager))
	http.ListenAndServe(":3000", nil)
}
