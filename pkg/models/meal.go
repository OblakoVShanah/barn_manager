package models

import (
	"fmt"
	"time"
)

type Meal struct {
	Id               string
	EatingTime       time.Time
	Ingridients      []Ingridient
	NutritionalValue NutritionalValueAbsolute
	Price            float32
}

type Ingridient struct {
	FoodProduct FoodProduct
	Weight      float32
}

func (m *Meal) AddIngredient(ingredient Ingridient) {
	m.Ingridients = append(m.Ingridients, ingredient)
	m.updateNutritionalValueAndWeight()
}

func (m *Meal) RemoveIngredient(index int) error {
	if index < 0 || index >= len(m.Ingridients) {
		return fmt.Errorf("invalid index")
	}
	m.Ingridients = append(m.Ingridients[:index], m.Ingridients[index+1:]...)
	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) UpdateIngredientWeight(index int, newWeight float32) error {
	if index < 0 || index >= len(m.Ingridients) {
		return fmt.Errorf("invalid index")
	}
	m.Ingridients[index].Weight = newWeight
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

	for _, ingredient := range m.Ingridients {
		m.NutritionalValue = m.NutritionalValue.AddRelativeValue(ingredient.FoodProduct.NutritionalValueRelative, ingredient.Weight / 100)
	}
}

func (m *Meal) CalculateTotalPrice() float32 {
	var totalPrice float32
	for _, ingredient := range m.Ingridients {
		totalPrice += ingredient.FoodProduct.PricePerPkg * (ingredient.Weight / ingredient.FoodProduct.WeightPerPkg)
	}
	m.Price = totalPrice
	return totalPrice
}
