// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package application

import "context"

// Command is a marker interface for commands.
type Command interface {
	CommandName() string
}

// Query is a marker interface for queries.
type Query interface {
	QueryName() string
}

// CommandHandler defines the interface for a command handler.
// T is the command type, R is the result type.
type CommandHandler[T Command, R any] interface {
	Handle(ctx context.Context, cmd T) (R, error)
}

// QueryHandler defines the interface for a query handler.
// T is the query type, R is the result type.
type QueryHandler[T Query, R any] interface {
	Handle(ctx context.Context, query T) (R, error)
}
