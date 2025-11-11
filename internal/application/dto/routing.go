// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package dto

// RouteRequest is the DTO for a routing request.
type RouteRequest struct {
	ServerID string `json:"server_id"`
	Payload  []byte `json:"payload"`
}

// RouteResponse is the DTO for a routing response.
type RouteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// In a real implementation, this would be a structured response.
	Data []byte `json:"data,omitempty"`
}
