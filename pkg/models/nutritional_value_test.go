package models

import (
	"testing"
)

func TestNutritionalValueAbsolute_Add(t *testing.T) {
	nv1 := NutritionalValueAbsolute{
		Proteins:      20,
		Fats:          10,
		Carbohydrates: 30,
		Calories:      300,
	}

	nv2 := NutritionalValueAbsolute{
		Proteins:      10,
		Fats:          5,
		Carbohydrates: 15,
		Calories:      150,
	}

	expected := NutritionalValueAbsolute{
		Proteins:      30,
		Fats:          15,
		Carbohydrates: 45,
		Calories:      450,
	}

	result := nv1.AddAbsoluteValue(nv2)

	// Check if the result matches the expected values
	if result.Proteins != expected.Proteins {
		t.Errorf("expected %d Proteins, got %d", expected.Proteins, result.Proteins)
	}
	if result.Fats != expected.Fats {
		t.Errorf("expected %d Fats, got %d", expected.Fats, result.Fats)
	}
	if result.Carbohydrates != expected.Carbohydrates {
		t.Errorf("expected %d Carbohydrates, got %d", expected.Carbohydrates, result.Carbohydrates)
	}
	if result.Calories != expected.Calories {
		t.Errorf("expected %d Calories, got %d", expected.Calories, result.Calories)
	}
}
