package os

import (
	"os"
)

// GetEnvWithDefault wraps the standard os.Getenv function by adding a default value.
func GetEnvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}
