package barn

import "context"

type AppService struct {
	storage Store
}

func NewService(storage Store) Service {
    return &AppService{storage: storage}
}

func (s *AppService) AvailableProducts(ctx context.Context) ([]FoodProduct, error) {
	products, err := s.storage.LoadProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *AppService) PlaceProduct(ctx context.Context, product FoodProduct) (id string, err error) {
	id, err = s.storage.SaveProduct(ctx, product)
	if err != nil {
		return "", err
	}

	return id, nil
}
