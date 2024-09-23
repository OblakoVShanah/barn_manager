package models

// Do tests after Foodstuff structs are created

// import (
//     "testing"
//     "time"
// )

// func TestAddIngredient(t *testing.T) {
//     meal := Meal{
//         Id:           "1",
//         EatingTime:   time.Now(),
//         Ingridients:    []Ingridient{},
//         Price:        0,
//         Weight:       0,
//     }

//     ingredient := Ingridient{
//         FoodStaff: FoodStaff{
//             Name: "Chicken",
//             PricePerKg: 10,
//             NutritionalValue: NutritionalValueAbsolute{
//                 Proteins: 25,
//                 Fats: 5,
//                 Carbohydrates: 0,
//                 Calories: 165,
//             },
//         },
//         Weight: 200,
//     }

//     meal.AddIngredient(ingredient)

//     if len(meal.Foodstuff) != 1 {
//         t.Errorf("expected 1 ingredient, got %d", len(meal.Foodstuff))
//     }
// }

// func TestRemoveIngredient(t *testing.T) {
//     meal := Meal{
//         Id:           "1",
//         EatingTime:   time.Now(),
//         Foodstuff:    []Ingridient{{FoodStaff: FoodStaff{Name: "Rice", PricePerKg: 2}, Weight: 100}},
//         Price:        0,
//         Weight:       0,
//     }

//     err := meal.RemoveIngredient(0)
//     if err != nil {
//         t.Errorf("unexpected error: %v", err)
//     }

//     if len(meal.Foodstuff) != 0 {
//         t.Errorf("expected 0 ingredients, got %d", len(meal.Foodstuff))
//     }
// }

// func TestCalculateTotalPrice(t *testing.T) {
//     meal := Meal{
//         Id:           "1",
//         EatingTime:   time.Now(),
//         Ingridients: []Ingridient{
//             {FoodStaff: FoodStaff{Name: "Chicken", PricePerKg: 10}, Weight: 200},
//             {FoodStaff: FoodStaff{Name: "Rice", PricePerKg: 2}, Weight: 100},
//         },
//     }

//     price := meal.CalculateTotalPrice()

//     if price != 2.2 { // (10 * 0.2) + (2 * 0.1) = 2.2
//         t.Errorf("expected price 2.2, got %f", price)
//     }
// }
