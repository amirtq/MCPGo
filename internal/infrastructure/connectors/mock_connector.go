// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package connectors

import (
	"context"
	"log"

	"mcpgo/internal/application/ports"
)

// MockConnectorClient is a mock implementation of the ConnectorClient.
type MockConnectorClient struct {
	logger *log.Logger
}

// NewMockConnectorClient creates a new mock connector client.
func NewMockConnectorClient(logger *log.Logger) *MockConnectorClient {
	return &MockConnectorClient{logger: logger}
}

// RouteCall simulates routing a call to an external server.
func (c *MockConnectorClient) RouteCall(ctx context.Context, serverAddress string, payload []byte) ([]byte, error) {
	c.logger.Printf("Mock routing call to %s with payload: %s", serverAddress, string(payload))
	// In a real implementation, this would make an HTTP/gRPC call.
	return []byte("mock response"), nil
}

// Static type check
var _ ports.ConnectorClient = (*MockConnectorClient)(nil)
