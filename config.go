package config_go

import (
	"os"
)

// Get returns the value of the key.
// If the key is a file secret, it reads is and return the content.
// Otherwise it returns an empty string.
func Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}

	// check if value is a file secret
	content, err := os.ReadFile(value)
	if err == nil {
		return string(content)
	}

	return value
}
