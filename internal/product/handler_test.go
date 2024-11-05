package product_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	common "github.com/OblakoVShanah/barn_manager/internal/models"
	"github.com/OblakoVShanah/barn_manager/internal/product"
	"github.com/OblakoVShanah/barn_manager/internal/product/mock"
	"github.com/go-chi/chi/v5"
	"github.com/google/go-cmp/cmp"
)

func TestHandler_getProducts(t *testing.T) {
	mockStore := mock.NewStore()
	router := chi.NewRouter()
	service := product.NewService(mockStore)
	handler := product.NewHandler(router, service)
	handler.Register()

	testProducts := []product.FoodProduct{
		{
			ID:              "test-id-1",
			Name:            "Test Product 1",
			WeightPerPkg:    100,
			Amount:          1,
			PricePerPkg:     9.99,
			ExpirationDate:  time.Now().Add(24 * time.Hour),
			PresentInFridge: true,
			NutritionalValueRelative: common.NutritionalValueRelative{
				Proteins:      10,
				Fats:          20,
				Carbohydrates: 30,
				Calories:      400,
			},
		},
	}

	t.Run("successful get", func(t *testing.T) {
		mockStore.SetProducts(testProducts)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var got []product.FoodProduct
		err := json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if diff := cmp.Diff(testProducts, got); diff != "" {
			t.Errorf("handler returned wrong body (-want +got):\n%s", diff)
		}
	})

	t.Run("error case", func(t *testing.T) {
		mockStore.SetError(errors.New("database error"))

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusInternalServerError)
		}
	})
}
