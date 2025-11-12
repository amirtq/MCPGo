package gateway

import (
	"context"
	"log"
	"net/http"
	"strings"

	gateway_app "mcpgo/backend/apps/gateway"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// Router wires HTTP/WebSocket requests to the gateway application.
type Router struct {
	app    *gateway_app.App
	logger *log.Logger
}

// NewRouter creates a new router for the gateway API.
func NewRouter(app *gateway_app.App, logger *log.Logger) *Router {
	if logger == nil {
		logger = log.Default()
	}
	return &Router{
		app:    app,
		logger: logger,
	}
}

// RegisterRoutes attaches the gateway routes to the provided mux.Router.
func (r *Router) RegisterRoutes(mux *mux.Router) {
	mux.Handle("/mcp", r.websocketHandler()).Methods(http.MethodGet)
}

func (r *Router) websocketHandler() http.Handler {
	return websocket.Server{
		Handshake: r.handshake,
		Handler:   r.handleWebSocket,
	}
}

func (r *Router) handshake(cfg *websocket.Config, req *http.Request) error {
	requested := selectSubprotocol(req.Header["Sec-Websocket-Protocol"])
	if requested == "" {
		requested = "mcp"
	}
	cfg.Protocol = []string{requested}
	if cfg.Origin == nil {
		if origin, err := websocket.Origin(cfg, req); err == nil {
			cfg.Origin = origin
		}
	}
	return nil
}

func (r *Router) handleWebSocket(conn *websocket.Conn) {
	var ctx = context.Background()
	if req := conn.Request(); req != nil {
		ctx = req.Context()
	}

	if err := r.app.HandleConnection(ctx, conn); err != nil {
		r.logger.Printf("gateway error: %v", err)
	}
}

func selectSubprotocol(headers []string) string {
	for _, header := range headers {
		parts := strings.Split(header, ",")
		for _, part := range parts {
			candidate := strings.TrimSpace(part)
			if candidate == "" {
				continue
			}
			if strings.EqualFold(candidate, "mcp") {
				return candidate
			}
			return candidate
		}
	}
	return ""
}
