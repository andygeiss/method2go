package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/andygeiss/method2go_template/clients/api/handlers"
	"github.com/andygeiss/method2go_template/services/status"
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
