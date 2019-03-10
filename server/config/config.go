package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration
type Config struct {
	// HTTPPort is the port the API will serve HTTP requests on
	HTTPPort int `envconfig:"http_port" required:"true"`

	// DBName is the database name
	DBName string `envconfig:"db_name" required:"true"`

	// DBUser is the database user
	DBUser string `envconfig:"db_user" required:"true"`
}

// NewConfigFromEnv load configuration options from environment variables
func NewConfigFromEnv() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("push_server", &cfg); err != nil {
		return nil, fmt.Errorf("error loading configuration from "+
			"environment variables: %s", err.Error())
	}

	return &cfg, nil
}
