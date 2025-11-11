// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package shared

import "time"

// DomainEvent represents an event that occurred in the domain.
type DomainEvent interface {
	OccurredOn() time.Time
	EventName() string
}

// BaseEvent provides a base implementation for DomainEvent.
type BaseEvent struct {
	occurredOn time.Time
}

// NewBaseEvent creates a new base event.
func NewBaseEvent() BaseEvent {
	return BaseEvent{occurredOn: time.Now().UTC()}
}

// OccurredOn returns the time the event occurred.
func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}

// --- Example Domain Event ---

// ServerRegistered is an example of a specific domain event.
type ServerRegistered struct {
	BaseEvent
	ServerID  ID
	Name      string
	Address   string
}

// EventName returns the name of the event.
func (e ServerRegistered) EventName() string {
	return "domain.server.registered"
}
