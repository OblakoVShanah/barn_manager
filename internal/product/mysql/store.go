package mysql

import (
	"context"
	"github.com/OblakoVShanah/barn_manager/internal/product"
	"github.com/OblakoVShanah/barn_manager/internal/oops"
	"github.com/jmoiron/sqlx"
)

// Storage представляет собой хранилище продуктов в MySQL базе данных
type Storage struct {
	db *sqlx.DB
}

// NewStorage создает новый экземпляр хранилища продуктов
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// SaveProduct сохраняет продукт в базу данных.
// Если продукт с таким ID уже существует, он будет обновлен.
// Возвращает ID сохраненного продукта или ошибку.
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
			price_per_pkg = (amount * price_per_pkg + VALUES(amount) * VALUES(price_per_pkg)) / (amount + VALUES(amount)),
			expiration_date = least(VALUES(expiration_date), expiration_date),
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

// LoadProducts загружает все продукты из базы данных.
// Возвращает список продуктов, отсортированный по имени, или ошибку.
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