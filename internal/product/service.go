package product

import (
	"context"
	"errors"
	"time"

	"github.com/OblakoVShanah/barn_manager/internal/oops"
)

// AppService implements the business logic for product management operations.
// It provides methods for retrieving available products, adding new products,
// and checking product availability for recipes.
type AppService struct {
	storage Store
}

// NewService creates a new instance of the product service.
// It initializes the service with the provided storage implementation.
//
// Parameters:
//   - storage: The storage implementation to use
//
// Returns:
//   - Service: The initialized service
func NewService(storage Store) Service {
	return &AppService{storage: storage}
}

// AvailableProducts returns a list of all available products in the storage.
// It retrieves all products that are currently in stock and not expired.
//
// Parameters:
//   - ctx: Context for the operation
//
// Returns:
//   - []FoodProduct: List of available products
//   - error: Any error that occurred during the operation
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

// PlaceProduct adds a new product to the storage.
// It validates the product data and stores it in the database.
// If a product with the same ID already exists, it will be updated.
//
// Parameters:
//   - ctx: Context for the operation
//   - product: The product to be added
//
// Returns:
//   - string: The ID of the saved product
//   - error: Any error that occurred during the operation
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

// CheckAvailability checks the availability of products for cooking a dish.
// It compares the required ingredients with available products and returns
// a shopping list of missing or insufficient products.
//
// Parameters:
//   - ctx: Context for the operation
//   - requirements: Map of product IDs to required amounts
//
// Returns:
//   - ShoppingList: List of products that need to be purchased
//   - error: Any error that occurred during the operation
func (s *AppService) CheckAvailability(ctx context.Context, requirements map[string]uint) (ShoppingList, error) {
	// Получаем все доступные продукты
	products, err := s.storage.LoadProducts(ctx)
	if err != nil && !errors.Is(err, oops.ErrNoData) {
		return ShoppingList{}, err
	}

	// Создаем мапу для быстрого поиска продуктов
	availableProducts := make(map[string]FoodProduct)
	for _, p := range products {
		availableProducts[p.ID] = p
	}

	// Проверяем каждый требуемый продукт
	var shoppingList ShoppingList
	for productID, requiredAmount := range requirements {
		availableProduct, exists := availableProducts[productID]
		if !exists {
			// Если продукт отсутствует, добавляем его в список покупок
			shoppingList.Products = append(shoppingList.Products, FoodProduct{
				ID:     productID,
				Amount: requiredAmount,
			})
			continue
		}

		// Проверяем, достаточно ли продукта
		if availableProduct.Amount < requiredAmount {
			// Если продукта недостаточно, добавляем недостающее количество в список покупок
			neededAmount := requiredAmount - availableProduct.Amount
			productToAdd := availableProduct
			productToAdd.Amount = neededAmount
			shoppingList.Products = append(shoppingList.Products, productToAdd)
		}
	}

	return shoppingList, nil
}
