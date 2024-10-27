package memory

import (
	"context"
	"sync"

	"github.com/OblakoVShanah/havchik_podbirator/internal/oops"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
)

// Storage реализует хранилище продуктов в памяти
type Storage struct {
	products map[string]product.FoodProduct
	mu       sync.RWMutex
}

// NewStorage создает новое хранилище в памяти
func NewStorage() *Storage {
	return &Storage{
		products: make(map[string]product.FoodProduct),
	}
}

// LoadProducts загружает все продукты из хранилища в памяти
func (s *Storage) LoadProducts(ctx context.Context) ([]product.FoodProduct, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]product.FoodProduct, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, oops.ErrNoData
	}
	return products, nil
}

// SaveProduct сохраняет продукт в хранилище в памяти
func (s *Storage) SaveProduct(ctx context.Context, product product.FoodProduct) (id string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.products[product.ID] = product
	return product.ID, nil
}
