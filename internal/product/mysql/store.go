package mysql

import (
	"context"
	"github.com/OblakoVShanah/barn_manager/internal/product"
	"github.com/OblakoVShanah/barn_manager/internal/oops"
	"github.com/jmoiron/sqlx"
)

// Storage represents a MySQL database storage for products.
// It implements the Store interface and provides methods for saving and loading products.
type Storage struct {
	db *sqlx.DB
}

// NewStorage creates a new instance of MySQL storage.
// It initializes the storage with the provided database connection.
//
// Parameters:
//   - db: The database connection to use
//
// Returns:
//   - *Storage: The initialized storage
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// SaveProduct saves a product to the database.
// If a product with the same ID already exists, it will be updated with the new information.
// The price is recalculated as a weighted average based on the existing and new amounts.
//
// Parameters:
//   - ctx: Context for the operation
//   - product: The product to be saved
//
// Returns:
//   - string: The ID of the saved product
//   - error: Any error that occurred during the operation
func (s *Storage) SaveProduct(ctx context.Context, product product.FoodProduct) (string, error) {
	query := `
		INSERT INTO products (
			id, name, weight_per_pkg, amount, price_per_pkg, 
			expiration_date, present_in_fridge, 
			proteins, fats, carbohydrates, calories
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			name = VALUES(name),
			weight_per_pkg = VALUES(weight_per_pkg),
			amount = amount + VALUES(amount),
			price_per_pkg = CASE 
				WHEN amount + VALUES(amount) > 0 
				THEN (amount * price_per_pkg + VALUES(amount) * VALUES(price_per_pkg)) / (amount + VALUES(amount))
				ELSE VALUES(price_per_pkg)
			END,
			expiration_date = LEAST(VALUES(expiration_date), expiration_date),
			present_in_fridge = VALUES(present_in_fridge),
			proteins = VALUES(proteins),
			fats = VALUES(fats),
			carbohydrates = VALUES(carbohydrates),
			calories = VALUES(calories)
	`

	_, err := s.db.ExecContext(ctx, query,
		product.ID,
		product.Name,
		product.WeightPerPkg,
		product.Amount,
		product.PricePerPkg,
		product.ExpirationDate,
		product.PresentInFridge,
		product.NutritionalValueRelative.Proteins,
		product.NutritionalValueRelative.Fats,
		product.NutritionalValueRelative.Carbohydrates,
		product.NutritionalValueRelative.Calories,
	)

	if err != nil {
		return "", oops.NewDBError(err, "SaveProduct", product.ID)
	}

	return product.ID, nil
}

// LoadProducts loads all products from the database.
// It retrieves all products sorted by name and returns them as a slice.
//
// Parameters:
//   - ctx: Context for the operation
//
// Returns:
//   - []product.FoodProduct: List of all products
//   - error: Any error that occurred during the operation
func (s *Storage) LoadProducts(ctx context.Context) ([]product.FoodProduct, error) {
	query := `
		SELECT 
			id, name, weight_per_pkg, amount, price_per_pkg,
			expiration_date, present_in_fridge,
			proteins, fats, carbohydrates, calories
		FROM products
		ORDER BY name
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, oops.NewDBError(err, "LoadProducts", "")
	}
	defer rows.Close()

	var products []product.FoodProduct
	for rows.Next() {
		var p product.FoodProduct
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.WeightPerPkg,
			&p.Amount,
			&p.PricePerPkg,
			&p.ExpirationDate,
			&p.PresentInFridge,
			&p.NutritionalValueRelative.Proteins,
			&p.NutritionalValueRelative.Fats,
			&p.NutritionalValueRelative.Carbohydrates,
			&p.NutritionalValueRelative.Calories,
		)
		if err != nil {
			return nil, oops.NewDBError(err, "LoadProducts.Scan", "")
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, oops.NewDBError(err, "LoadProducts.Rows", "")
	}

	return products, nil
} 