// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package session

import (
	"time"

	"mcpgo/internal/domain/shared"
)

// Session represents a single, stateful interaction between an agent and the gateway.
type Session struct {
	ID        shared.ID
	AgentID   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// NewSession creates a new session for a given agent.
func NewSession(agentID string, duration time.Duration) (*Session, error) {
	if agentID == "" {
		return nil, &shared.ValidationError{Field: "agentID", Msg: "cannot be empty"}
	}

	return &Session{
		ID:        shared.NewID(),
		AgentID:   agentID,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(duration),
	}, nil
}

// IsExpired checks if the session has expired.
func (s *Session) IsExpired() bool {
	return time.Now().UTC().After(s.ExpiresAt)
}
