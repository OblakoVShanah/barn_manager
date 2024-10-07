package models

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"time"
)

var ErrExpDateEmpty = errors.New("ExpirationDate is not specified")

// A FoodProduct is a data structure for groceries.
// One instance of the structure corresponds to one type of food product,
// but there can be several peaces (packages) of a product in one structure.
type FoodProduct struct {
	Name                     string
	WeightPerPkg             uint // in gramms
	Amount                   uint
	PricePerPkg              float32
	ExpirationDate           time.Time
	PresentInFridge          bool
	NutritionalValueRelative NutritionalValueRelative
}

// IsSpoiled reports whether an expiration date of a product < current date.
func (fp *FoodProduct) IsSpoiled() (bool, error) {
	if fp.ExpirationDate.IsZero() {
		return false, fmt.Errorf("IsSpoiled() error: %w", ErrExpDateEmpty)
	}
	if fp.ExpirationDate.Unix() < time.Now().Unix() {
		fmt.Println(fp.Name, "was rotten. The expiration date --", fp.ExpirationDate)
		return true, nil
	}
	fmt.Println(fp.Name, "is fine. The expiration date --", fp.ExpirationDate)
	return false, nil
}

// WillSpoilSoon reports whether and expiration date of a product is in a day from current date
func (fp *FoodProduct) WillSpoilSoon() (bool, error) {
	if fp.ExpirationDate.IsZero() {
		return false, fmt.Errorf("WillSpoilSoon() error: %w", ErrExpDateEmpty)
	}
	if fp.ExpirationDate.Unix() < time.Now().AddDate(0, 0, 1).Unix() {
		fmt.Println(fp.Name, "will spoil soon. The expiration date --", fp.ExpirationDate)
		return true, nil
	}
	fmt.Println(fp.Name, "is fine. The expiration date --", fp.ExpirationDate)
	return false, nil
}

// UpdateProductWeight changes weight per package of a product
func (fp *FoodProduct) UpdateProductWeight(newWeight uint) {
	fp.WeightPerPkg = newWeight
	fmt.Println("New weight --", fp.WeightPerPkg)
}

// UpdateProductAmount changes an amount of a product and automaticaly updates PresentInFridge state
func (fp *FoodProduct) UpdateProductAmount(newAmount uint) {
	fp.Amount = newAmount
	if newAmount == 0 {
		fp.PresentInFridge = false
	} else {
		fp.PresentInFridge = true
	}
	fmt.Println("New amount --", fp.Amount, " Present in fridge --", fp.PresentInFridge)
}

// UpdateProductPrice changes a price per package of a product
func (fp *FoodProduct) UpdateProductPrice(newPrice float32) {
	fp.PricePerPkg = newPrice
	fmt.Println("New price --", fp.PricePerPkg)
}

// UpdateProductName changes a product name
func (fp *FoodProduct) UpdateProductName(newName string) {
	fp.Name = newName
	fmt.Println("New name --", fp.Name)
}

// MapExample shows an the example of usage map
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

// SliceExample shows en example of usage slice
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
		flag, err := v.IsSpoiled()
		if err != nil {
			if errors.Is(err, ErrExpDateEmpty) {
				log.Fatalln(err)
			}
			log.Fatalln("unexpected error")
		}
		if !v.PresentInFridge || flag {
			index_to_delete = append(index_to_delete, i)
			fmt.Printf("Product %v is rotten or out\n", v.Name)
			continue
		}
		flag, err = v.WillSpoilSoon()
		if err != nil {
			if errors.Is(err, ErrExpDateEmpty) {
				log.Fatalln(err)
			}
			log.Fatalln("unexpected error")
		}
		if flag {
			fmt.Printf("Product %v should be eaten today, otherwise it will rot\n", v.Name)
		}
	}

	for _, v := range index_to_delete {
		// delete "bad" products
		productSlice = slices.Delete(productSlice, v, v+1)
	}

	fmt.Println(productSlice)

}
