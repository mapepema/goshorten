package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration holds data necessary for capplication configuration
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Server holds data necessary for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Database holds data necessary for database configuration
type Database struct {
	LogQueries bool `yaml:"log_queries,omitempty"`
	Timeout    int  `yaml:"timeout,omitempty"`
}

// Application holds data necessary for application configuration
type Application struct {
	SwaggerUIPath string `yaml:"swagger_ui_path,omitempty"`
}

// Load returns Configuration structs
func Load(path string) (*Configuration, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("error while parsing config file, %s", err)
	}
	return cfg, nil
}
