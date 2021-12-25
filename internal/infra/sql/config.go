package sql

import "fmt"

// Config defines the SQL connection configuration
type Config struct {
	URL string
}

// DatabaseURL returns the url prepared with its param values.
func (c *Config) DatabaseURL() string {
	return fmt.Sprint(c.URL)
}
