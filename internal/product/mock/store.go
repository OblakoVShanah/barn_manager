package mock

import (
	"context"
	"sync"

	"github.com/OblakoVShanah/havchik_podbirator/internal/oops"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
)

// Store представляет тестовое хранилище продуктов.
// Реализует интерфейс product.Store для использования в тестах.
type Store struct {
	products map[string]product.FoodProduct
	mu       sync.RWMutex
	err      error
}

// NewStore создает новый экземпляр тестового хранилища.
func NewStore() *Store {
	return &Store{
		products: make(map[string]product.FoodProduct),
	}
}

// SetError устанавливает ошибку, которая будет возвращена при следующем вызове методов хранилища.
// Используется для тестирования сценариев с ошибками.
func (s *Store) SetError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.err = err
}

// SetProducts устанавливает предопределенный набор продуктов в хранилище.
// Используется для подготовки тестовых данных.
func (s *Store) SetProducts(products []product.FoodProduct) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products = make(map[string]product.FoodProduct)
	for _, p := range products {
		s.products[p.ID] = p
	}
}

// LoadProducts возвращает все продукты из тестового хранилища.
// Если установлена ошибка через SetError, вернет её.
func (s *Store) LoadProducts(ctx context.Context) ([]product.FoodProduct, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.err != nil {
		return nil, s.err
	}

	products := make([]product.FoodProduct, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, oops.ErrNoData
	}
	return products, nil
}

// SaveProduct сохраняет продукт в тестовое хранилище.
// Если установлена ошибка через SetError, вернет её.
// При попытке сохранить продукт с существующим ID вернет oops.ErrDuplicateKey.
func (s *Store) SaveProduct(ctx context.Context, product product.FoodProduct) (id string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.err != nil {
		return "", s.err
	}

	if _, exists := s.products[product.ID]; exists {
		return "", oops.ErrDuplicateKey
	}

	s.products[product.ID] = product
	return product.ID, nil
}
