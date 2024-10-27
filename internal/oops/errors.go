package oops

import (
	"errors"
	"fmt"
)

// Базовые ошибки
var (
	// Ошибки базы данных
	ErrNoData = errors.New("нет данных в базе данных")
	ErrDuplicateKey = errors.New("дублирование ключа")
	ErrDBConnection = errors.New("ошибка подключения к базе данных")

	// Ошибки бизнес-логики
	ErrProductNotFound = errors.New("продукт не найден")
	ErrInvalidProduct = errors.New("некорректные данные продукта")
	ErrExpiredProduct = errors.New("срок годности продукта истек")
)

// DBError представляет ошибку базы данных
type DBError struct {
	Err error
	ID  string
	Op  string
}

// Error возвращает строковое представление ошибки базы данных
func (e *DBError) Error() string {
	if e.ID != "" {
		return fmt.Sprintf("операция БД '%s' для ID '%s': %v", e.Op, e.ID, e.Err)
	}
	return fmt.Sprintf("операция БД '%s': %v", e.Op, e.Err)
}

// NewDBError создает новую ошибку базы данных
func NewDBError(err error, op string, id string) *DBError {
	return &DBError{
		Err: err,
		ID:  id,
		Op:  op,
	}
}

// ValidationError представляет ошибку валидации
type ValidationError struct {
	Field string
	Err   error
}

// Error возвращает строковое представление ошибки валидации
func (e *ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации поля '%s': %v", e.Field, e.Err)
}

// NewValidationError создает новую ошибку валидации
func NewValidationError(field string, err error) *ValidationError {
	return &ValidationError{
		Field: field,
		Err:   err,
	}
}

// Is реализует интерфейс errors.Is для проверки типов ошибок
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As реализует интерфейс errors.As для приведения типов ошибок
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
