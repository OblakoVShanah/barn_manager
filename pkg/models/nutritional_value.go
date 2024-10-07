package models

import (
	"errors"
)

var (
	// ErrWeightMustBeGreaterThanZero -- Вес должен быть больше нуля
	ErrWeightMustBeGreaterThanZero = errors.New("weight must be greater than zero")

	// ErrIncorrectNutritionalValue -- Пищевая ценность указана некорректно
	ErrIncorrectNutritionalValue = errors.New("nutritional value is not correct")
)

// A NutritionalValueRelative represents a nutritional value per 100g of a product
type NutritionalValueRelative struct {
	Proteins      int `json:"proteins"`
	Fats          int `json:"fats"`
	Carbohydrates int `json:"carbohydrates"`
	Calories      int `json:"calories"`
}

// NutritionalValueAbsolute represents a nutritional value of a whole product
type NutritionalValueAbsolute struct {
	Proteins      int `json:"proteins"`
	Fats          int `json:"fats"`
	Carbohydrates int `json:"carbohydrates"`
	Calories      int `json:"calories"`
}

// AddAbsoluteValue adds absolute nutritional value to existed NutritionalValueAbsolute and
// returns new instance of NutritionalValueAbsolute.
func (nvLeft NutritionalValueAbsolute) AddAbsoluteValue(nvRight NutritionalValueAbsolute) NutritionalValueAbsolute {
	return NutritionalValueAbsolute{
		Proteins:      nvLeft.Proteins + nvRight.Proteins,
		Fats:          nvLeft.Fats + nvRight.Fats,
		Carbohydrates: nvLeft.Carbohydrates + nvRight.Carbohydrates,
		Calories:      nvLeft.Calories + nvRight.Calories,
	}
}

// AddRelativeValue adds relative nutritional value multiplied by weight to existed NutritionalValueAbsolute and
// returns new instance of NutritionalValueAbsolute and error in case of wrong weight.
func (nvLeft NutritionalValueAbsolute) AddRelativeValue(nvRight NutritionalValueRelative, weightRight int) (NutritionalValueAbsolute, error) {
	if weightRight <= 0 {
		return NutritionalValueAbsolute{}, ErrWeightMustBeGreaterThanZero
	}
	return NutritionalValueAbsolute{
		Proteins:      nvLeft.Proteins + int(nvRight.Proteins*weightRight/100),
		Fats:          nvLeft.Fats + int(nvRight.Fats*weightRight/100),
		Carbohydrates: nvLeft.Carbohydrates + int(nvRight.Carbohydrates*weightRight/100),
		Calories:      nvLeft.Calories + int(nvRight.Calories*weightRight/100),
	}, nil
}

// ValidateNutritionalValueRelative -- Проверяет, корректны ли значения пищевой ценности.
// Если одно из значений белков, жиров, углеводов или калорий отрицательное, возвращает ошибку ErrIncorrectNutritionalValue.
func ValidateNutritionalValueRelative(nv NutritionalValueRelative) error {
	if nv.Proteins < 0 || nv.Fats < 0 || nv.Carbohydrates < 0 || nv.Calories < 0 {
		return ErrIncorrectNutritionalValue
	}
	return nil
}
