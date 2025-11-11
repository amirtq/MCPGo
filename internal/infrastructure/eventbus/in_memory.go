// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package eventbus

import (
	"context"
	"log"

	"mcpgo/internal/application/ports"
	"mcpgo/internal/domain/shared"
)

// InMemoryEventBus is a simple, in-memory event bus for demonstration purposes.
type InMemoryEventBus struct {
	logger *log.Logger
}

// NewInMemoryEventBus creates a new in-memory event bus.
func NewInMemoryEventBus(logger *log.Logger) *InMemoryEventBus {
	return &InMemoryEventBus{logger: logger}
}

// Publish logs the event to the console.
func (b *InMemoryEventBus) Publish(ctx context.Context, event shared.DomainEvent) error {
	b.logger.Printf("Publishing event: %s, Occurred on: %s", event.EventName(), event.OccurredOn())
	// In a real implementation, this would dispatch the event to subscribers.
	return nil
}

// Static type check
var _ ports.EventBus = (*InMemoryEventBus)(nil)
