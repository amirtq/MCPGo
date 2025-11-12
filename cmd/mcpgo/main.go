// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

// @title MCPGo API
// @version 1.0
// @description This is a sample server for MCPGo.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
// @schemes https
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mcpgo/internal/application/commands"
	"mcpgo/internal/application/queries"
	"mcpgo/internal/infrastructure/connectors"
	"mcpgo/internal/infrastructure/eventbus"
	"mcpgo/internal/infrastructure/obs"
	"mcpgo/internal/infrastructure/persistence/memory"
	http_iface "mcpgo/internal/interfaces/http"
	"mcpgo/internal/platform/certs"

	_ "mcpgo/docs"
)

func main() {
	// 1. Initialize Infrastructure
	logger := obs.NewLogger()
	serverRepo := memory.NewInMemoryServerRepo()
	eventBus := eventbus.NewInMemoryEventBus(logger.Logger)
	connectorClient := connectors.NewMockConnectorClient(logger.Logger)

	// 2. Initialize Application Layer (CQRS Handlers)
	registerServerHandler := commands.NewRegisterServerHandler(serverRepo, eventBus)
	listServersHandler := queries.NewListServersHandler(serverRepo)
	routeCallHandler := commands.NewRouteCallHandler(serverRepo, connectorClient)

	// 3. Initialize Interface Layer (HTTP Handlers)
	handlers := &http_iface.Handlers{
		RegisterServer: registerServerHandler,
		ListServers:    listServersHandler,
		RouteCall:      routeCallHandler,
	}

	// 4. Create Router and Server
	router := http_iface.NewRouter(handlers)
	server := &http.Server{
		Addr:    ":443", // This would come from config in a real app
		Handler: router,
	}

	// 5. Ensure certificates are available and start server with Graceful Shutdown
	if err := certs.EnsureCerts(); err != nil {
		logger.Fatalf("Could not ensure certificates: %v\n", err)
	}

	go func() {
		logger.Println("Starting server on https://localhost" + server.Addr)
		if err := server.ListenAndServeTLS("configs/certs/cert.pem", "configs/certs/key.pem"); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exiting")
}
