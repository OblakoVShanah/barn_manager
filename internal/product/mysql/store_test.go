package mysql_test

import (
	"context"
	"testing"
	"time"

	common "github.com/OblakoVShanah/barn_manager/internal/models"
	"github.com/OblakoVShanah/barn_manager/internal/product"
	"github.com/OblakoVShanah/barn_manager/internal/product/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// setupTestDB создает тестовое подключение к базе данных
func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("mysql", "barn_manager:barn_manager@tcp(localhost:3306)/barn_test?parseTime=true&loc=Local")
	require.NoError(t, err)
	
	// Очищаем таблицу перед тестами
	_, err = db.Exec("TRUNCATE TABLE products")
	require.NoError(t, err)
	
	return db
}

func TestStorage_SaveAndLoadProducts(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	storage := mysql.NewStorage(db)
	ctx := context.Background()

	testProduct := product.FoodProduct{
		ID:              "test-id",
		Name:            "Тестовый продукт",
		WeightPerPkg:    100,
		Amount:          1,
		PricePerPkg:     99.99,
		ExpirationDate:  time.Now().Add(24 * time.Hour).Truncate(time.Second),
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

	t.Run("обновление существующего продукта", func(t *testing.T) {
		updatedProduct := testProduct
		updatedProduct.Name = "Обновленный продукт"
		updatedProduct.Amount = 2

		id, err := storage.SaveProduct(ctx, updatedProduct)
		require.NoError(t, err)
		require.Equal(t, testProduct.ID, id)

		products, err := storage.LoadProducts(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, products)

		found := false
		for _, p := range products {
			if p.ID == testProduct.ID {
				found = true
				require.Equal(t, "Обновленный продукт", p.Name)
				require.Equal(t, uint(3), p.Amount)
				break
			}
		}
		require.True(t, found, "Обновленный продукт не найден")
	})

	t.Run("загрузка продуктов", func(t *testing.T) {
		products, err := storage.LoadProducts(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, products)

		found := false
		for _, p := range products {
			if p.ID == testProduct.ID {
				found = true
				require.Equal(t, "Обновленный продукт", p.Name)
				require.Equal(t, testProduct.WeightPerPkg, p.WeightPerPkg)
				require.Equal(t, testProduct.PricePerPkg, p.PricePerPkg)
				require.Equal(t, testProduct.ExpirationDate.Unix(), p.ExpirationDate.Unix())
				require.Equal(t, testProduct.PresentInFridge, p.PresentInFridge)
				require.Equal(t, testProduct.NutritionalValueRelative.Proteins, p.NutritionalValueRelative.Proteins)
				require.Equal(t, testProduct.NutritionalValueRelative.Fats, p.NutritionalValueRelative.Fats)
				require.Equal(t, testProduct.NutritionalValueRelative.Carbohydrates, p.NutritionalValueRelative.Carbohydrates)
				require.Equal(t, testProduct.NutritionalValueRelative.Calories, p.NutritionalValueRelative.Calories)
				break
			}
		}
		require.True(t, found, "Сохраненный продукт не найден")
	})
} 