package models

import (
	"testing"
	"time"
)

func TestFoodProduct_IsSpoiled(t *testing.T) {
	tests := []struct {
		name string
		fp   *FoodProduct
		want bool
	}{
		{
			"norm Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			false,
		},
		{
			"spoiled Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, -1),
				PresentInFridge: true,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fp.IsSpoiled(); got != tt.want {
				t.Errorf("FoodProduct.CheckExpirationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodProduct_UpdateProductWeight(t *testing.T) {

	tests := []struct {
		name      string
		fp        *FoodProduct
		newWeight int
	}{
		{
			"Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fp.UpdateProductWeight(tt.newWeight)
			if tt.fp.WeightPerPkg != tt.newWeight {
				t.Errorf("FoodProduct.WeightPerPkg = %v, want %v", tt.fp.WeightPerPkg, tt.newWeight)
			}
		})
	}
}

func TestFoodProduct_UpdateProductAmount(t *testing.T) {
	tests := []struct {
		name          string
		fp            *FoodProduct
		newAmount     uint
		inFridgeState bool
	}{
		{
			"Milk1to2",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			2,
			true,
		},
		{
			"Milk1to0",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fp.UpdateProductAmount(tt.newAmount)
			if tt.fp.Amount != tt.newAmount {
				t.Errorf("FoodProduct.Amount = %v, want %v", tt.fp.Amount, tt.newAmount)
			}
			if tt.fp.PresentInFridge != tt.inFridgeState {
				t.Errorf("FoodProduct.PresentInFridge = %v, want %v", tt.fp.PresentInFridge, tt.inFridgeState)
			}
		})
	}
}

func TestFoodProduct_UpdateProductPrice(t *testing.T) {
	tests := []struct {
		name     string
		fp       *FoodProduct
		newPrice float32
	}{
		{
			"Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			111.11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fp.UpdateProductPrice(tt.newPrice)
			if tt.fp.PricePerPkg != tt.newPrice {
				t.Errorf("FoodProduct.PricePerPkg = %v, want %v", tt.fp.PricePerPkg, tt.newPrice)
			}
		})
	}
}

func TestFoodProduct_UpdateProductName(t *testing.T) {
	tests := []struct {
		name    string
		fp      *FoodProduct
		newName string
	}{
		{
			"Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			"Milk2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fp.UpdateProductName(tt.newName)
			if tt.fp.Name != tt.newName {
				t.Errorf("FoodProduct.Name = %v, want %v", tt.fp.Name, tt.newName)
			}
		})
	}
}

func TestMapExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Map test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapExample()
		})
	}
}

func TestFoodProduct_WillSpoilSoon(t *testing.T) {
	tests := []struct {
		name string
		fp   *FoodProduct
		want bool
	}{
		{
			"norm Milk",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 5),
				PresentInFridge: true,
			},
			false,
		},
		{
			"Milk almost spoiled",
			&FoodProduct{
				Name:            "Milk",
				WeightPerPkg:    1000,
				Amount:          1,
				PricePerPkg:     98.50,
				ExpirationDate:  time.Now().AddDate(0, 0, 1).Add(time.Hour * -1),
				PresentInFridge: true,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fp.WillSpoilSoon(); got != tt.want {
				t.Errorf("FoodProduct.IsSpoilSoon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Slice tst"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceExample()
		})
	}
}
