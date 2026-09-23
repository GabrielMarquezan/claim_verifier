package config

import (
	"os"
)

type Config struct {
	DatabaseURL       string
	MigatrionsDirPath string
	RAGWorkerAddr     string
}

func Load() *Config {
	return &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		MigatrionsDirPath: os.Getenv("MIGRATIONS_DIR_PATH"),
		RAGWorkerAddr:     os.Getenv("RAG_WORKER_ADDR"),
	}
}
