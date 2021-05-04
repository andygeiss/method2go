package contracts

// StatusRequest ...
type StatusRequest struct{}

// StatusResponse ...
type StatusResponse struct {
	Text  string `json:"text,omitempty"`
	Error string `json:"error,omitempty"`
}
