package menu

import "context"

type AppService struct {
	storage Store
}

func NewAppService(storage Store) *AppService {
	return &AppService{storage: storage}
}

func (s *AppService) Menu(ctx context.Context) (Menu, error) {
	menu, err := s.storage.LoadMenu(ctx)
	if err != nil {
		return Menu{}, err
	}

	return menu, nil
}

func (s *AppService) Place(ctx context.Context, menu Menu) (id string, err error) {
	id, err = s.storage.SaveMenu(ctx, menu)
	if err != nil {
		return "", err
	}

	return id, nil
}
