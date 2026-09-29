package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig              `yaml':server`
	Providers map[string]ProviderConfig `yaml':providers`
}

type ServerConfig struct {
	Addr                   string `yaml':addr`
	ShutdownTimeoutSeconds int    `yaml':shutdown_timeout_seconds`
}

type ProviderConfig struct {
	ApiKeyEnv string `yaml':api_key_env`
	BaseUrl   string `yaml':base_url`
}

func (p ProviderConfig) APIKey() string {
	return os.Getenv(p.ApiKeyEnv)
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config file could not be read: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("config file could not be parsed: %w", err)
	}

	return &cfg, nil
}
