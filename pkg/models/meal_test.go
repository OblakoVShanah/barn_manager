package models

import (
	"testing"
	"time"
)

func TestMeal_AddIngredient(t *testing.T) {
	t.Parallel()
	meal := Meal{
		Id:            "1",
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

	meal.AddIngredient(ingredient)

	if len(meal.IngridientMap) != 1 {
		t.Errorf("expected 1 ingredient, got %d", len(meal.IngridientMap))
	}
}

func TestMeal_RemoveIngredient(t *testing.T) {
	t.Parallel()
	meal := Meal{
		Id:         "1",
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
}

func TestMeal_CalculateTotalPrice(t *testing.T) {
	t.Parallel()
	meal := Meal{
		Id:         "1",
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

	price := meal.CalculateTotalPrice()

	if price != 48 { // 23 * (100/100) + 50 * (200/400) = 48
		t.Errorf("expected price 48, got %f", price)
	}
}

func TestMeal_updateNutritionalValueAndWeight(t *testing.T) {
	t.Parallel()
	type fields struct {
		Id               string
		EatingTime       time.Time
		IngridientMap    map[string]Ingridient
		NutritionalValue NutritionalValueAbsolute
		Price            float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Meal{
				Id:               tt.fields.Id,
				EatingTime:       tt.fields.EatingTime,
				IngridientMap:    tt.fields.IngridientMap,
				NutritionalValue: tt.fields.NutritionalValue,
				Price:            tt.fields.Price,
			}
			m.updateNutritionalValueAndWeight()
		})
	}
}
