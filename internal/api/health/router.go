// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package health

import (
	"net/http"

	"mcpgo/internal/apps/health"
)

// Router for the health API.
type Router struct {
	app *health.App
}

// NewRouter creates a new health API router.
func NewRouter(app *health.App) *Router {
	return &Router{app: app}
}

// RegisterRoutes registers the health check routes.
func (r *Router) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", r.healthHandler)
}
