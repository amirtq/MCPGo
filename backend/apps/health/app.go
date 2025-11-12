// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package health

// App represents the health check application.
type App struct {
	// In the future, we can add dependencies like a database connection
	// to perform more comprehensive health checks.
}

// NewApp creates a new HealthApp.
func NewApp() *App {
	return &App{}
}

// CheckHealth performs the health check and returns the status.
func (a *App) CheckHealth() map[string]string {
	return map[string]string{"status": "healthy"}
}
