package services

import (
	"fmt"
	"shareen/src/models"
	"shareen/src/repositories"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type LobbiesService struct {
	lobbiesRepository *repositories.LobbiesRepository
	usersRepository   *repositories.UsersRepository
}

func NewLobbiesService(usersRepository *repositories.UsersRepository,
	lobbiesRepository *repositories.LobbiesRepository) *LobbiesService {
	return &LobbiesService{
		lobbiesRepository: lobbiesRepository,
		usersRepository:   usersRepository,
	}
}

func (ls *LobbiesService) CreateLobby() (*models.Lobby, error) {
	lobbyId := uuid.New().ID()
	lobby := &models.Lobby{
		LobbyURL:  createUniqueLobbyURL(lobbyId),
		CreatedAt: time.Now().GoString(),
	}
	return ls.lobbiesRepository.CreateLobby(lobby)
}

func (ls *LobbiesService) GetLobby(lobbyId string) (*models.Lobby, error) {
	if lobbyId == "" {
		return nil, fmt.Errorf("lobby id can't be null")
	}
	return ls.GetLobby(lobbyId)
}

func (ls *LobbiesService) GetAllLobbies() ([]*models.Lobby, error) {
	return ls.lobbiesRepository.GetAllLobbies()
}

func createUniqueLobbyURL(id uint32) string {

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var alpabetLen = uint32(utf8.RuneCountInString(alphabet))

	var (
		nums    []uint32
		num     = id
		builder strings.Builder
	)
	for num > 0 {
		nums = append(nums, num&alpabetLen)
		num /= alpabetLen

	}
	slices.Reverse(nums)
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}
