package env

import "os"

func Get(key, default_ string) (string) {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return default_
}
