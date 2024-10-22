package barn

import (
	"context"
	"github.com/OblakoVShanah/havchik_podbirator/internal/common"
	"time"
)

type FoodProduct struct {
	ID                       string                          `json:"id"`
	Name                     string                          `json:"name"`
	WeightPerPkg             uint                            `json:"weight_per_pkg"`
	Amount                   uint                            `json:"amount"`
	PricePerPkg              float32                         `json:"price_per_pkg"`
	ExpirationDate           time.Time                       `json:"expiration_date"`
	PresentInFridge          bool                            `json:"present_in_fridge"`
	NutritionalValueRelative common.NutritionalValueRelative `json:"nutritional_value_relative"`
}

type Service interface {
	AvailableProducts(ctx context.Context) ([]FoodProduct, error)
	PlaceProduct(ctx context.Context, product FoodProduct) (id string, err error)
}

type Store interface {
	LoadProducts(ctx context.Context) ([]FoodProduct, error)
	SaveProduct(ctx context.Context, product FoodProduct) (id string, err error)
}
