package models

import (
	"errors"
)

var (
	ErrWeightMustBeGreaterThanZero = errors.New("weight must be greater than zero")
	ErrIncorrectNutritionalValue   = errors.New("Nutritional value is not correct")
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
func (nv_left NutritionalValueAbsolute) AddAbsoluteValue(nv_right NutritionalValueAbsolute) NutritionalValueAbsolute {
	return NutritionalValueAbsolute{
		Proteins:      nv_left.Proteins + nv_right.Proteins,
		Fats:          nv_left.Fats + nv_right.Fats,
		Carbohydrates: nv_left.Carbohydrates + nv_right.Carbohydrates,
		Calories:      nv_left.Calories + nv_right.Calories,
	}
}

// AddRelativeValue adds relative nutritional value multiplied by weight to existed NutritionalValueAbsolute and
// returns new instance of NutritionalValueAbsolute and error in case of wrong weight.
func (nv_left NutritionalValueAbsolute) AddRelativeValue(nv_right NutritionalValueRelative, weight_right int) (NutritionalValueAbsolute, error) {
	if weight_right <= 0 {
		return NutritionalValueAbsolute{}, ErrWeightMustBeGreaterThanZero
	}
	return NutritionalValueAbsolute{
		Proteins:      nv_left.Proteins + int(nv_right.Proteins*weight_right/100),
		Fats:          nv_left.Fats + int(nv_right.Fats*weight_right/100),
		Carbohydrates: nv_left.Carbohydrates + int(nv_right.Carbohydrates*weight_right/100),
		Calories:      nv_left.Calories + int(nv_right.Calories*weight_right/100),
	}, nil
}

func ValidateNutritionalValueRelative(nv NutritionalValueRelative) error {
	if nv.Proteins < 0 || nv.Fats < 0 || nv.Carbohydrates < 0 || nv.Calories < 0 {
		return ErrIncorrectNutritionalValue
	}
	return nil
}
