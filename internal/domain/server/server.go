// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package server

import (
	"time"

	"mcpgo/internal/domain/shared"
)

// Server represents an MCP server that can be connected to.
type Server struct {
	ID        shared.ID
	Name      string
	Address   string
	Protocol  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewServer creates a new MCP server instance.
func NewServer(name, address, protocol string) (*Server, error) {
	// Basic validation can go here.
	// For more complex validation, a dedicated factory or builder might be used.
	if name == "" || address == "" || protocol == "" {
		// In a real app, we'd use custom error types.
		return nil, &shared.ValidationError{Field: "name/address/protocol", Msg: "cannot be empty"}
	}

	return &Server{
		ID:        shared.NewID(),
		Name:      name,
		Address:   address,
		Protocol:  protocol,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}

// UpdateDetails updates the server's mutable properties.
func (s *Server) UpdateDetails(name, address string) error {
	if name == "" || address == "" {
		return &shared.ValidationError{Field: "name/address", Msg: "cannot be empty"}
	}
	s.Name = name
	s.Address = address
	s.UpdatedAt = time.Now().UTC()
	return nil
}
