package menu

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler handles HTTP requests for the menu service.
type Handler struct {
	router  *chi.Mux
	service Service
}

// NewHandler creates a new Handler with the given router and service.
func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router: router,
		service: service,
	}
}

func (handler *Handler) Register() {
	handler.router.Group(func(r chi.Router) {
		r.Get("/api/v1/menu", handler.getMenu)
		r.Post("/api/v1/menu", handler.postMenu)
	})
}

func (handler *Handler) getMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := handler.service.Menu(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", menu)
}

func (handler *Handler) postMenu(w http.ResponseWriter, r *http.Request) {
	var menu Menu
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.service.Place(r.Context(), menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Menu created with ID: %s", id)
}
