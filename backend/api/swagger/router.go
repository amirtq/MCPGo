// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package swagger

import (
	"mcpgo/backend/apps/swagger"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Router is the router for the swagger app.
type Router struct {
	app *swagger.App
}

// NewRouter creates a new swagger router.
func NewRouter(app *swagger.App) *Router {
	return &Router{
		app: app,
	}
}

// RegisterRoutes registers the routes for the swagger app.
func (r *Router) RegisterRoutes(mux *mux.Router) {
	mux.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
