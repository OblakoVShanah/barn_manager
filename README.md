# Barn Manager Service

A microservice responsible for managing product storage and creating shopping lists as part of the Havchik Podbirator ecosystem.

## Service Responsibilities

- Tracking products in user storage (refrigerator, pantry, etc.)
- Monitoring product expiration dates
- Calculating required ingredient quantities
- Creating shopping lists based on meal plans and current inventory

## Architecture

This service is part of the Havchik Podbirator system:
- Exchanges data with the menu management service to receive meal plans
- Provides inventory data to the menu management service
- Creates shopping lists based on meal plans and current stock levels

## API Endpoints

### GET /api/v1/products
Retrieves all available products in the storage.

**Response:**
```json
[
  {
    "id": "product-id",
    "name": "Product Name",
    "weight_per_pkg": 100,
    "amount": 5,
    "price_per_pkg": 10.50,
    "expiration_date": "2024-12-31T23:59:59Z",
    "present_in_fridge": true,
    "nutritional_value_relative": {
      "proteins": 10,
      "fats": 5,
      "carbohydrates": 15,
      "calories": 150
    }
  }
]
```

### POST /api/v1/products
Adds a new product to the storage.

**Request Body:**
```json
{
  "id": "product-id",
  "name": "Product Name",
  "weight_per_pkg": 100,
  "amount": 5,
  "price_per_pkg": 10.50,
  "expiration_date": "2024-12-31T23:59:59Z",
  "present_in_fridge": true,
  "nutritional_value_relative": {
    "proteins": 10,
    "fats": 5,
    "carbohydrates": 15,
    "calories": 150
  }
}
```

### POST /api/v1/products/check-availability
Checks product availability for recipes and returns a shopping list.

**Request Body:**
```json
[
  {
    "steps": ["step1", "step2"],
    "ingredients": [
      {"product_id": "milk", "amount": 500},
      {"product_id": "flour", "amount": 250}
    ]
  }
]
```

**Response:**
```json
{
  "products": [
    {
      "id": "missing-product",
      "amount": 100
    }
  ]
}
```

## Configuration

The service uses a YAML configuration file (`configs/config.yaml`):

```yaml
host: "0.0.0.0"
port: "8082"
db:
  dsn: "user:password@tcp(localhost:3306)/database?parseTime=true&loc=Local"
```

## Development

### Prerequisites
- Go 1.21+
- MySQL 8.0+

### Running the Service
```bash
go run main.go
```

### Testing
```bash
go test ./...
```

## Curl Examples

```bash 
# Get all products
curl http://localhost:8082/api/v1/products

# Add a new product
curl -X POST http://localhost:8082/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "id": "apple",
    "name": "Apple", 
    "weight_per_pkg": 100, 
    "amount": 5,
    "price_per_pkg": 2.50,
    "expiration_date": "2024-12-31T23:59:59Z",
    "present_in_fridge": true,
    "nutritional_value_relative": {
      "proteins": 1,
      "fats": 0,
      "carbohydrates": 14,
      "calories": 52
    }
  }'

# Check product availability for recipes
curl -X POST http://localhost:8082/api/v1/products/check-availability \
  -H "Content-Type: application/json" \
  -d '[
    {
      "steps": ["Mix ingredients"],
      "ingredients": [
        {"product_id": "milk", "amount": 500},
        {"product_id": "flour", "amount": 250}
      ]
    }
  ]'
```