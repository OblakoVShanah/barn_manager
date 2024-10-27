package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	// Добавьте другие поля конфигурации
}

// Load загружает конфигурацию из переменных окружения
func Load() (*Config, error) {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}, nil
}
