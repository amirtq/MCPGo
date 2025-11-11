// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package shared

import (
	"fmt"

	"github.com/google/uuid"
)

// ID represents a unique entity identifier.
type ID struct {
	value string
}

// NewID creates a new, random ID.
func NewID() ID {
	return ID{value: uuid.New().String()}
}

// FromString creates an ID from a string representation.
func FromString(value string) (ID, error) {
	parsed, err := uuid.Parse(value)
	if err != nil {
		return ID{}, fmt.Errorf("invalid ID format: %w", err)
	}
	return ID{value: parsed.String()}, nil
}

// String returns the string representation of the ID.
func (id ID) String() string {
	return id.value
}

// Equals checks if two IDs are the same.
func (id ID) Equals(other ID) bool {
	return id.value == other.value
}

// ValidationError is a custom error type for domain validation failures.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Msg)
}
