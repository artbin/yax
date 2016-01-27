package httpx

import "net/http"

type HostSwitch map[string]http.Handler

func NewHostSwitch() *HostSwitch {
	return &HostSwitch{}
}

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		// Handle host names for wich no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func (hs HostSwitch) Handle(host string, handler http.Handler) {
	hs[host] = handler
}
