package models

import (
	"errors"
	"time"
)

var (
	ErrIngrNotFound          = errors.New("ingredient not found")
	ErrNegativeWeight        = errors.New("negative weight")
	ErrNegativePrice         = errors.New("negative price")
	ErrNegativeCalories      = errors.New("negative calories")
	ErrNegativeProteins      = errors.New("negative proteins")
	ErrNegativeFats          = errors.New("negative fats")
	ErrNegativeCarbohydrates = errors.New("negative carbohydrates")
)

type Meal struct {
	Id               string
	EatingTime       time.Time
	IngridientMap    map[string]Ingridient // Using map instead of slice
	NutritionalValue NutritionalValueAbsolute
	Price            float32
}

type Ingridient struct {
	FoodProduct FoodProduct
	Weight      int
}

func (ingredient *Ingridient) ValidateIngridient() error {
	if ingredient.Weight < 0 {
		return ErrNegativeWeight
	}
	if error := ingredient.FoodProduct.ValidateFoodproduct() ; error != nil {
		return error
	}
	
	return nil
}

func (m *Meal) AddIngredient(ingredient Ingridient) error {
	// Validate ingridient
	if error := ingredient.ValidateIngridient(); error != nil {
		return error
	}

	// Initialize the map if it's nil
	if m.IngridientMap == nil {
		m.IngridientMap = make(map[string]Ingridient)
	}

	// Add or update the ingredient in the map
	m.IngridientMap[ingredient.FoodProduct.Name] = ingredient

	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) RemoveIngredient(name string) error {
	// Check if ingredient exists
	if _, exists := m.IngridientMap[name]; !exists {
		return ErrIngrNotFound
	}

	// Remove from map
	delete(m.IngridientMap, name)

	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) UpdateIngredientWeight(name string, newWeight int) error {
	// Lookup ingredient in map
	ingredient, exists := m.IngridientMap[name]
	if !exists {
		return ErrIngrNotFound
	}

	// Check for negative weight
	if newWeight < 0 {
		return ErrNegativeWeight
	}

	ingredient.Weight = newWeight
	m.IngridientMap[name] = ingredient

	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) updateNutritionalValueAndWeight() error {
	m.NutritionalValue = NutritionalValueAbsolute{
		Proteins:      0,
		Fats:          0,
		Carbohydrates: 0,
		Calories:      0,
	}

	for _, ingredient := range m.IngridientMap {
		nv, err := m.NutritionalValue.AddRelativeValue(ingredient.FoodProduct.NutritionalValueRelative, ingredient.Weight/100)
		if err != nil {
			return err
		}
		m.NutritionalValue = nv
	}
	return nil
}

func (m *Meal) CalculateTotalPrice() (float32, error) {
	var totalPrice float32
	for _, ingredient := range m.IngridientMap {
		if ingredient.FoodProduct.PricePerPkg < 0 {
			return 0, ErrNegativePrice
		}
		totalPrice += ingredient.FoodProduct.PricePerPkg * float32(ingredient.Weight) / float32(ingredient.FoodProduct.WeightPerPkg)
	}
	m.Price = totalPrice
	return totalPrice, nil
}
