// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package commands

import (
	"context"
	"fmt"

	"mcpgo/internal/application"
	"mcpgo/internal/application/dto"
	"mcpgo/internal/application/ports"
	"mcpgo/internal/domain/shared"
)

// RouteCallCommand is the command to route a call to a server.
type RouteCallCommand struct {
	ServerID string
	Payload  []byte
}

func (c RouteCallCommand) CommandName() string {
	return "commands.RouteCall"
}

// RouteCallHandler is the handler for the RouteCallCommand.
type RouteCallHandler struct {
	serverRepo      ports.ServerRepo
	connectorClient ports.ConnectorClient
}

// NewRouteCallHandler creates a new RouteCallHandler.
func NewRouteCallHandler(serverRepo ports.ServerRepo, connectorClient ports.ConnectorClient) *RouteCallHandler {
	return &RouteCallHandler{
		serverRepo:      serverRepo,
		connectorClient: connectorClient,
	}
}

// Handle executes the RouteCallCommand.
func (h *RouteCallHandler) Handle(ctx context.Context, cmd RouteCallCommand) (dto.RouteResponse, error) {
	serverID, err := shared.FromString(cmd.ServerID)
	if err != nil {
		return dto.RouteResponse{}, fmt.Errorf("invalid server ID format: %w", err)
	}

	server, err := h.serverRepo.FindByID(ctx, serverID)
	if err != nil {
		return dto.RouteResponse{}, fmt.Errorf("server not found: %w", err)
	}

	// In a real implementation, we would use the connectorClient:
	// _, err = h.connectorClient.RouteCall(ctx, server.Address, cmd.Payload)
	// if err != nil {
	// 	return dto.RouteResponse{}, fmt.Errorf("failed to route call: %w", err)
	// }

	// For now, return a mock success response as per requirements.
	return dto.RouteResponse{
		Status:  "success",
		Message: fmt.Sprintf("Mock call routed to server '%s' at %s", server.Name, server.Address),
	}, nil
}

// Static type check
var _ application.CommandHandler[RouteCallCommand, dto.RouteResponse] = (*RouteCallHandler)(nil)
