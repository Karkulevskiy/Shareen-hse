package services

import (
	"net/http"
	"shareen/src/models"
	"shareen/src/repositories"
	"shareen/src/utils"
)

type UsersService struct {
	usersRepository   *repositories.UsersRepository
	lobbiesRepository *repositories.LobbiesRepository
}

func NewUsersService(userRepository *repositories.UsersRepository,
	lobbiesRepository *repositories.LobbiesRepository) *UsersService {
	return &UsersService{
		usersRepository:   userRepository,
		lobbiesRepository: lobbiesRepository,
	}
}

func (us *UsersService) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	if user.Name == "" {
		return nil, &models.ResponseError{
			Message: "user name can't be empty",
			Status:  http.StatusBadRequest,
		}
	}
	return us.usersRepository.CreateUser(user)
}

func (us *UsersService) GetUser(userId string) (*models.User, *models.ResponseError) {
	err := utils.ValidateId(userId)
	if err != nil {
		return nil, err
	}
	return us.usersRepository.GetUser(userId)
}

func (us *UsersService) GetAllUsers() ([]*models.User, *models.ResponseError) {
	return us.usersRepository.GetAllUsers()
}

func (us *UsersService) DeleteUser(userId string) *models.ResponseError {
	err := utils.ValidateId(userId)
	if err != nil {
		return err
	}
	return us.usersRepository.DeleteUser(userId)
}

func (us *UsersService) UpdateUser(user *models.User) *models.ResponseError {
	err := utils.ValidateId(user.ID)
	if err != nil {
		return err
	}
	return us.usersRepository.UpdateUser(user)
}
