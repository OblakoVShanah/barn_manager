package models

type NutritionalValue interface {
	NutritionalValueRelative | NutritionalValueAbsolute
}


// Nutritional value on 100 gramms of product
type NutritionalValueRelative struct {
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
	Calories      float64 `json:"calories"`
}


// Nutritional value of whole product
type NutritionalValueAbsolute struct {
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
	Calories      float64 `json:"calories"`
}


func (nv NutritionalValueAbsolute) Add(nv2 NutritionalValueAbsolute) NutritionalValueAbsolute {
	return NutritionalValueAbsolute{
		Proteins:      nv.Proteins + nv2.Proteins,
		Fats:          nv.Fats + nv2.Fats,
		Carbohydrates: nv.Carbohydrates + nv2.Carbohydrates,
		Calories:      nv.Calories + nv2.Calories,
	}
}
