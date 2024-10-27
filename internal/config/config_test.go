package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	// Подготавливаем тестовые переменные окружения
	t.Setenv("DATABASE_URL", "test-database-url")
	t.Setenv("SERVER_PORT", "8080")

	config, err := Load()
	if err != nil {
		t.Fatalf("Неожиданная ошибка при загрузке конфигурации: %v", err)
	}

	// Проверяем значения
	if config.DatabaseURL != "test-database-url" {
		t.Errorf("Ожидался DatabaseURL 'test-database-url', получено: %s", config.DatabaseURL)
	}

	if config.ServerPort != "8080" {
		t.Errorf("Ожидался ServerPort '8080', получено: %s", config.ServerPort)
	}
}
