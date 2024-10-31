package app

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"
	"github.com/go-chi/chi/v5"
)

// Тестирование конфигурации
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

// Тестирование создания приложения
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

// Тестирование настройки приложения
func TestApp_Setup(t *testing.T) {
	config := &Config{
		Host: "localhost",
		Port: "8080",
		DB: struct {
			DSN string
		}{
			DSN: "postgres://test:test@localhost:5432/testdb?sslmode=disable",
		},
	}

	app, err := New(context.Background(), config)
	if err != nil {
		t.Fatalf("Неожиданная ошибка при создании приложения: %v", err)
	}

	err = app.Setup(context.Background(), config.DB.DSN)
	if err != nil {
		t.Fatalf("Неожиданная ошибка при настройке приложения: %v", err)
	}

	// Проверяем, что все компоненты были инициализированы
	routes := []string{"/api/v1/products"}
	for _, route := range routes {
		if !routeExists(app.router, route) {
			t.Errorf("Маршрут %s не зарегистрирован", route)
		}
	}
}

// Вспомогательная функция для проверки существования маршрута
func routeExists(router *chi.Mux, path string) bool {
	found := false
	walkFn := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == path {
			found = true
		}
		return nil
	}
	chi.Walk(router, walkFn)
	return found
}

// Тестирование запуска приложения
func TestApp_Start(t *testing.T) {
	config := &Config{
		Host: "localhost",
		Port: "0", // Используем порт 0 для автоматического выбора свободного порта
		DB: struct {
			DSN string
		}{
			DSN: "postgres://test:test@localhost:5432/testdb?sslmode=disable",
		},
	}

	app, err := New(context.Background(), config)
	if err != nil {
		t.Fatalf("Неожиданная ошибка при создании приложения: %v", err)
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		if err := app.Start(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Неожиданная ошибка при запуске сервера: %v", err)
		}
	}()

	// Даем серверу время на запуск
	time.Sleep(100 * time.Millisecond)

	// Отправляем сигнал для graceful shutdown
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("Не удалось найти текущий процесс: %v", err)
	}
	
	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatalf("Не удалось отправить сигнал: %v", err)
	}

	// Даем серверу время на graceful shutdown
	time.Sleep(100 * time.Millisecond)
}
