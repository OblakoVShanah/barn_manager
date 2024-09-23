package models

import (
	"fmt"
	"time"
)

type FoodProduct struct {
	Name            string
	WeightPerPkg    int // in gramms
	Amount          int
	PricePerPkg     float32
	ExpirationDate  time.Time
	PresentInFridge bool
	Proteins        int // per 100g
	Fats            int // per 100g
	Carbohydrates   int // per 100g
	Calories        int // per 100g
}

// compare an expiration date of a product with current date,
// return false if expiration date < current date otherwise true
func (fp *FoodProduct) CheckExpirationDate() bool {
	if fp.ExpirationDate.Unix() < time.Now().Unix() {
		fmt.Println(fp.Name, "was rotten. The expiration date --", fp.ExpirationDate)
		return false
	} else {
		fmt.Println(fp.Name, "is fine. The expiration date --", fp.ExpirationDate)
		return true
	}
}

// change FoodProduct's weight per package
func (fp *FoodProduct) UpdateProductWeight(newWeight int) {
	fp.WeightPerPkg = newWeight
	fmt.Println("New weight --", fp.WeightPerPkg)
}

// change FoodProduct's amount and automaticaly update PresentInFridge state
func (fp *FoodProduct) UpdateProductAmount(newAmount int) {
	fp.Amount = newAmount
	if newAmount == 0 {
		fp.PresentInFridge = false
	} else {
		fp.PresentInFridge = true
	}
	fmt.Println("New amount --", fp.Amount, " Present in fridge --", fp.PresentInFridge)
}

// change FoodProduct's price per package
func (fp *FoodProduct) UpdateProductPrice(newPrice float32) {
	fp.PricePerPkg = newPrice
	fmt.Println("New price --", fp.PricePerPkg)
}

// change FoodProduct's name
func (fp *FoodProduct) UpdateProductName(newName string) {
	fp.Name = newName
	fmt.Println("New name --", fp.Name)
}

// map example
func MapExample() {
	productMap := make(map[string]*FoodProduct)

	productMap["Milk"] = &FoodProduct{
		Name:            "Milk",
		WeightPerPkg:    1000,
		Amount:          1,
		PricePerPkg:     98.50,
		ExpirationDate:  time.Now().AddDate(0, 0, 5),
		PresentInFridge: true,
	}

	productMap["Beer"] = &FoodProduct{
		Name:            "Beer",
		WeightPerPkg:    500,
		Amount:          1,
		PricePerPkg:     float32(999),
		ExpirationDate:  time.Now().AddDate(0, 1, 0),
		PresentInFridge: true,
	}

	for key, value := range productMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	delete(productMap, "Milk")
	fmt.Println("Milk deleted.")

	fmt.Println("Beer was run out")
	productMap["Beer"].UpdateProductAmount(0)

	for key, value := range productMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
