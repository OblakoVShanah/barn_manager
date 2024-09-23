package models

//Do tests after FoodProduct structs are created

import (
	"testing"
	"time"
)

var testProduct = FoodProduct{
	Name:            "Milk",
	WeightPerPkg:    1000,
	Amount:          1,
	PricePerPkg:     98.50,
	ExpirationDate:  time.Now().AddDate(0, 0, 5),
	PresentInFridge: true,
}

func TestChechExpirationDate(t *testing.T) {
	tmpBool := testProduct.CheckExpirationDate()
	if tmpBool != true {
		t.Errorf("expected expiration true, got %t", tmpBool)
	}
}

func TestUpdateProductWeight(t *testing.T) {
	testProduct.UpdateProductWeight(930)

	if testProduct.WeightPerPkg != 930 {
		t.Errorf("expected weight 930, got %d", testProduct.WeightPerPkg)
	}
}

func TestUpdateProductAmount(t *testing.T) {
	testProduct.UpdateProductAmount(2)

	if testProduct.Amount != 2 {
		t.Errorf("expected amount 2, got %d", testProduct.Amount)
	}
	if testProduct.PresentInFridge != true {
		t.Errorf("expected true PresentInFridge, got %t", testProduct.PresentInFridge)
	}

	testProduct.UpdateProductAmount(0)

	if testProduct.Amount != 0 {
		t.Errorf("expected amount 0, got %d", testProduct.Amount)
	}
	if testProduct.PresentInFridge != false {
		t.Errorf("expected false PresentInFridge, got %t", testProduct.PresentInFridge)
	}
}

func TestUpdateProductPrice(t *testing.T) {
	testProduct.UpdateProductPrice(100.21)

	if testProduct.PricePerPkg != 100.21 {
		t.Errorf("expected price 100.21, got %f", testProduct.PricePerPkg)
	}
}

func TestUpdateProductName(t *testing.T) {
	testProduct.UpdateProductName("CowMilk")

	if testProduct.Name != "CowMilk" {
		t.Errorf("expected name 'CowMilk', got %s", testProduct.Name)
	}
}
