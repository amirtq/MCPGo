package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.yaml.in/yaml/v3"
)

// Duration wraps time.Duration to provide YAML unmarshalling from strings.
type Duration struct {
	time.Duration
}

// UnmarshalYAML parses a duration string (e.g., "10s") into a time.Duration.
func (d *Duration) UnmarshalYAML(value *yaml.Node) error {
	var raw string
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if raw == "" {
		d.Duration = 0
		return nil
	}
	parsed, err := time.ParseDuration(raw)
	if err != nil {
		return err
	}
	d.Duration = parsed
	return nil
}

// AgentConfig controls how the agent-facing HTTP/WebSocket endpoints behave.
type AgentConfig struct {
	HTTP struct {
		Addr    string   `yaml:"addr"`
		Timeout Duration `yaml:"timeout"`
	} `yaml:"http"`
	WS struct {
		Addr string `yaml:"addr"`
	} `yaml:"ws"`
}

// ServerConfig defines an upstream MCP server.
type ServerConfig struct {
	ID       string `yaml:"id"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Protocol string `yaml:"protocol"`
}

// Config represents the full gateway configuration.
type Config struct {
	Agent   AgentConfig    `yaml:"agent"`
	Servers []ServerConfig `yaml:"servers"`
}

// Load reads configuration from the provided path. If the file does not exist,
// an error is returned.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("config file %s not found", path)
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config %s: %w", path, err)
	}

	return &cfg, nil
}

// LoadFromEnv loads configuration from the MCPGO_CONFIG environment variable.
// If the variable is empty, it attempts to load configs/config.yaml. When that
// file is missing it falls back to configs/config.example.yaml.
func LoadFromEnv() (*Config, string, error) {
	path := os.Getenv("MCPGO_CONFIG")
	candidates := []string{}
	if path != "" {
		candidates = append(candidates, path)
	}
	candidates = append(candidates,
		filepath.Join("configs", "config.yaml"),
		filepath.Join("configs", "config.example.yaml"),
	)

	var lastErr error
	for _, candidate := range candidates {
		cfg, err := Load(candidate)
		if err == nil {
			return cfg, candidate, nil
		}
		lastErr = err
	}
	if lastErr == nil {
		lastErr = errors.New("no configuration files found")
	}
	return nil, "", lastErr
}

// DefaultServer returns the first configured server or an error when none are defined.
func (c *Config) DefaultServer() (ServerConfig, error) {
	if len(c.Servers) == 0 {
		return ServerConfig{}, errors.New("no upstream servers configured")
	}
	return c.Servers[0], nil
}
