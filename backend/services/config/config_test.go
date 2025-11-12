package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"mcpgo/backend/services/config"
)

func TestLoadParsesServerConfiguration(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	content := []byte("agent:\n  http:\n    addr: \":9000\"\n    timeout: 5s\nservers:\n  - id: \"test\"\n    name: \"Test Server\"\n    address: \"ws://localhost:1234/mcp\"\n    protocol: \"mcp/v1\"\n")
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	cfg, err := config.Load(path)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Agent.HTTP.Addr != ":9000" {
		t.Fatalf("expected http addr :9000, got %q", cfg.Agent.HTTP.Addr)
	}

	server, err := cfg.DefaultServer()
	if err != nil {
		t.Fatalf("unexpected error from DefaultServer: %v", err)
	}
	if server.Address != "ws://localhost:1234/mcp" {
		t.Fatalf("unexpected server address %q", server.Address)
	}

	t.Setenv("MCPGO_CONFIG", path)
	loaded, loadedPath, err := config.LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv returned error: %v", err)
	}
	if loadedPath != path {
		t.Fatalf("expected loaded path %q, got %q", path, loadedPath)
	}
	if loaded.Agent.HTTP.Addr != cfg.Agent.HTTP.Addr {
		t.Fatalf("expected loaded config http addr %q, got %q", cfg.Agent.HTTP.Addr, loaded.Agent.HTTP.Addr)
	}
}
