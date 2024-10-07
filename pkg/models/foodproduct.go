package models

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"time"
)

// ErrExpDateEmpty -- Дата истечения не указана
var ErrExpDateEmpty = errors.New("ExpirationDate is not specified")

// ErrNegativeAmount -- Отрицательное количество
var ErrNegativeAmount = errors.New("negative amount")

// ErrNegativeWeightPerPkg -- Отрицательный вес на упаковку
var ErrNegativeWeightPerPkg = errors.New("negative weight per package")

// ErrNegativePricePerPkg -- Отрицательная цена
var ErrNegativePricePerPkg = errors.New("negative price")

// FoodProduct представляет структуру данных для продуктов питания.
// Один экземпляр этой структуры соответствует одному типу продукта,
// но в одном экземпляре может быть несколько штук (упаковок) продукта.
type FoodProduct struct {
	Name                     string
	WeightPerPkg             uint                     // вес в граммах на упаковку
	Amount                   uint                     // количество упаковок
	PricePerPkg              float32                  // цена за упаковку
	ExpirationDate           time.Time                // срок годности
	PresentInFridge          bool                     // признак наличия в холодильнике
	NutritionalValueRelative NutritionalValueRelative // относительное содержание питательных веществ
}

// ValidateFoodproduct проверяет значения атрибутов продукта на корректность
func (fp *FoodProduct) ValidateFoodproduct() error {
	if fp.WeightPerPkg <= 0 {
		return ErrNegativeWeightPerPkg
	}
	if fp.Amount <= 0 {
		return ErrNegativeAmount
	}
	if fp.PricePerPkg <= 0 {
		return ErrNegativePricePerPkg
	}
	if error := ValidateNutritionalValueRelative(fp.NutritionalValueRelative); error != nil {
		return error
	}

	return nil
}

// IsSpoiled проверяет, испортился ли продукт (срок годности истек)
func (fp *FoodProduct) IsSpoiled() (bool, error) {
	if fp.ExpirationDate.IsZero() {
		return false, fmt.Errorf("IsSpoiled() error: %w", ErrExpDateEmpty)
	}
	if fp.ExpirationDate.Unix() < time.Now().Unix() {
		fmt.Println(fp.Name, "испортился. Срок годности --", fp.ExpirationDate)
		return true, nil
	}
	fmt.Println(fp.Name, "в порядке. Срок годности --", fp.ExpirationDate)
	return false, nil
}

// WillSpoilSoon проверяет, испортится ли продукт в течение следующего дня
func (fp *FoodProduct) WillSpoilSoon() (bool, error) {
	if fp.ExpirationDate.IsZero() {
		return false, fmt.Errorf("WillSpoilSoon() error: %w", ErrExpDateEmpty)
	}
	if fp.ExpirationDate.Unix() < time.Now().AddDate(0, 0, 1).Unix() {
		fmt.Println(fp.Name, "скоро испортится. Срок годности --", fp.ExpirationDate)
		return true, nil
	}
	fmt.Println(fp.Name, "в порядке. Срок годности --", fp.ExpirationDate)
	return false, nil
}

// UpdateProductWeight изменяет вес на упаковку продукта
func (fp *FoodProduct) UpdateProductWeight(newWeight uint) {
	fp.WeightPerPkg = newWeight
	fmt.Println("Новый вес --", fp.WeightPerPkg)
}

// UpdateProductAmount изменяет количество продукта и автоматически обновляет статус PresentInFridge
func (fp *FoodProduct) UpdateProductAmount(newAmount uint) {
	fp.Amount = newAmount
	if newAmount == 0 {
		fp.PresentInFridge = false
	} else {
		fp.PresentInFridge = true
	}
	fmt.Println("Новое количество --", fp.Amount, " Наличие в холодильнике --", fp.PresentInFridge)
}

// UpdateProductPrice изменяет цену за упаковку продукта
func (fp *FoodProduct) UpdateProductPrice(newPrice float32) {
	fp.PricePerPkg = newPrice
	fmt.Println("Новая цена --", fp.PricePerPkg)
}

// UpdateProductName изменяет название продукта
func (fp *FoodProduct) UpdateProductName(newName string) {
	fp.Name = newName
	fmt.Println("Новое название --", fp.Name)
}

// MapExample демонстрирует пример использования map
func MapExample() {
	productMap := make(map[string]*FoodProduct)

	productMap["Milk"] = &FoodProduct{
		Name:            "Молоко",
		WeightPerPkg:    1000,
		Amount:          1,
		PricePerPkg:     98.50,
		ExpirationDate:  time.Now().AddDate(0, 0, 5),
		PresentInFridge: true,
	}

	productMap["Beer"] = &FoodProduct{
		Name:            "Пиво",
		WeightPerPkg:    500,
		Amount:          1,
		PricePerPkg:     float32(999),
		ExpirationDate:  time.Now().AddDate(0, 1, 0),
		PresentInFridge: true,
	}

	for key, value := range productMap {
		fmt.Println("Ключ:", key, "Значение:", value)
	}

	delete(productMap, "Milk")
	fmt.Println("Молоко удалено.")

	fmt.Println("Пиво закончилось")
	productMap["Beer"].UpdateProductAmount(0)

	for key, value := range productMap {
		fmt.Println("Ключ:", key, "Значение:", value)
	}
}

// SliceExample демонстрирует пример использования среза
func SliceExample() {
	productSlice := []FoodProduct{
		{
			Name:            "Молоко",
			WeightPerPkg:    1000,
			Amount:          1,
			PricePerPkg:     98.50,
			ExpirationDate:  time.Now().AddDate(0, 0, 5),
			PresentInFridge: true,
		},
		{
			Name:            "Яйцо",
			WeightPerPkg:    100,
			Amount:          5,
			PricePerPkg:     9.99,
			ExpirationDate:  time.Now().AddDate(0, 0, -5),
			PresentInFridge: true,
		},
		{
			Name:            "Помидор",
			WeightPerPkg:    100,
			Amount:          0,
			PricePerPkg:     5.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 10),
			PresentInFridge: false,
		},
		{
			Name:            "Курица",
			WeightPerPkg:    500,
			Amount:          1,
			PricePerPkg:     19.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 1),
			PresentInFridge: true,
		},
		{
			Name:            "Сыр",
			WeightPerPkg:    200,
			Amount:          1,
			PricePerPkg:     49.99,
			ExpirationDate:  time.Now().AddDate(0, 0, 1),
			PresentInFridge: true,
		},
	}

	fmt.Println(productSlice)

	indexToDelete := make([]int, 0, len(productSlice))
	for i, v := range productSlice {
		// Проходит по срезу и проверяет каждый продукт
		// если он закончился или испортился
		flag, err := v.IsSpoiled()
		if err != nil {
			if errors.Is(err, ErrExpDateEmpty) {
				log.Fatalln(err)
			}
			log.Fatalln("неожиданная ошибка")
		}
		if !v.PresentInFridge || flag {
			indexToDelete = append(indexToDelete, i)
			fmt.Printf("Продукт %v испортился или закончился\n", v.Name)
			continue
		}
		flag, err = v.WillSpoilSoon()
		if err != nil {
			if errors.Is(err, ErrExpDateEmpty) {
				log.Fatalln(err)
			}
			log.Fatalln("неожиданная ошибка")
		}
		if flag {
			fmt.Printf("Продукт %v нужно съесть сегодня, иначе он испортится\n", v.Name)
		}
	}

	for _, v := range indexToDelete {
		// Удаляет "плохие" продукты
		productSlice = slices.Delete(productSlice, v, v+1)
	}

	fmt.Println(productSlice)

}
