package memory

import (
	"context"
	"github.com/OblakoVShanah/havchik_podbirator/internal/barn"
	"sync"
)

type Storage struct {
	products map[string]barn.FoodProduct
	mu       sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		products: make(map[string]barn.FoodProduct),
	}
}

func (s *Storage) LoadProducts(ctx context.Context) ([]barn.FoodProduct, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]barn.FoodProduct, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) SaveProduct(ctx context.Context, product barn.FoodProduct) (id string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.products[product.ID] = product
	return product.ID, nil
}
