// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package obs

import (
	"log"
	"os"
)

// Logger is a placeholder for a structured logger.
// In a real app, this would be an interface with implementations for different logging libraries.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger instance.
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[mcpgo] ", log.LstdFlags|log.Lshortfile),
	}
}

// For now, we just expose the standard logger methods.
// A real implementation would have leveled logging, e.g., Info, Warn, Error.
