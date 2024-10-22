package common

type NutritionalValueRelative struct {
	Proteins      int `json:"proteins"`
	Fats          int `json:"fats"`
	Carbohydrates int `json:"carbohydrates"`
	Calories      int `json:"calories"`
}

type NutritionalValueAbsolute struct {
	Proteins      int `json:"proteins"`
	Fats          int `json:"fats"`
	Carbohydrates int `json:"carbohydrates"`
	Calories      int `json:"calories"`
}
