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

// @BasePath /
// @schemes https
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	gateway_api "mcpgo/backend/api/gateway"
	health_api "mcpgo/backend/api/health"
	swagger_api "mcpgo/backend/api/swagger"
	"mcpgo/backend/apps/gateway"
	"mcpgo/backend/apps/health"
	swagger_app "mcpgo/backend/apps/swagger"
	"mcpgo/backend/services/config"
	"mcpgo/backend/services/ssl"

	"github.com/gorilla/mux"
)

func main() {
	swagger_app.SwaggerInfo.Host = "localhost:443"
	// 1. Initialize Infrastructure
	logger := log.Default()

	cfg, configPath, err := config.LoadFromEnv()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
	logger.Printf("loaded configuration from %s", configPath)

	serverCfg, err := cfg.DefaultServer()
	if err != nil {
		logger.Fatalf("no upstream server configured: %v", err)
	}

	gatewayApp, err := gateway.NewApp(serverCfg.Address, logger)
	if err != nil {
		logger.Fatalf("failed to create gateway app: %v", err)
	}

	// 4. Create Router and Server
	router := mux.NewRouter()

	// Initialize and register apps
	healthApp := health.NewApp()
	healthAPI := health_api.NewRouter(healthApp)
	healthAPI.RegisterRoutes(router)

	swaggerApp := swagger_app.NewApp()
	swaggerAPI := swagger_api.NewRouter(swaggerApp)
	swaggerAPI.RegisterRoutes(router)

	gatewayAPI := gateway_api.NewRouter(gatewayApp, logger)
	gatewayAPI.RegisterRoutes(router)

	httpAddr := cfg.Agent.HTTP.Addr
	if httpAddr == "" {
		httpAddr = cfg.Agent.WS.Addr
	}
	if httpAddr == "" {
		httpAddr = ":443"
	}

	hostPort := httpAddr
	if strings.HasPrefix(hostPort, ":") {
		hostPort = "localhost" + hostPort
	}
	swagger_app.SwaggerInfo.Host = hostPort
	httpTimeout := cfg.Agent.HTTP.Timeout.Duration
	if httpTimeout <= 0 {
		httpTimeout = 10 * time.Second
	}

	server := &http.Server{
		Addr:              httpAddr,
		Handler:           router,
		ReadHeaderTimeout: httpTimeout,
		WriteTimeout:      httpTimeout,
		IdleTimeout:       httpTimeout,
	}

	// 5. Ensure certificates are available and start server with Graceful Shutdown
	if err := ssl.EnsureSSL(); err != nil {
		logger.Fatalf("Could not ensure certificates: %v\n", err)
	}

	go func() {
		logger.Println("Starting MCP gateway on https://localhost" + server.Addr)
		if err := server.ListenAndServeTLS("backend/services/ssl/cert.pem", "backend/services/ssl/key.pem"); err != nil && err != http.ErrServerClosed {
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
