package models

import (
	"fmt"
	"time"
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

func (m *Meal) AddIngredient(ingredient Ingridient) {
	// Initialize the map if it's nil
	if m.IngridientMap == nil {
		m.IngridientMap = make(map[string]Ingridient)
	}

	// Add or update the ingredient in the map
	m.IngridientMap[ingredient.FoodProduct.Name] = ingredient

	m.updateNutritionalValueAndWeight()
}

func (m *Meal) RemoveIngredient(name string) error {
	// Check if ingredient exists
	if _, exists := m.IngridientMap[name]; !exists {
		return fmt.Errorf("ingredient not found")
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
		return fmt.Errorf("ingredient not found")
	}

	// Update weight in map
	ingredient.Weight = newWeight
	m.IngridientMap[name] = ingredient

	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) updateNutritionalValueAndWeight() {
	m.NutritionalValue = NutritionalValueAbsolute{
		Proteins:      0,
		Fats:          0,
		Carbohydrates: 0,
		Calories:      0,
	}

	for _, ingredient := range m.IngridientMap {
		m.NutritionalValue = m.NutritionalValue.AddRelativeValue(ingredient.FoodProduct.NutritionalValueRelative, float32(ingredient.Weight)/100)
	}
}

func (m *Meal) CalculateTotalPrice() float32 {
	var totalPrice float32
	for _, ingredient := range m.IngridientMap {
		totalPrice += ingredient.FoodProduct.PricePerPkg * float32(ingredient.Weight) / float32(ingredient.FoodProduct.WeightPerPkg)
	}
	m.Price = totalPrice
	return totalPrice
}
