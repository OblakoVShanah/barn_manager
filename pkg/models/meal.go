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
	Price            float64
	Weight           float64
}

type Ingridient struct {
	FoodStaff FoodStaff
	Weight    float64
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

func (m *Meal) UpdateIngredientWeight(index int, newWeight float64) error {
	if index < 0 || index >= len(m.Ingridients) {
		return fmt.Errorf("invalid index")
	}
	m.Ingridients[index].Weight = newWeight
	m.updateNutritionalValueAndWeight()
	return nil
}

func (m *Meal) updateNutritionalValueAndWeight() {
	var totalProteins, totalFats, totalCarbs, totalCalories, totalWeight float64

	for _, ingredient := range m.Ingridients {
		totalWeight += ingredient.Weight
		totalProteins += ingredient.FoodStaff.NutritionalValueRelative.Proteins * (ingredient.Weight / 100)
		totalFats += ingredient.FoodStaff.NutritionalValueRelative.Fats * (ingredient.Weight / 100)
		totalCarbs += ingredient.FoodStaff.NutritionalValueRelative.Carbohydrates * (ingredient.Weight / 100)
		totalCalories += ingredient.FoodStaff.NutritionalValueRelative.Calories * (ingredient.Weight / 100)
	}

	m.Weight = totalWeight
	m.NutritionalValue = NutritionalValueAbsolute{
		Proteins:      totalProteins,
		Fats:          totalFats,
		Carbohydrates: totalCarbs,
		Calories:      totalCalories,
	}
}

func (m *Meal) CalculateTotalPrice() float64 {
	var totalPrice float64
	for _, ingredient := range m.Ingridients {
		totalPrice += ingredient.FoodStaff.PricePerKg * (ingredient.Weight / 1000)
	}
	m.Price = totalPrice
	return totalPrice
}
