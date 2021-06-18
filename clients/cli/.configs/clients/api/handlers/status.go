package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"{{ .Module }}/clients/api/contracts"
	"{{ .Module }}/services/status"
	"github.com/andygeiss/utilities/logging"
	"github.com/andygeiss/utilities/tracing"
)

// NewStatusHandler ...
func NewStatusHandler(manager *status.Manager) (hf http.HandlerFunc) {
	handlerID := "handlers.NewStatusHandler"
	// Init the state here ...
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		switch r.Method {
		case "POST":
			var req *contracts.StatusRequest
			var res *contracts.StatusResponse
			var err error
			start := time.Now()
			// Add tracing to the context
			ctx, rid := tracing.NewRequestContextWithID()
			// Set a timeout for the whole use case
			ctx, cancel := context.WithTimeout(ctx, time.Second*3)
			defer cancel()
			tracePath := filepath.Join("tracing")
			// Add actors ...
			trace := tracing.FromContext(ctx)
			trace.Register(handlerID)
			trace.Register(manager.ID())
			ctx = trace.ToContext(ctx)
			// Do the calls ...
			ctx = tracing.Call(ctx, handlerID, manager.ID(), "GetStatus", func() {
				// Decode the request
				err = json.NewDecoder(r.Body).Decode(&req)
				// Call the manager ...
				text := manager.GetStatus(ctx)
				// Handle errors and response
				if manager.Error() != nil {
					err = manager.Error()
				}
				if err != nil {
					res = &contracts.StatusResponse{Error: err.Error()}
					err = json.NewEncoder(w).Encode(res)
					logging.Response(r, rid, req, err, start)
					return
				}
				// OK Response
				res = &contracts.StatusResponse{Text: text}
				err = json.NewEncoder(w).Encode(res)
				logging.Response(r, rid, req, err, start)
			})
			defer func() {
				tracing.FromContext(ctx).ToFile(tracePath)
			}()
		}
	}
}
