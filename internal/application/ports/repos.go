// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package ports

import (
	"context"

	"mcpgo/internal/domain/server"
	"mcpgo/internal/domain/session"
	"mcpgo/internal/domain/shared"
)

// ServerRepo defines the persistence operations for the Server entity.
type ServerRepo interface {
	Save(ctx context.Context, s *server.Server) error
	FindByID(ctx context.Context, id shared.ID) (*server.Server, error)
	FindAll(ctx context.Context) ([]*server.Server, error)
}

// SessionRepo defines the persistence operations for the Session entity.
type SessionRepo interface {
	Save(ctx context.Context, s *session.Session) error
	FindByID(ctx context.Context, id shared.ID) (*session.Session, error)
}

// UnitOfWork defines a transactional boundary for operations.
// This is a placeholder for now and would be more complex in a real DB implementation.
type UnitOfWork interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
