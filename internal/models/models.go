package common

// NutritionalValueRelative represents the relative nutritional value of a product per 100g.
// This structure is used to store nutritional information that is standardized per 100 grams
// of the product, making it easy to calculate nutritional content for different amounts.
type NutritionalValueRelative struct {
	Proteins      int `json:"proteins"`      // Proteins in grams per 100g
	Fats          int `json:"fats"`          // Fats in grams per 100g
	Carbohydrates int `json:"carbohydrates"` // Carbohydrates in grams per 100g
	Calories      int `json:"calories"`      // Calories per 100g
}

// NutritionalValueAbsolute represents the absolute nutritional value of a product.
// This structure is used to store the total nutritional information for a specific
// amount of product, rather than per 100g.
type NutritionalValueAbsolute struct {
	Proteins      int `json:"proteins"`      // Total proteins in grams
	Fats          int `json:"fats"`          // Total fats in grams
	Carbohydrates int `json:"carbohydrates"` // Total carbohydrates in grams
	Calories      int `json:"calories"`      // Total calories
}
