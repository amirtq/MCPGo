// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package queries

import (
	"context"

	"mcpgo/internal/application"
	"mcpgo/internal/application/dto"
	"mcpgo/internal/application/ports"
)

// ListServersQuery is the query to get all registered servers.
type ListServersQuery struct{}

func (q ListServersQuery) QueryName() string {
	return "queries.ListServers"
}

// ListServersHandler is the handler for the ListServersQuery.
type ListServersHandler struct {
	serverRepo ports.ServerRepo
}

// NewListServersHandler creates a new ListServersHandler.
func NewListServersHandler(serverRepo ports.ServerRepo) *ListServersHandler {
	return &ListServersHandler{serverRepo: serverRepo}
}

// Handle executes the ListServersQuery.
func (h *ListServersHandler) Handle(ctx context.Context, query ListServersQuery) ([]dto.ServerResponse, error) {
	servers, err := h.serverRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToServerResponseList(servers), nil
}

// Static type check
var _ application.QueryHandler[ListServersQuery, []dto.ServerResponse] = (*ListServersHandler)(nil)
