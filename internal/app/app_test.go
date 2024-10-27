package app

import (
	"context"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	// Подготавливаем тестовые переменные окружения
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("DATABASE_URL", "postgres://test:test@localhost:5432/testdb")

	// Тестируем создание конфигурации
	config, err := NewConfig("")
	if err != nil {
		t.Fatalf("Неожиданная ошибка при создании конфигурации: %v", err)
	}

	// Проверяем значения
	if config.Port != "8080" {
		t.Errorf("Ожидался порт 8080, получено: %s", config.Port)
	}
	if config.DB.DSN != "postgres://test:test@localhost:5432/testdb" {
		t.Errorf("Неверный DSN, получено: %s", config.DB.DSN)
	}
}

func TestNew(t *testing.T) {
	// Подготавливаем тестовую конфигурацию
	config := &Config{
		Host: "localhost",
		Port: "8080",
	}

	// Тестируем создание приложения
	app, err := New(context.Background(), config)
	if err != nil {
		t.Fatalf("Неожиданная ошибка при создании приложения: %v", err)
	}

	// Проверяем инициализацию компонентов
	if app.router == nil {
		t.Error("Роутер не был инициализирован")
	}
	if app.http == nil {
		t.Error("HTTP сервер не был инициализирован")
	}
	if app.http.ReadTimeout != 15*time.Second {
		t.Error("Неверный таймаут чтения")
	}
}
