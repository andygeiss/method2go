package main

import (
	"context"
	"log"
	"time"
	
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
	statusResourceAccess := resources.NewInMemoryStatusResourceAccess()
	// Engines
	transformationEngine := engines.NewLowerCaseTransformationEngine()
	// Managers
	statusManager := status.NewManager(transformationEngine, statusResourceAccess)
	// Set a timeout for the whole use case
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	// Run the action
	text := statusManager.GetStatus(ctx)
	// Handle errors and response
	if err := statusManager.Error(); err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("result text: %s", text)
}
