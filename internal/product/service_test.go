package product_test

import (
	"context"
	"errors"
	"testing"
	"time"

	common "github.com/OblakoVShanah/havchik_podbirator/internal/models"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product"
	"github.com/OblakoVShanah/havchik_podbirator/internal/product/mock"
	"github.com/OblakoVShanah/havchik_podbirator/internal/oops"
)

func TestAvailableProducts(t *testing.T) {
	// Подготавливаем мок хранилища
	store := mock.NewStore()
	service := product.NewService(store)

	// Тестовые продукты
	testProducts := []product.FoodProduct{
		{
			ID:              "1",
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
		},
	}

	t.Run("успешное получение продуктов", func(t *testing.T) {
		store.SetProducts(testProducts)

		products, err := service.AvailableProducts(context.Background())
		if err != nil {
			t.Fatalf("Неожиданная ошибка: %v", err)
		}

		if len(products) != len(testProducts) {
			t.Errorf("Ожидалось %d продуктов, получено %d", len(testProducts), len(products))
		}
	})

	t.Run("ошибка при получении продуктов", func(t *testing.T) {
		expectedErr := oops.ErrNoData
		store.SetError(expectedErr)

		_, err := service.AvailableProducts(context.Background())
		if err != expectedErr {
			t.Errorf("Ожидалась ошибка %v, получено %v", expectedErr, err)
		}
	})
}

func TestPlaceProduct(t *testing.T) {
	store := mock.NewStore()
	service := product.NewService(store)

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

	t.Run("успешное добавление продукта", func(t *testing.T) {
		id, err := service.PlaceProduct(context.Background(), testProduct)
		if err != nil {
			t.Fatalf("Неожиданная ошибка: %v", err)
		}

		if id != testProduct.ID {
			t.Errorf("Ожидался ID %s, получен %s", testProduct.ID, id)
		}
	})

	t.Run("ошибка при добавлении продукта", func(t *testing.T) {
		expectedErr := errors.New("ошибка сохранения")
		store.SetError(expectedErr)

		_, err := service.PlaceProduct(context.Background(), testProduct)
		if err != expectedErr {
			t.Errorf("Ожидалась ошибка %v, получено %v", expectedErr, err)
		}
	})
}
