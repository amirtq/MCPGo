// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package http

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter creates a new HTTP router and registers the handlers.
func NewRouter(handlers *Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /health", handlers.HealthHandler)

	mux.HandleFunc("POST /servers", handlers.RegisterServerHandler)
	mux.HandleFunc("GET /servers", handlers.ListServersHandler)

	mux.HandleFunc("POST /route", handlers.RouteCallHandler)

	return mux
}
