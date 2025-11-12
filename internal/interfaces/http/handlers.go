// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package http

import (
	"encoding/json"
	"net/http"

	"mcpgo/internal/application"
	"mcpgo/internal/application/commands"
	"mcpgo/internal/application/dto"
	"mcpgo/internal/application/queries"
)

// Handlers holds the CQRS handlers.
type Handlers struct {
	RegisterServer application.CommandHandler[commands.RegisterServerCommand, string]
	ListServers    application.QueryHandler[queries.ListServersQuery, []dto.ServerResponse]
	RouteCall      application.CommandHandler[commands.RouteCallCommand, dto.RouteResponse]
}

// RegisterServerHandler handles the HTTP request to register a server.
// @Summary Register a new server
// @Description Register a new server
// @Tags servers
// @Accept  json
// @Produce  json
// @Param server body dto.RegisterServerRequest true "Server to register"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /servers [post]
func (h *Handlers) RegisterServerHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterServerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cmd := commands.RegisterServerCommand{
		Name:     req.Name,
		Address:  req.Address,
		Protocol: req.Protocol,
	}

	serverID, err := h.RegisterServer.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": serverID})
}

// ListServersHandler handles the HTTP request to list all servers.
// @Summary List all servers
// @Description List all servers
// @Tags servers
// @Produce  json
// @Success 200 {array} dto.ServerResponse
// @Failure 500 {string} string "Internal server error"
// @Router /servers [get]
func (h *Handlers) ListServersHandler(w http.ResponseWriter, r *http.Request) {
	query := queries.ListServersQuery{}
	servers, err := h.ListServers.Handle(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servers)
}

// RouteCallHandler handles the HTTP request to route a call.
// @Summary Route a call to a server
// @Description Route a call to a server
// @Tags routing
// @Accept  json
// @Produce  json
// @Param route body dto.RouteRequest true "Route request"
// @Success 200 {object} dto.RouteResponse
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /route [post]
func (h *Handlers) RouteCallHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cmd := commands.RouteCallCommand{
		ServerID: req.ServerID,
		Payload:  req.Payload,
	}

	resp, err := h.RouteCall.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// HealthzHandler is a simple health check endpoint.
// @Summary Health check
// @Description Health check
// @Tags health
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handlers) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
