package services

import (
	"net/http"
	"shareen/src/internal/models"
	"shareen/src/internal/repositories"
	"shareen/src/internal/utils"
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

func (us *UsersService) CreateUser(userName string) (*models.User, *models.ResponseError) {
	if userName == "" {
		return nil, &models.ResponseError{
			Message: "user name can't be empty",
			Status:  http.StatusBadRequest,
		}
	} // ПР что я тут имел ввиду - хз
	user := &models.User{
		Name: userName,
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
	transactionErr := repositories.BeginTransaction(us.lobbiesRepository, us.usersRepository)
	if transactionErr != nil {
		return transactionErr
	}
	responserErr := us.usersRepository.DeleteUser(userId)
	if responserErr != nil {
		transactionErr = repositories.RollbackTransaction(us.lobbiesRepository, us.usersRepository)
		if transactionErr != nil {
			return transactionErr
		}
		return responserErr
	}
	transactionErr = repositories.CommitTransaction(us.lobbiesRepository, us.usersRepository)
	if transactionErr != nil {
		return transactionErr
	}
	return nil
}

func (us *UsersService) UpdateUser(user *models.User) *models.ResponseError {
	err := utils.ValidateId(user.ID)
	if err != nil {
		return err
	}
	return us.usersRepository.UpdateUser(user)
}
