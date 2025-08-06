package product

import (
	"context"
	"time"

	common "github.com/OblakoVShanah/barn_manager/internal/models"
)

// FoodProduct represents a food product in the storage system.
// It contains all the necessary information about a product including
// its nutritional value, expiration date, and storage location.
type FoodProduct struct {
	ID                       string                          `json:"id"`                       // Unique product identifier
	Name                     string                          `json:"name"`                     // Product name
	WeightPerPkg             uint                            `json:"weight_per_pkg"`           // Weight per package in grams
	Amount                   uint                            `json:"amount"`                   // Number of packages available
	PricePerPkg              float32                         `json:"price_per_pkg"`            // Price per package
	ExpirationDate           time.Time                       `json:"expiration_date"`          // Product expiration date
	PresentInFridge          bool                            `json:"present_in_fridge"`        // Whether the product is stored in fridge
	NutritionalValueRelative common.NutritionalValueRelative `json:"nutritional_value_relative"` // Nutritional value per 100g
}

// ShoppingList represents a list of products needed for cooking a dish.
// It contains products that need to be purchased or are missing from the inventory.
type ShoppingList struct {
	Products []FoodProduct `json:"products"` // List of products needed
}

// Service defines the interface for product management operations.
// It provides methods for retrieving available products, adding new products,
// and checking product availability for recipes.
type Service interface {
	// AvailableProducts returns a list of all available products in the storage.
	// It retrieves all products that are currently in stock and not expired.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//
	// Returns:
	//   - []FoodProduct: List of available products
	//   - error: Any error that occurred during the operation
	AvailableProducts(ctx context.Context) ([]FoodProduct, error)

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
	PlaceProduct(ctx context.Context, product FoodProduct) (id string, err error)

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
	CheckAvailability(ctx context.Context, requirements map[string]uint) (ShoppingList, error)
}

// Store defines the interface for product storage operations.
// It provides methods for loading and saving products to the underlying storage system.
type Store interface {
	// LoadProducts loads all products from the storage.
	// It retrieves all products from the database, sorted by name.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//
	// Returns:
	//   - []FoodProduct: List of all products
	//   - error: Any error that occurred during the operation
	LoadProducts(ctx context.Context) ([]FoodProduct, error)

	// SaveProduct saves a product to the storage.
	// If a product with the same ID already exists, it will be updated.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - product: The product to be saved
	//
	// Returns:
	//   - string: The ID of the saved product
	//   - error: Any error that occurred during the operation
	SaveProduct(ctx context.Context, product FoodProduct) (id string, err error)
}
