package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// EnvManagerConfig is a struct that implements the Config interface.
type EnvManagerConfig struct {
	config ConfigInfo
}

// Get returns the configuration information stored in the EnvManagerConfig instance.
func (f EnvManagerConfig) Get() ConfigInfo {
	return f.config
}

// NewEnvManagerConfig creates a new configuration using environment variables.
func NewEnvManagerConfig() (conf ConfigInfo, err error) {
	conf, err = newConfigWithEnvVars()
	return
}

// newConfigWithEnvVars initializes a new configuration using environment variables.
func newConfigWithEnvVars() (conf ConfigInfo, err error) {
	// Read server port from environment variable "PORT"
	srvPort, err := getEnvInt("PORT")
	if err != nil {
		return
	}

	// Read database port from environment variable "DB_PORT"
	dbPort, err := getEnvInt("DB_PORT")
	if err != nil {
		return
	}

	// Create a new ConfigInfo instance using environment variables
	conf = ConfigInfo{
		Server: server{
			Port:           srvPort,
			Host:           os.Getenv("SRV_HOST"),
			AllowedOrigins: strings.Split(os.Getenv("SRV_ALLOWED_ORIGINS"), ";"),
			SecretKey:      os.Getenv("SRV_SECRET_KEY"),
		},
		PostgreSQLProperties: postgreSQLProperties{
			URL:      os.Getenv("DATABASE_URL"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
		},
	}
	return
}

// getEnvInt retrieves an integer environment variable and converts it.
func getEnvInt(n string) (i int, err error) {
	i, err = strconv.Atoi(os.Getenv(n))
	if err != nil {
		err = fmt.Errorf("failed to load env var int %s: %s", n, err)
	}
	return
}
