package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"{{ .Module }}/clients/api/handlers"
	"{{ .Module }}/services/status"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)
	// Managers
	statusManager := status.NewManager("status.Manager")
	// Handlers
	http.HandleFunc("/status", handlers.NewStatusHandler(statusManager))
	http.ListenAndServe(":3000", nil)
}
