package gateway_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	gateway_api "mcpgo/backend/api/gateway"
	gateway_app "mcpgo/backend/apps/gateway"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

func TestGatewayProxiesMessages(t *testing.T) {
	upstream := httptest.NewServer(websocket.Server{
		Handshake: func(cfg *websocket.Config, req *http.Request) error {
			cfg.Protocol = []string{"mcp"}
			return nil
		},
		Handler: func(conn *websocket.Conn) {
			defer conn.Close()
			for {
				var message string
				if err := websocket.Message.Receive(conn, &message); err != nil {
					return
				}
				if err := websocket.Message.Send(conn, message); err != nil {
					return
				}
			}
		},
	})
	defer upstream.Close()

	upstreamURL := "ws" + strings.TrimPrefix(upstream.URL, "http")

	app, err := gateway_app.NewApp(upstreamURL, log.New(io.Discard, "", 0))
	if err != nil {
		t.Fatalf("failed to create app: %v", err)
	}

	router := mux.NewRouter()
	api := gateway_api.NewRouter(app, log.New(io.Discard, "", 0))
	api.RegisterRoutes(router)

	gatewayServer := httptest.NewServer(router)
	defer gatewayServer.Close()

	gatewayURL := "ws" + strings.TrimPrefix(gatewayServer.URL, "http") + "/mcp"
	conn, err := websocket.Dial(gatewayURL, "mcp", "http://localhost")
	if err != nil {
		t.Fatalf("failed to dial gateway: %v", err)
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("failed to set deadline: %v", err)
	}

	payload := `{"jsonrpc":"2.0","id":1,"method":"ping"}`
	if err := websocket.Message.Send(conn, payload); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	var reply string
	if err := websocket.Message.Receive(conn, &reply); err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	if reply != payload {
		t.Fatalf("expected reply %q, got %q", payload, reply)
	}
}
