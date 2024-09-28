package models

import (
	"testing"
)

func TestNutritionalValueAbsolute_Add(t *testing.T) {
	t.Parallel()
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
		t.Errorf("expected %f Proteins, got %f", expected.Proteins, result.Proteins)
	}
	if result.Fats != expected.Fats {
		t.Errorf("expected %f Fats, got %f", expected.Fats, result.Fats)
	}
	if result.Carbohydrates != expected.Carbohydrates {
		t.Errorf("expected %f Carbohydrates, got %f", expected.Carbohydrates, result.Carbohydrates)
	}
	if result.Calories != expected.Calories {
		t.Errorf("expected %f Calories, got %f", expected.Calories, result.Calories)
	}
}
