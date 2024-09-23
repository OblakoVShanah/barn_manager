package models

import (
	"testing"
	"time"
)

func TestAddIngredient(t *testing.T) {
	meal := Meal{
		Id:          "1",
		EatingTime:  time.Now(),
		Ingridients: []Ingridient{},
		Price:       0,
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

	meal.AddIngredient(ingredient)

	if len(meal.Ingridients) != 1 {
		t.Errorf("expected 1 ingredient, got %d", len(meal.Ingridients))
	}
}

func TestRemoveIngredient(t *testing.T) {
	meal := Meal{
		Id:         "1",
		EatingTime: time.Now(),
		Ingridients: []Ingridient{{
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
			Weight: 200}},
		Price: 0,
	}

	err := meal.RemoveIngredient(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(meal.Ingridients) != 0 {
		t.Errorf("expected 0 ingredients, got %d", len(meal.Ingridients))
	}
}

func TestCalculateTotalPrice(t *testing.T) {
	meal := Meal{
		Id:         "1",
		EatingTime: time.Now(),
		Ingridients: []Ingridient{
			{
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
				Weight: 100},

			{
			FoodProduct: FoodProduct{
				Name:            "Chicken",
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
			Weight: 200},
		},
	}

	price := meal.CalculateTotalPrice()

	if price != 48 { // 23 * (100/100) + 50 * (200/400) = 48
		t.Errorf("expected price 48, got %f", price)
	}
}
