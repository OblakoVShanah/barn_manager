package product_test

import (
	"context"
	"errors"
	"testing"
	"time"

	common "github.com/OblakoVShanah/barn_manager/internal/models"
	"github.com/OblakoVShanah/barn_manager/internal/oops"
	"github.com/OblakoVShanah/barn_manager/internal/product"
	"github.com/OblakoVShanah/barn_manager/internal/product/mock"
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

func TestCheckAvailability(t *testing.T) {
	store := mock.NewStore()
	service := product.NewService(store)
	ctx := context.Background()

	// Setup test products in store
	availableProducts := []product.FoodProduct{
		{
			ID:              "milk",
			Name:            "Молоко",
			WeightPerPkg:    1000,
			Amount:          1000,
			PresentInFridge: true,
		},
		{
			ID:              "flour",
			Name:            "Мука",
			WeightPerPkg:    1000,
			Amount:          1500,
			PresentInFridge: true,
		},
	}
	store.SetProducts(availableProducts)

	t.Run("все продукты доступны", func(t *testing.T) {
		requirements := map[string]uint{
			"milk":  500,
			"flour": 250,
		}

		shoppingList, err := service.CheckAvailability(ctx, requirements)
		if err != nil {
			t.Fatalf("Неожиданная ошибка: %v", err)
		}

		if len(shoppingList.Products) != 0 {
			t.Errorf("Ожидался пустой список покупок, получено %d продуктов", len(shoppingList.Products))
		}
	})

	t.Run("недостаточно продуктов", func(t *testing.T) {
		requirements := map[string]uint{
			"milk":  3000, // 3 литра молока (больше чем есть)
			"flour": 250,  // 250 г муки
		}

		shoppingList, err := service.CheckAvailability(ctx, requirements)
		if err != nil {
			t.Fatalf("Неожиданная ошибка: %v", err)
		}

		if len(shoppingList.Products) != 1 {
			t.Fatalf("Ожидался 1 продукт в списке покупок, получено %d", len(shoppingList.Products))
		}

		if shoppingList.Products[0].ID != "milk" {
			t.Errorf("Ожидался продукт milk в списке покупок, получен %s", shoppingList.Products[0].ID)
		}
	})

	t.Run("продукт отсутствует", func(t *testing.T) {
		requirements := map[string]uint{
			"sugar": 100, // сахар отсутствует в хранилище
		}

		shoppingList, err := service.CheckAvailability(ctx, requirements)
		if err != nil {
			t.Fatalf("Неожиданная ошибка: %v", err)
		}

		if len(shoppingList.Products) != 1 {
			t.Fatalf("Ожидался 1 продукт в списке покупок, получено %d", len(shoppingList.Products))
		}

		if shoppingList.Products[0].ID != "sugar" {
			t.Errorf("Ожидался продукт sugar в списке покупок, получен %s", shoppingList.Products[0].ID)
		}
	})

	t.Run("ошибка хранилища", func(t *testing.T) {
		store.SetError(oops.ErrDBConnection)

		requirements := map[string]uint{
			"milk": 500,
		}

		_, err := service.CheckAvailability(ctx, requirements)
		if !errors.Is(err, oops.ErrDBConnection) {
			t.Errorf("Ожидалась ошибка %v, получено %v", oops.ErrDBConnection, err)
		}
	})
}
