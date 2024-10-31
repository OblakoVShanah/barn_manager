package postgres_test

import (
	"context"
	"testing"
	"time"

	common "github.com/OblakoVShanah/havchik_podbirator/internal/models"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

// setupTestDB создает тестовое подключение к базе данных
func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://test:test@localhost:5432/testdb?sslmode=disable") // TODO: заменить на переменные окружения
	require.NoError(t, err)
	return db
}

func TestStorage_SaveAndLoadProducts(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	storage := postgres.NewStorage(db)
	ctx := context.Background()

	testProduct := product.FoodProduct{
		ID:              "test-id",
		Name:            "Тестовый продукт",
		WeightPerPkg:    100,
		Amount:          1,
		PricePerPkg:     99.99,
		ExpirationDate:  time.Now().Add(24 * time.Hour),
		PresentInFridge: true,
		NutritionalValueRelative: common.NutritionalValueRelative{
			Proteins:      10,
			Fats:          5,
			Carbohydrates: 15,
			Calories:      150,
		},
	}

	t.Run("сохранение продукта", func(t *testing.T) {
		id, err := storage.SaveProduct(ctx, testProduct)
		require.NoError(t, err)
		require.Equal(t, testProduct.ID, id)
	})

	t.Run("загрузка продуктов", func(t *testing.T) {
		products, err := storage.LoadProducts(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, products)

		found := false
		for _, p := range products {
			if p.ID == testProduct.ID {
				found = true
				require.Equal(t, testProduct.Name, p.Name)
				require.Equal(t, testProduct.WeightPerPkg, p.WeightPerPkg)
				break
			}
		}
		require.True(t, found, "Сохраненный продукт не найден")
	})
}
