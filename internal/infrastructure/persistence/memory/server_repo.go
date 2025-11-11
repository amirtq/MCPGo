// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package memory

import (
	"context"
	"fmt"
	"sync"

	"mcpgo/internal/application/ports"
	"mcpgo/internal/domain/server"
	"mcpgo/internal/domain/shared"
)

// InMemoryServerRepo is an in-memory implementation of the ServerRepo.
type InMemoryServerRepo struct {
	mu      sync.RWMutex
	servers map[string]*server.Server
}

// NewInMemoryServerRepo creates a new in-memory server repository.
func NewInMemoryServerRepo() *InMemoryServerRepo {
	return &InMemoryServerRepo{
		servers: make(map[string]*server.Server),
	}
}

// Save persists a server to memory.
func (r *InMemoryServerRepo) Save(ctx context.Context, s *server.Server) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.servers[s.ID.String()] = s
	return nil
}

// FindByID retrieves a server from memory by its ID.
func (r *InMemoryServerRepo) FindByID(ctx context.Context, id shared.ID) (*server.Server, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s, ok := r.servers[id.String()]
	if !ok {
		return nil, fmt.Errorf("server with ID %s not found", id.String())
	}
	return s, nil
}

// FindAll retrieves all servers from memory.
func (r *InMemoryServerRepo) FindAll(ctx context.Context) ([]*server.Server, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	allServers := make([]*server.Server, 0, len(r.servers))
	for _, s := range r.servers {
		allServers = append(allServers, s)
	}
	return allServers, nil
}

// Static type check
var _ ports.ServerRepo = (*InMemoryServerRepo)(nil)
