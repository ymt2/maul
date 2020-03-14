package config

import (
	"log"
	"os"
)

// Env represents configuration should be defined by environment variables.
type Env struct {
	GitHubAuthToken string
}

// ReadFromEnv retrieves configuration from environment variables.
func ReadFromEnv() *Env {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	return &Env{
		GitHubAuthToken: token,
	}
}
