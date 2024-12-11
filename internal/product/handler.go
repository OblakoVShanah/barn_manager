package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Handler обрабатывает HTTP-запросы для работы с продуктами
type Handler struct {
	router  *chi.Mux
	service Service
}

// NewHandler создает новый обработчик HTTP-запросов
func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

// Register регистрирует все обработчики маршрутов
func (handler *Handler) Register() {
	handler.router.Group(func(r chi.Router) {
		r.Get("/api/v1/products", handler.getProducts)
		r.Post("/api/v1/products", handler.postProducts)
		r.Post("/api/v1/products/check-availability", handler.checkProductsAvailability)
	})
}

// getProducts возвращает список всех продуктов
func (handler *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.service.AvailableProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// postProducts добавляет новый продукт
func (handler *Handler) postProducts(w http.ResponseWriter, r *http.Request) {
	var product FoodProduct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.service.PlaceProduct(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product placed with ID: %s", id)
}

// checkProductsAvailability проверяет наличие необходимых продуктов
func (handler *Handler) checkProductsAvailability(w http.ResponseWriter, r *http.Request) {
	var recipes []struct {
		Steps       []string                 `json:"steps"`
		Ingredients []map[string]interface{} `json:"ingredients"`
	}

	// Декодируем массив рецептов из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&recipes); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Print(time.Now().String() + " -> ")
	log.Println(recipes)

	// requirements будет содержать суммарные требования по продуктам
	requirements := make(map[string]uint)

	// Пробегаемся по всем рецептам и их ингредиентам
	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			productID, ok := ingredient["product_id"].(string)
			if !ok {
				http.Error(w, "Invalid product_id in ingredients", http.StatusBadRequest)
				return
			}

			amount, ok := ingredient["amount"].(float64)
			if !ok {
				http.Error(w, "Invalid amount in ingredients", http.StatusBadRequest)
				return
			}

			// Добавляем количество к общей карте требований
			requirements[productID] += uint(amount)
		}
	}

	// Проверяем доступность продуктов
	shoppingList, err := handler.service.CheckAvailability(r.Context(), requirements)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Кодируем результат в JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(shoppingList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
