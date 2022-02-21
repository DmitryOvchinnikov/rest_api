package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WithCORS struct {
	Router *mux.Router
}

// Simple wrapper to Allow CORS for setting HTTP headers, except OPTIONS.
// See: http://stackoverflow.com/a/24818638/1058612.
func (s *WithCORS) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.Router.ServeHTTP(w, req)
}
