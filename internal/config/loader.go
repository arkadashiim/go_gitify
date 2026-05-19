package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() (Config, error) {
	err := godotenv.Load()

	if err != nil {
		return Config{}, err
	}

	config := Config{
		RepoPath:       os.Getenv("GRAPH_REPO_PATH"),
		GitAuthorName:  os.Getenv("GIT_AUTHOR_NAME"),
		GitAuthorEmail: os.Getenv("GIT_AUTHOR_EMAIL"),
		SSHKeyPath:     os.Getenv("SSH_KEY_PATH"),
		GraphBranch:    envOrDefault("GRAPH_BRANCH", "main"),
	}

	if err := validate(config); err != nil {
		return Config{}, err
	} else {
		return config, nil
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
