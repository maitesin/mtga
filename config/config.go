package config

import (
	"github.com/maitesin/mtga/internal/infra/sql"
	"github.com/maitesin/mtga/internal/infra/storage"
	"os"
)

// Config defines the configuration of the mtga application
type Config struct {
	SQL     sql.Config
	Storage storage.Config
}

// NewConfig is the constructor for the mtga application configuration
func NewConfig() Config {
	return Config{
		SQL: sql.Config{
			URL: GetEnvOrDefault("DATABASE_URL", "cards.db"),
		},
		Storage: storage.Config{
			Path: GetEnvOrDefault("STORAGE_PATH", "devops/storage"),
		},
	}
}

func GetEnvOrDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}

	return defaultValue
}
