package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"{{ .Module }}/clients/api/handlers"
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
	// ResourceAccess
	repository := repositories.NewInMemoryStatus()
	// Engines
	engine := engines.NewLowerCaseTransformationEngine()
	// Managers
	statusManager := status.NewManager(engine, repository)
	// Handlers
	http.HandleFunc("/status", handlers.NewStatusHandler(statusManager))
	http.ListenAndServe(":3000", nil)
}
