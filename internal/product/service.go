package product

import (
	"context"
	"errors"
	"time"

	"github.com/OblakoVShanah/barn_manager/internal/oops"
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

// CheckAvailability проверяет доступность продуктов для покупки
func (s *AppService) CheckAvailability(ctx context.Context, requirements map[string]uint) (ShoppingList, error) {
	// Получаем все доступные продукты
	products, err := s.storage.LoadProducts(ctx)
	if err != nil && !errors.Is(err, oops.ErrNoData) {
		return ShoppingList{}, err
	}

	// Создаем мапу для быстрого поиска продуктов
	availableProducts := make(map[string]FoodProduct)
	for _, p := range products {
		if p.PresentInFridge {
			availableProducts[p.ID] = p
		}
	}

	// Проверяем каждый требуемый продукт
	var shoppingList ShoppingList
	for productID, requiredAmount := range requirements {
		product, exists := availableProducts[productID]
		if !exists {
			// Если продукт отсутствует, добавляем его в список покупок
			shoppingList.Products = append(shoppingList.Products, FoodProduct{
				ID:     productID,
				Amount: requiredAmount,
			})
			continue
		}

		// Проверяем, достаточно ли продукта
		availableAmount := product.Amount * product.WeightPerPkg
		if availableAmount < requiredAmount {
			// Если продукта недостаточно, добавляем недостающее количество в список покупок
			neededAmount := requiredAmount - availableAmount
			product.Amount = neededAmount
			shoppingList.Products = append(shoppingList.Products, product)
		}
	}

	return shoppingList, nil
}
