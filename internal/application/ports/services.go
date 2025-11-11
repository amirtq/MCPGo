// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package ports

import (
	"context"

	"mcpgo/internal/domain/shared"
)

// ConnectorClient defines the interface for communicating with an external MCP server.
type ConnectorClient interface {
	// RouteCall would be a more complex method in a real scenario,
	// likely taking a structured request and returning a structured response.
	RouteCall(ctx context.Context, serverAddress string, payload []byte) ([]byte, error)
}

// EventBus defines the interface for a domain event publisher.
type EventBus interface {
	Publish(ctx context.Context, event shared.DomainEvent) error
}
