package menu

import (
	"context"
	"time"
	"github.com/OblakoVShanah/havchik_podbirator/internal/common"
	"github.com/OblakoVShanah/havchik_podbirator/internal/barn"
)

type Menu struct {
	ID    string
	Meals []Meal
}

type Meal struct {
	ID               string
	EatingTime       time.Time
	IngredientMap    map[string]Ingredient
	NutritionalValue common.NutritionalValueAbsolute
	Price            float32
}

type Ingredient struct {
	FoodProduct barn.FoodProduct
	Weight      int
}

type Service interface {
	Menu(ctx context.Context) (Menu, error)
	Place(ctx context.Context, menu Menu) (id string, err error)
}

type Store interface {
	LoadMenu(ctx context.Context) (Menu, error)
	SaveMenu(ctx context.Context, menu Menu) (id string, err error)
}
