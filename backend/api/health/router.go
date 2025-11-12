// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package health

import (
	"encoding/json"
	"net/http"

	"mcpgo/backend/apps/health"

	"github.com/gorilla/mux"
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
func (r *Router) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/health", r.healthHandler).Methods("GET")
}

// @Summary Health check
// @Description Health check
// @Tags health
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (r *Router) healthHandler(w http.ResponseWriter, req *http.Request) {
	status := r.app.CheckHealth()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
