package config

import (
	"log"
	"os"
)

func LoadEnvVar(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("key %q not found in env", key)
    }
    return value
}
