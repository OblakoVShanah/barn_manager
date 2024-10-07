package models

import (
	"testing"
	"time"
)

func TestMeal_AddIngredient(t *testing.T) {
	t.Parallel()
	meal := Meal{
		ID:            "1",
		EatingTime:    time.Now(),
		IngridientMap: map[string]Ingridient{},
		Price:         0,
	}

	ingredient := Ingridient{
		FoodProduct: FoodProduct{
			Name:            "Chicken",
			WeightPerPkg:    100,
			Amount:          10,
			PricePerPkg:     123,
			ExpirationDate:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			PresentInFridge: true,
			NutritionalValueRelative: NutritionalValueRelative{
				Proteins:      25,
				Fats:          5,
				Carbohydrates: 0,
				Calories:      165,
			},
		},
		Weight: 200,
	}

	err := meal.AddIngredient(ingredient)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(meal.IngridientMap) != 1 {
		t.Errorf("expected 1 ingredient, got %d", len(meal.IngridientMap))
	}
}

func TestMeal_RemoveIngredient(t *testing.T) {
	t.Parallel()
	meal := Meal{
		ID:         "1",
		EatingTime: time.Now(),
		IngridientMap: map[string]Ingridient{
			"Chicken": {
				FoodProduct: FoodProduct{
					Name:            "Chicken",
					WeightPerPkg:    100,
					Amount:          10,
					PricePerPkg:     123,
					ExpirationDate:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					PresentInFridge: true,
					NutritionalValueRelative: NutritionalValueRelative{
						Proteins:      25,
						Fats:          5,
						Carbohydrates: 0,
						Calories:      165,
					},
				},
				Weight: 200,
			},
		},
		Price: 0,
	}

	err := meal.RemoveIngredient("Chicken")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(meal.IngridientMap) != 0 {
		t.Errorf("expected 0 ingredients, got %d", len(meal.IngridientMap))
	}

	err = meal.RemoveIngredient("Fish")
	if err == nil {
		t.Errorf("expected error: %v", ErrIngrNotFound)
	}
}

func TestMeal_CalculateTotalPrice(t *testing.T) {
	t.Parallel()
	meal := Meal{
		ID:         "1",
		EatingTime: time.Now(),
		IngridientMap: map[string]Ingridient{
			"Chicken": {
				FoodProduct: FoodProduct{
					Name:            "Chicken",
					WeightPerPkg:    100,
					Amount:          10,
					PricePerPkg:     23,
					ExpirationDate:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					PresentInFridge: true,
					NutritionalValueRelative: NutritionalValueRelative{
						Proteins:      25,
						Fats:          5,
						Carbohydrates: 0,
						Calories:      165,
					},
				},
				Weight: 100,
			},
			"Banana": {
				FoodProduct: FoodProduct{
					Name:            "Banana",
					WeightPerPkg:    400,
					Amount:          10,
					PricePerPkg:     50,
					ExpirationDate:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					PresentInFridge: true,
					NutritionalValueRelative: NutritionalValueRelative{
						Proteins:      25,
						Fats:          5,
						Carbohydrates: 0,
						Calories:      165,
					},
				},
				Weight: 200,
			},
		},
		Price: 0,
	}

	price, error := meal.CalculateTotalPrice()
	if error != nil {
		t.Errorf("unexpected error: %v", error)
	}

	if price != 48 { // 23 * (100/100) + 50 * (200/400) = 48
		t.Errorf("expected price 48, got %f", price)
	}

	meal.IngridientMap["Fish"] = Ingridient{
		FoodProduct: FoodProduct{
			Name:            "Fish",
			WeightPerPkg:    100,
			Amount:          10,
			PricePerPkg:     -1,
			ExpirationDate:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			PresentInFridge: true,
			NutritionalValueRelative: NutritionalValueRelative{
				Proteins:      25,
				Fats:          5,
				Carbohydrates: 0,
				Calories:      165,
			},
		},
		Weight: 200,
	}

	_, error = meal.CalculateTotalPrice()
	if error == nil {
		t.Errorf("expected error: %v", ErrNegativePrice)
	}
}
