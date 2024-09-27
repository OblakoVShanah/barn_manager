package models

// Nutritional value on 100 gramms of product
type NutritionalValueRelative struct {
	Proteins      float32 `json:"proteins"`
	Fats          float32 `json:"fats"`
	Carbohydrates float32 `json:"carbohydrates"`
	Calories      float32 `json:"calories"`
}

// Nutritional value of whole product
type NutritionalValueAbsolute struct {
	Proteins      float32 `json:"proteins"`
	Fats          float32 `json:"fats"`
	Carbohydrates float32 `json:"carbohydrates"`
	Calories      float32 `json:"calories"`
}

func (nv_left NutritionalValueAbsolute) AddAbsoluteValue(nv_right NutritionalValueAbsolute) NutritionalValueAbsolute {
	return NutritionalValueAbsolute{
		Proteins:      nv_left.Proteins + nv_right.Proteins,
		Fats:          nv_left.Fats + nv_right.Fats,
		Carbohydrates: nv_left.Carbohydrates + nv_right.Carbohydrates,
		Calories:      nv_left.Calories + nv_right.Calories,
	}
}

func (nv_left NutritionalValueAbsolute) AddRelativeValue(nv_right NutritionalValueRelative, weight_right float32) NutritionalValueAbsolute {
	return NutritionalValueAbsolute{
		Proteins:      nv_left.Proteins + nv_right.Proteins * weight_right,
		Fats:          nv_left.Fats + nv_right.Fats * weight_right,
		Carbohydrates: nv_left.Carbohydrates + nv_right.Carbohydrates * weight_right,
		Calories:      nv_left.Calories + nv_right.Calories * weight_right,
	}
}
