package repositories

import (
	"shareen/src/internal/models"
	"shareen/src/internal/repositories"
)

type InMemoryStorage struct {
	lobbies map[string]models.Lobby
}

func NewMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		lobbies: map[string]models.Lobby{},
	}
}

// TODO: посмотреть на каком слое создавать кастомные ошибки
// TODO: в какой момент обновлять данные в лобби
func (m *InMemoryStorage) Update(lobbyID string) error {
	if _, ok := m.lobbies[lobbyID]; !ok {
		return repositories.ErrUserNotFound
	}

}
