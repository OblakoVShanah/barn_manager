package postgres

import (
	"context"
	// "database/sql"
	"fmt"

	"github.com/OblakoVShanah/havchik_podbirator/internal/menu"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) LoadMenu(ctx context.Context) (menu.Menu, error) {
	return menu.Menu{}, fmt.Errorf("not implemented")
}

func (s *Storage) SaveMenu(ctx context.Context, m menu.Menu) (id string, err error) {
	return "", fmt.Errorf("not implemented")
}
