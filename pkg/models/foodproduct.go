package models

import (
	"fmt"
	"time"
)

type FoodProduct struct {
	Name                     string
	WeightPerPkg             float32 // in gramms
	Amount                   int
	PricePerPkg              float32
	ExpirationDate           time.Time
	PresentInFridge          bool
	NutritionalValueRelative NutritionalValueRelative
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
func (fp *FoodProduct) UpdateProductWeight(newWeight float32) {
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
