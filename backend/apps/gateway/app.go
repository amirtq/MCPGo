package gateway

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

// rawMessage captures both the payload bytes and the frame type so that we can
// faithfully forward text and binary frames.
type rawMessage struct {
	Data        []byte
	PayloadType byte
}

var rawCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		msg, ok := v.(rawMessage)
		if !ok {
			return nil, websocket.UnknownFrame, fmt.Errorf("unexpected type %T", v)
		}
		return msg.Data, msg.PayloadType, nil
	},
	Unmarshal: func(data []byte, payloadType byte, v interface{}) error {
		msg, ok := v.(*rawMessage)
		if !ok {
			return fmt.Errorf("unexpected type %T", v)
		}
		msg.PayloadType = payloadType
		msg.Data = append(msg.Data[:0], data...)
		return nil
	},
}

// App encapsulates the MCP gateway logic. It maintains a connection to an
// upstream MCP server and proxies messages between the client and the upstream
// service.
type App struct {
	upstreamURL *url.URL
	baseConfig  *websocket.Config
	dialTimeout time.Duration
	logger      *log.Logger
}

// NewApp creates a new gateway app for the provided upstream address. The
// address must be a valid WebSocket URL.
func NewApp(upstreamAddress string, logger *log.Logger) (*App, error) {
	if upstreamAddress == "" {
		return nil, errors.New("upstream address is required")
	}
	parsed, err := url.Parse(upstreamAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid upstream address %q: %w", upstreamAddress, err)
	}
	if parsed.Scheme != "ws" && parsed.Scheme != "wss" {
		return nil, fmt.Errorf("unsupported upstream scheme %q", parsed.Scheme)
	}
	if logger == nil {
		logger = log.Default()
	}

	originScheme := "http"
	if parsed.Scheme == "wss" {
		originScheme = "https"
	}
	origin := fmt.Sprintf("%s://%s", originScheme, parsed.Host)
	baseConfig, err := websocket.NewConfig(upstreamAddress, origin)
	if err != nil {
		return nil, fmt.Errorf("failed to build upstream config: %w", err)
	}
	baseConfig.Protocol = []string{"mcp"}
	baseConfig.Dialer = &net.Dialer{Timeout: 10 * time.Second}

	return &App{
		upstreamURL: parsed,
		baseConfig:  baseConfig,
		dialTimeout: 10 * time.Second,
		logger:      logger,
	}, nil
}

// HandleConnection establishes a new upstream connection and proxies all MCP
// traffic between the connected client and the upstream server.
func (a *App) HandleConnection(ctx context.Context, clientConn *websocket.Conn) error {
	if clientConn == nil {
		return errors.New("client connection is nil")
	}

	subproto := "mcp"
	if cfg := clientConn.Config(); cfg != nil && len(cfg.Protocol) > 0 {
		if cfg.Protocol[0] != "" {
			subproto = cfg.Protocol[0]
		}
	}

	clientAddr := "unknown"
	if req := clientConn.Request(); req != nil {
		clientAddr = req.RemoteAddr
	}
	a.logger.Printf("Connecting client %s to upstream %s using protocol %s", clientAddr, a.upstreamURL, subproto)

	upstreamConfig := cloneConfig(a.baseConfig)
	upstreamConfig.Protocol = []string{subproto}
	upstreamConfig.Header = cloneHeader(a.baseConfig.Header)
	if upstreamConfig.Dialer == nil {
		upstreamConfig.Dialer = &net.Dialer{}
	}
	upstreamConfig.Dialer.Timeout = a.dialTimeout

	upstreamConn, err := upstreamConfig.DialContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to upstream %s: %w", a.upstreamURL, err)
	}
	defer upstreamConn.Close()
	defer clientConn.Close()

	errCh := make(chan error, 2)
	var once sync.Once
	closeBoth := func() {
		once.Do(func() {
			_ = clientConn.Close()
			_ = upstreamConn.Close()
		})
	}

	go func() {
		errCh <- pump("client->upstream", clientConn, upstreamConn)
	}()
	go func() {
		errCh <- pump("upstream->client", upstreamConn, clientConn)
	}()

	select {
	case <-ctx.Done():
		closeBoth()
		return ctx.Err()
	case err := <-errCh:
		closeBoth()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if errors.Is(err, context.Canceled) {
			return nil
		}
		return err
	}
}

func pump(direction string, src, dst *websocket.Conn) error {
	for {
		var msg rawMessage
		if err := rawCodec.Receive(src, &msg); err != nil {
			return fmt.Errorf("%s receive: %w", direction, err)
		}
		if err := rawCodec.Send(dst, msg); err != nil {
			return fmt.Errorf("%s send: %w", direction, err)
		}
	}
}

func cloneConfig(cfg *websocket.Config) *websocket.Config {
	if cfg == nil {
		return &websocket.Config{}
	}
	copyCfg := *cfg
	copyCfg.Header = cloneHeader(cfg.Header)
	if cfg.Dialer != nil {
		dialerCopy := *cfg.Dialer
		copyCfg.Dialer = &dialerCopy
	}
	return &copyCfg
}

func cloneHeader(h http.Header) http.Header {
	if h == nil {
		return nil
	}
	clone := make(http.Header, len(h))
	for k, values := range h {
		vv := make([]string, len(values))
		copy(vv, values)
		clone[k] = vv
	}
	return clone
}
