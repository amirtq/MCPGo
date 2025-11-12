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
)

// healthHandler is the HTTP handler for the health check.
// It calls the application layer to get the health status and returns it.
func (r *Router) healthHandler(w http.ResponseWriter, req *http.Request) {
	status := r.app.CheckHealth()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
