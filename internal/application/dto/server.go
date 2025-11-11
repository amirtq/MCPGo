// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package dto

import (
	"time"

	"mcpgo/internal/domain/server"
)

// RegisterServerRequest is the DTO for registering a new server.
type RegisterServerRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Protocol string `json:"protocol"`
}

// ServerResponse is the DTO for returning server information.
type ServerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Protocol  string    `json:"protocol"`
	CreatedAt time.Time `json:"created_at"`
}

// ToServerResponse converts a domain Server to a DTO.
func ToServerResponse(s *server.Server) ServerResponse {
	return ServerResponse{
		ID:        s.ID.String(),
		Name:      s.Name,
		Address:   s.Address,
		Protocol:  s.Protocol,
		CreatedAt: s.CreatedAt,
	}
}

// ToServerResponseList converts a list of domain Servers to a list of DTOs.
func ToServerResponseList(servers []*server.Server) []ServerResponse {
	res := make([]ServerResponse, len(servers))
	for i, s := range servers {
		res[i] = ToServerResponse(s)
	}
	return res
}
