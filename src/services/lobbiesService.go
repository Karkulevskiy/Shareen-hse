package services

import (
	"shareen/src/models"
	"shareen/src/repositories"
	"shareen/src/utils"
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

func (ls *LobbiesService) CreateLobby() (*models.Lobby, *models.ResponseError) {
	lobbyId := uuid.New().ID()
	lobby := &models.Lobby{
		LobbyURL:  createUniqueLobbyURL(lobbyId),
		CreatedAt: time.Now().GoString(),
	}
	return ls.lobbiesRepository.CreateLobby(lobby)
}

func (ls *LobbiesService) GetLobby(lobbyId string) (*models.Lobby, *models.ResponseError) {
	err := utils.ValidateId(lobbyId)
	if err != nil {
		return nil, err
	}
	return ls.lobbiesRepository.GetLobby(lobbyId)
}

func (ls *LobbiesService) GetAllLobbies() ([]*models.Lobby, *models.ResponseError) {
	return ls.lobbiesRepository.GetAllLobbies()
}

func (ls *LobbiesService) DeleteLobby(lobbyId string) *models.ResponseError {
	err := utils.ValidateId(lobbyId)
	if err != nil {
		return err
	}
	return ls.lobbiesRepository.DeleteLobby(lobbyId)
}

func (ls *LobbiesService) DeleteAllLobbies() *models.ResponseError {
	return ls.lobbiesRepository.DeleteAllLobbies()
}

func (ls *LobbiesService) UpdateLobby(lobby *models.Lobby) *models.ResponseError {
	err := utils.ValidateId(lobby.ID)
	if err != nil {
		return err
	}
	return ls.lobbiesRepository.UpdateLobby(lobby)
}

func (ls *LobbiesService) GetLobbyUsers(lobbyId string) ([]*models.User, *models.ResponseError) {
	err := utils.ValidateId(lobbyId)
	if err != nil {
		return nil, err
	}
	return ls.lobbiesRepository.GetLobbyUsers(lobbyId)
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
