package models

import (
	"fmt"
	"slices"
	"time"
)

type FoodProduct struct {
	Name                     string
	WeightPerPkg             int // in gramms
	Amount                   uint
	PricePerPkg              float32
	ExpirationDate           time.Time
	PresentInFridge          bool
	NutritionalValueRelative NutritionalValueRelative
}

// compare an expiration date of a product with current date,
// return true if expiration date < current date otherwise false
func (fp *FoodProduct) IsSpoiled() bool {
	if fp.ExpirationDate.Unix() < time.Now().Unix() {
		fmt.Println(fp.Name, "was rotten. The expiration date --", fp.ExpirationDate)
		return true
	} else {
		fmt.Println(fp.Name, "is fine. The expiration date --", fp.ExpirationDate)
		return false
	}
}

// compare an expiration date of a product with current date,
// return true if expiration date is in a day from current date
// otherwise false
func (fp *FoodProduct) WillSpoilSoon() bool {
	if fp.ExpirationDate.Unix() < time.Now().AddDate(0, 0, 1).Unix() {
		fmt.Println(fp.Name, "will spoil soon. The expiration date --", fp.ExpirationDate)
		return true
	}
	fmt.Println(fp.Name, "is fine. The expiration date --", fp.ExpirationDate)
	return false
}

// change FoodProduct's weight per package
func (fp *FoodProduct) UpdateProductWeight(newWeight int) {
	fp.WeightPerPkg = newWeight
	fmt.Println("New weight --", fp.WeightPerPkg)
}

// change FoodProduct's amount and automaticaly update PresentInFridge state
func (fp *FoodProduct) UpdateProductAmount(newAmount uint) {
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

// slice examle
func SliceExample() {
	productSlice := []FoodProduct{
		{
			Name:            "Milk",
			WeightPerPkg:    1000,
			Amount:          1,
			PricePerPkg:     98.50,
			ExpirationDate:  time.Now().AddDate(0, 0, 5),
			PresentInFridge: true,
		},
		{
			Name:            "Egg",
			WeightPerPkg:    100,
			Amount:          5,
			PricePerPkg:     9.99,
			ExpirationDate:  time.Now().AddDate(0, 0, -5),
			PresentInFridge: true,
		},
		{
			Name:            "Tomato",
			WeightPerPkg:    100,
			Amount:          0,
			PricePerPkg:     5.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 10),
			PresentInFridge: false,
		},
		{
			Name:            "Chicken",
			WeightPerPkg:    500,
			Amount:          1,
			PricePerPkg:     19.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 1),
			PresentInFridge: true,
		},
		{
			Name:            "Cheese",
			WeightPerPkg:    200,
			Amount:          1,
			PricePerPkg:     49.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 1),
			PresentInFridge: true,
		},
	}

	fmt.Println(productSlice)

	index_to_delete := make([]int, 0, len(productSlice))
	for i, v := range productSlice {
		// goes through the slice and check each product
		// if it is out or already rotten
		if !v.PresentInFridge || v.IsSpoiled() {
			index_to_delete = append(index_to_delete, i)
			fmt.Printf("Product %v is rotten or out\n", v.Name)
			continue
		}
		if v.WillSpoilSoon() {
			fmt.Printf("Product %v should be eaten today, otherwise it will rot\n", v.Name)
		}
	}

	for _, v := range index_to_delete {
		// delete "bad" products
		productSlice = slices.Delete(productSlice, v, v+1)
	}

	fmt.Println(productSlice)

}
