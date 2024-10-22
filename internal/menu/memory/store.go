package memory

import (
	"context"
	"sync"
	"fmt"

	"github.com/OblakoVShanah/havchik_podbirator/internal/menu"
)

type Storage struct {
	menus map[string]menu.Menu
	mu    sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		menus: make(map[string]menu.Menu),
	}
}

func (s *Storage) LoadMenu(ctx context.Context) (menu.Menu, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, m := range s.menus {
		return m, nil
	}

	return menu.Menu{}, fmt.Errorf("no menu found")
}

func (s *Storage) SaveMenu(ctx context.Context, m menu.Menu) (id string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.menus[m.ID] = m
	return m.ID, nil
}
