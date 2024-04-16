package services

import (
	"fmt"
	"shareen/src/internal/models"
	"shareen/src/internal/repositories"
	"shareen/src/internal/utils"
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
	fmt.Println(lobbyId)
	lobby := &models.Lobby{
		LobbyURL:  createUniqueLobbyURL(lobbyId),
		CreatedAt: time.Now().GoString(),
	}
	err := repositories.BeginTransaction(ls.lobbiesRepository, ls.usersRepository)
	if err != nil {
		return nil, err
	}
	response, responseErr := ls.lobbiesRepository.CreateLobby(lobby)
	if responseErr != nil {
		err := repositories.RollbackTransaction(ls.lobbiesRepository, ls.usersRepository)
		if err != nil {
			return nil, err
		}
		return nil, responseErr
	}
	err = repositories.CommitTransaction(ls.lobbiesRepository, ls.usersRepository)
	if err != nil {
		return nil, err
	}
	return response, nil
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
	transactionErr := repositories.BeginTransaction(ls.lobbiesRepository, ls.usersRepository)
	if transactionErr != nil {
		return transactionErr
	}
	responseErr := ls.lobbiesRepository.DeleteLobby(lobbyId)
	if responseErr != nil {
		err = repositories.RollbackTransaction(ls.lobbiesRepository, ls.usersRepository)
		if err != nil {
			return err
		}
		return responseErr
	}
	err = repositories.CommitTransaction(ls.lobbiesRepository, ls.usersRepository)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LobbiesService) DeleteAllLobbies() *models.ResponseError {
	transactionErr := repositories.BeginTransaction(ls.lobbiesRepository, ls.usersRepository)
	if transactionErr != nil {
		return transactionErr
	}
	responseErr := ls.lobbiesRepository.DeleteAllLobbies()
	if responseErr != nil {
		transactionErr := repositories.RollbackTransaction(ls.lobbiesRepository, ls.usersRepository)
		if transactionErr != nil {
			return transactionErr
		}
		return responseErr
	}
	transactionErr = repositories.CommitTransaction(ls.lobbiesRepository, ls.usersRepository)
	if transactionErr != nil {
		return transactionErr
	}
	return nil
}

func (ls *LobbiesService) UpdateLobby(lobby *models.Lobby) *models.ResponseError {
	err := utils.ValidateId(lobby.ID)
	if err != nil {
		return err
	}
	lobby.ChangedAt = time.Now().String()
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
		nums = append(nums, num%alpabetLen)
		num /= alpabetLen
	}
	slices.Reverse(nums)
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}
