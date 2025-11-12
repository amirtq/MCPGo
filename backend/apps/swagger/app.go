// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// App represents the swagger application.
type App struct {
}

// NewApp creates a new swagger app.
func NewApp() *App {
	return &App{}
}

// SwaggerHandler returns the handler for swagger.
func (a *App) SwaggerHandler() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.json"), //The url pointing to API definition
	)
}
