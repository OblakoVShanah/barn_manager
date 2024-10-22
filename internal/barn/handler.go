package barn

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	router  *chi.Mux
	service Service
}

func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (handler *Handler) Register() {
	handler.router.Group(func(r chi.Router) {
		r.Get("/api/v1/products", handler.getProducts)
		r.Post("/api/v1/products", handler.postProducts)
	})
}

func (handler *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.service.AvailableProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", products)
}

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
