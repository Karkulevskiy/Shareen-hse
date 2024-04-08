package services

import (
	"fmt"
	"shareen/src/models"
	"shareen/src/repositories"
)
type UsersService struct {
	userRepository *repositories.UsersRepository
	lobbiesRepository *repositories.LobbiesRepository
}

func NewUsersService(userRepository *repositories.UsersRepository,
	 lobbiesRepository *repositories.LobbiesRepository) *UsersService{
	return &UsersService{
		userRepository: userRepository,
		lobbiesRepository: lobbiesRepository,
	}
}

func (us *UsersService) CreateUser(user *models.User) (*models.User, error){
	if user.Name == ""{
		return nil, fmt.Errorf("name can't be empty")
	}
	return us.userRepository.CreateUser(user)
}

func (us *UsersService) GetUser(userId string) (*models.User, error){
	if userId == ""{
		return nil, fmt.Errorf("user id can't be null")
	}
	return us.userRepository.GetUser(userId)
}