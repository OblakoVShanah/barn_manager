package oops

import (
	"errors"
	"fmt"
)

// Base errors for common application scenarios
var (
	// Database errors
	ErrNoData       = errors.New("нет данных в базе данных")     // No data found in database
	ErrDuplicateKey = errors.New("дублирование ключа")           // Duplicate key error
	ErrDBConnection = errors.New("ошибка подключения к базе данных") // Database connection error

	// Business logic errors
	ErrProductNotFound = errors.New("продукт не найден")         // Product not found
	ErrInvalidProduct  = errors.New("некорректные данные продукта") // Invalid product data
	ErrExpiredProduct  = errors.New("срок годности продукта истек") // Product has expired
)

// DBError represents a database operation error.
// It provides additional context about the database operation that failed,
// including the operation type and the ID of the affected record.
type DBError struct {
	Err error  // The underlying error
	ID  string // The ID of the affected record (if applicable)
	Op  string // The operation that failed
}

// Error returns a string representation of the database error.
// It includes information about the operation and the affected record ID if available.
//
// Returns:
//   - string: A formatted error message
func (e *DBError) Error() string {
	if e.ID != "" {
		return fmt.Sprintf("операция БД '%s' для ID '%s': %v", e.Op, e.ID, e.Err)
	}
	return fmt.Sprintf("операция БД '%s': %v", e.Op, e.Err)
}

// NewDBError creates a new database error with the specified details.
// It wraps the underlying error with additional context about the operation.
//
// Parameters:
//   - err: The underlying error
//   - op: The operation that failed
//   - id: The ID of the affected record (optional)
//
// Returns:
//   - *DBError: The new database error
func NewDBError(err error, op string, id string) *DBError {
	return &DBError{
		Err: err,
		ID:  id,
		Op:  op,
	}
}

// ValidationError represents a validation error for a specific field.
// It provides information about which field failed validation and why.
type ValidationError struct {
	Field string // The field that failed validation
	Err   error  // The underlying validation error
}

// Error returns a string representation of the validation error.
// It includes information about the field that failed validation.
//
// Returns:
//   - string: A formatted error message
func (e *ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации поля '%s': %v", e.Field, e.Err)
}

// NewValidationError creates a new validation error for the specified field.
// It wraps the underlying error with field-specific context.
//
// Parameters:
//   - field: The field that failed validation
//   - err: The underlying validation error
//
// Returns:
//   - *ValidationError: The new validation error
func NewValidationError(field string, err error) *ValidationError {
	return &ValidationError{
		Field: field,
		Err:   err,
	}
}

// Is implements the errors.Is interface for checking error types.
// It provides a convenient way to check if an error is of a specific type.
//
// Parameters:
//   - err: The error to check
//   - target: The target error to compare against
//
// Returns:
//   - bool: True if the error matches the target
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As implements the errors.As interface for type assertion.
// It provides a convenient way to extract specific error types.
//
// Parameters:
//   - err: The error to check
//   - target: The target error type to extract
//
// Returns:
//   - bool: True if the error can be converted to the target type
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
