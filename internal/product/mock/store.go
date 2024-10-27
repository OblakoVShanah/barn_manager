package mock

import (
	"context"

	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
)

type Store struct {
	products []product.FoodProduct
	err      error
}

func NewStore() *Store {
	return &Store{
		products: make([]product.FoodProduct, 0),
	}
}

// SetError позволяет установить ошибку для тестирования сценариев с ошибками
func (s *Store) SetError(err error) {
	s.err = err
}

// SetProducts позволяет установить предопределенные продукты для тестирования
func (s *Store) SetProducts(products []product.FoodProduct) {
	s.products = products
}

func (s *Store) LoadProducts(ctx context.Context) ([]product.FoodProduct, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.products, nil
}

func (s *Store) SaveProduct(ctx context.Context, product product.FoodProduct) (id string, err error) {
	if s.err != nil {
		return "", s.err
	}
	s.products = append(s.products, product)
	return product.ID, nil
}
