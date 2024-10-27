package common

// NutritionalValueRelative представляет относительную пищевую ценность продукта (на 100г)
type NutritionalValueRelative struct {
	Proteins      int `json:"proteins"`      // Белки в граммах
	Fats          int `json:"fats"`          // Жиры в граммах
	Carbohydrates int `json:"carbohydrates"` // Углеводы в граммах
	Calories      int `json:"calories"`      // Калории
}

// NutritionalValueAbsolute представляет абсолютную пищевую ценность продукта
type NutritionalValueAbsolute struct {
	Proteins      int `json:"proteins"`      // Белки в граммах
	Fats          int `json:"fats"`          // Жиры в граммах
	Carbohydrates int `json:"carbohydrates"` // Углеводы в граммах
	Calories      int `json:"calories"`      // Калории
}
