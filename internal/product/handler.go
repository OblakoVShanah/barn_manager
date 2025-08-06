package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Handler handles HTTP requests for product operations.
// It provides REST API endpoints for managing products and checking availability.
type Handler struct {
	router  *chi.Mux
	service Service
}

// NewHandler creates a new HTTP request handler for product operations.
// It initializes the handler with the provided router and service.
//
// Parameters:
//   - router: The Chi router to register routes with
//   - service: The product service to use for business logic
//
// Returns:
//   - *Handler: The initialized handler
func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

// Register registers all route handlers with the router.
// It sets up the REST API endpoints for product management operations.
func (handler *Handler) Register() {
	handler.router.Group(func(r chi.Router) {
		r.Get("/api/v1/products", handler.getProducts)
		r.Post("/api/v1/products", handler.postProducts)
		r.Post("/api/v1/products/check-availability", handler.checkProductsAvailability)
	})
}

// getProducts handles GET requests to retrieve all available products.
// It returns a JSON response with the list of all products in the storage.
//
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func (handler *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.service.AvailableProducts(r.Context())
	if err != nil {
		log.Printf("Error retrieving products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding products response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// postProducts handles POST requests to add a new product.
// It accepts a JSON payload with product information and returns the ID of the created product.
//
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func (handler *Handler) postProducts(w http.ResponseWriter, r *http.Request) {
	var product FoodProduct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Printf("Error decoding product request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := handler.service.PlaceProduct(r.Context(), product)
	if err != nil {
		log.Printf("Error placing product: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product placed with ID: %s", id)
}

// checkProductsAvailability handles POST requests to check product availability for recipes.
// It accepts a JSON payload with recipe information and returns a shopping list of missing products.
//
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func (handler *Handler) checkProductsAvailability(w http.ResponseWriter, r *http.Request) {
	var recipes []struct {
		Steps       []string                 `json:"steps"`
		Ingredients []map[string]interface{} `json:"ingredients"`
	}

	// Декодируем массив рецептов из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&recipes); err != nil {
		log.Printf("Error decoding recipes request: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%s -> Processing %d recipes", time.Now().String(), len(recipes))

	// requirements будет содержать суммарные требования по продуктам
	requirements := make(map[string]uint)

	// Пробегаемся по всем рецептам и их ингредиентам
	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			productID, ok := ingredient["product_id"].(string)
			if !ok {
				log.Printf("Invalid product_id in ingredients: %v", ingredient)
				http.Error(w, "Invalid product_id in ingredients", http.StatusBadRequest)
				return
			}

			amount, ok := ingredient["amount"].(float64)
			if !ok {
				log.Printf("Invalid amount in ingredients: %v", ingredient)
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
		log.Printf("Error checking product availability: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Кодируем результат в JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(shoppingList); err != nil {
		log.Printf("Error encoding shopping list response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
