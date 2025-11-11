// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package commands

import (
	"context"

	"mcpgo/internal/application"
	"mcpgo/internal/application/ports"
	"mcpgo/internal/domain/server"
	"mcpgo/internal/domain/shared"
)

// RegisterServerCommand is the command to register a new server.
type RegisterServerCommand struct {
	Name     string
	Address  string
	Protocol string
}

func (c RegisterServerCommand) CommandName() string {
	return "commands.RegisterServer"
}

// RegisterServerHandler is the handler for the RegisterServerCommand.
type RegisterServerHandler struct {
	serverRepo ports.ServerRepo
	eventBus   ports.EventBus
}

// NewRegisterServerHandler creates a new RegisterServerHandler.
func NewRegisterServerHandler(serverRepo ports.ServerRepo, eventBus ports.EventBus) *RegisterServerHandler {
	return &RegisterServerHandler{
		serverRepo: serverRepo,
		eventBus:   eventBus,
	}
}

// Handle executes the RegisterServerCommand.
// The result is the ID of the newly created server.
func (h *RegisterServerHandler) Handle(ctx context.Context, cmd RegisterServerCommand) (string, error) {
	newServer, err := server.NewServer(cmd.Name, cmd.Address, cmd.Protocol)
	if err != nil {
		return "", err
	}

	if err := h.serverRepo.Save(ctx, newServer); err != nil {
		return "", err
	}

	// Publish a domain event (optional, but good practice)
	event := shared.ServerRegistered{
		BaseEvent: shared.NewBaseEvent(),
		ServerID:  newServer.ID,
		Name:      newServer.Name,
		Address:   newServer.Address,
	}
	_ = h.eventBus.Publish(ctx, event) // Fire-and-forget for now

	return newServer.ID.String(), nil
}

// Static type check to ensure handler implements the generic interface.
var _ application.CommandHandler[RegisterServerCommand, string] = (*RegisterServerHandler)(nil)
