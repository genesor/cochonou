package os

import (
	"fmt"
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

// MustGetEnv wraps the standard os.Getenv function by panicking
// if the value is empty.
func MustGetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("Env Variable %v must have a value", name))
	}

	return value
}
