package product_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"bytes"
	"strings"

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

func TestHandler_checkProductsAvailability(t *testing.T) {
	mockStore := mock.NewStore()
	router := chi.NewRouter()
	service := product.NewService(mockStore)
	handler := product.NewHandler(router, service)
	handler.Register()

	// Setup test products in store
	testProducts := []product.FoodProduct{
		{
			ID:              "овсяные_хлопья",
			Name:            "Овсяные хлопья",
			WeightPerPkg:    1000,
			Amount:          1000,
			PresentInFridge: true,
		},
		{
			ID:              "молоко",
			Name:            "Молоко",
			WeightPerPkg:    1000,
			Amount:          1000,
			PresentInFridge: true,
		},
	}
	mockStore.SetProducts(testProducts)

	t.Run("successful check", func(t *testing.T) {
		recipe := map[string]interface{}{
			"steps": []string{
				"Вскипятить молоко",
				"Добавить хлопья",
				"Варить 5 минут",
			},
			"ingredients": []map[string]interface{}{
				{
					"amount":     100,
					"product_id": "овсяные_хлопья",
				},
				{
					"amount":     200,
					"product_id": "молоко",
				},
			},
		}

		body, err := json.Marshal(recipe)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/check-availability", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var got product.ShoppingList
		err = json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		// Since all products are available, shopping list should be empty
		if len(got.Products) != 0 {
			t.Errorf("Expected empty shopping list, got %d items", len(got.Products))
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/check-availability", strings.NewReader("invalid json"))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusBadRequest)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mockStore.SetError(errors.New("database error"))

		recipe := map[string]interface{}{
			"steps": []string{"Step 1"},
			"ingredients": []map[string]interface{}{
				{
					"unit":       "г",
					"amount":     100.0,
					"product_id": "овсяные_хлопья",
				},
			},
		}

		body, err := json.Marshal(recipe)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/check-availability", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusInternalServerError)
		}
	})
}