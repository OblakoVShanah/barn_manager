package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product/memory"
	// "github.com/OblakoVShanah/havchik_podbirator/internal/product/postgres"
	"github.com/go-chi/chi/v5"
	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер postgres
)

// Config представляет конфигурацию приложения
type Config struct {
	Host string
	Port string
	DB   struct {
		DSN string
	}
}

// NewConfig создает конфигурацию приложения из переменных окружения
func NewConfig(configPath string) (*Config, error) {
	return &Config{
		Host: "0.0.0.0",
		Port: os.Getenv("SERVER_PORT"),
		DB: struct {
			DSN string
		}{
			DSN: os.Getenv("DATABASE_URL"),
		},
	}, nil
}

// App это структура приложения
type App struct {
	config *Config
	router *chi.Mux
	http   *http.Server
}

// New создает новое приложение
func New(ctx context.Context, config *Config) (*App, error) {
	r := chi.NewRouter()

	return &App{
		config: config,
		router: r,
		http: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
			Handler: r,
			// Разумные значения по умолчанию для таймаутов
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       30 * time.Second,
		},
	}, nil
}

// Setup инициализирует приложение
func (a *App) Setup(ctx context.Context, dsn string) error {
	// // Инициализация подключения к базе данных
	// db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	// if err != nil {
	// 	return fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	// }

	// // Тестирование подключения
	// if err := db.PingContext(ctx); err != nil {
	// 	return fmt.Errorf("не удалось выполнить ping базы данных: %w", err)
	// }

	// Инициализация хранилища
	// store := postgres.NewStorage(db)
	store := memory.NewStorage()

	// Инициализация сервиса
	service := product.NewService(store)

	// Инициализация и регистрация обработчиков
	handler := product.NewHandler(a.router, service)
	handler.Register()

	return nil
}

// Start запускает приложение
func (a *App) Start() error {
	// Создание контекста, который будет отменен при получении сигнала прерывания
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Запуск сервера в горутине
	go func() {
		log.Printf("запуск веб-сервера на %s", a.http.Addr)
		if err := a.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("не удалось запустить сервер: %v", err)
		}
	}()

	// Ожидание сигнала прерывания
	<-ctx.Done()

	// Восстановление стандартного поведения при получении сигнала прерывания и уведомление пользователя о завершении работы
	stop()
	log.Println("плавное завершение работы, нажмите Ctrl+C еще раз для принудительного завершения")

	// Создание дедлайна для ожидания завершения
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Завершение работы сервера
	if err := a.http.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("не удалось завершить работу сервера: %w", err)
	}

	log.Println("сервер успешно завершил работу")
	return nil
}
