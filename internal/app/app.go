package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/OblakoVShanah/barn_manager/internal/product"
	// "github.com/OblakoVShanah/barn_manager/internal/product/postgres"
	"github.com/OblakoVShanah/barn_manager/internal/product/mysql"
	// "github.com/OblakoVShanah/barn_manager/internal/product/memory"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql" // драйвер mysql
	"github.com/jmoiron/sqlx"

	// _ "github.com/lib/pq" // драйвер postgres
	"gopkg.in/yaml.v3"
)

// Config represents the application configuration structure.
// It contains settings for the HTTP server and database connection.
type Config struct {
	Host string `yaml:"host"` // HTTP server host
	Port string `yaml:"port"` // HTTP server port
	DB   struct {
		DSN string `yaml:"dsn"` // Database connection string
	} `yaml:"db"`
}

// NewConfig creates a new application configuration from a YAML file.
// It reads the configuration file from the specified path and unmarshals it into a Config struct.
//
// Parameters:
//   - configPath: Path to the YAML configuration file
//
// Returns:
//   - *Config: The parsed configuration
//   - error: Any error that occurred during file reading or parsing
func NewConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

// App represents the main application structure.
// It contains the HTTP server, router, and configuration.
type App struct {
	config *Config
	router *chi.Mux
	http   *http.Server
}

// New creates a new application instance with the provided configuration.
// It initializes the HTTP server with reasonable default timeouts and the Chi router.
//
// Parameters:
//   - ctx: Context for the operation
//   - config: Application configuration
//
// Returns:
//   - *App: The initialized application
//   - error: Any error that occurred during initialization
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

// Setup initializes the application components including database connection,
// storage layer, service layer, and HTTP handlers.
//
// Parameters:
//   - ctx: Context for the operation
//   - dsn: Database connection string
//
// Returns:
//   - error: Any error that occurred during setup
func (a *App) Setup(ctx context.Context, dsn string) error {
	// // Инициализация подключения к базе данных
	// db, err := sqlx.Connect("mysql", "barn_manager:barn_manager@tcp(localhost:3306)/barn?parseTime=true&loc=Local")
	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	// db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	// Тестирование подключения
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("не удалось выполнить ping базы данных: %w", err)
	}

	// Инициализация хранилища
	// store := postgres.NewStorage(db)
	store := mysql.NewStorage(db)
	// store := memory.NewStorage()

	// Инициализация сервиса
	service := product.NewService(store)

	// Инициализация и регистрация обработчиков
	handler := product.NewHandler(a.router, service)
	handler.Register()

	return nil
}

// Start runs the HTTP server with graceful shutdown handling.
// It starts the server in a goroutine and waits for an interrupt signal.
// When a signal is received, it gracefully shuts down the server.
//
// Returns:
//   - error: Any error that occurred during server operation or shutdown
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
