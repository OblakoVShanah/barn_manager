package product

import (
	"context"
	"errors"
	"time"

	"github.com/OblakoVShanah/havchik_podbirator/internal/oops"
)

// AppService реализует бизнес-логику работы с продуктами
type AppService struct {
	storage Store
}

// NewService создает новый экземпляр сервиса
func NewService(storage Store) Service {
	return &AppService{storage: storage}
}

// AvailableProducts возвращает список всех доступных продуктов
func (s *AppService) AvailableProducts(ctx context.Context) ([]FoodProduct, error) {
	products, err := s.storage.LoadProducts(ctx)
	if err != nil {
		if errors.Is(err, oops.ErrNoData) {
			return nil, err
		}
		return nil, err
	}
	return products, nil
}

// PlaceProduct добавляет новый продукт в хранилище
func (s *AppService) PlaceProduct(ctx context.Context, product FoodProduct) (id string, err error) {
	// Проверка валидности продукта
	if product.ID == "" || product.Name == "" || product.WeightPerPkg == 0 {
		return "", oops.ErrInvalidProduct
	}

	// Существующая проверка срока годности
	if time.Now().After(product.ExpirationDate) {
		return "", oops.ErrExpiredProduct
	}

	id, err = s.storage.SaveProduct(ctx, product)
	if err != nil {
		return "", err
	}
	return id, nil
}
