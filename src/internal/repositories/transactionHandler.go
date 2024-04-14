package repositories

import (
	"context"
	"database/sql"
	"net/http"
	"shareen/src/internal/models"
)

func BeginTransaction(lobbiesRepository *LobbiesRepository,
	usersRepository *UsersRepository) *models.ResponseError {
	ctx := context.Background()
	transaction, err := lobbiesRepository.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	lobbiesRepository.transaction = transaction
	usersRepository.transaction = transaction
	return nil
}

func RollbackTransaction(lobbiesRepository *LobbiesRepository,
	usersRepository *UsersRepository) *models.ResponseError {
	transaction := lobbiesRepository.transaction
	lobbiesRepository.transaction = nil
	usersRepository.transaction = nil
	err := transaction.Rollback()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func CommitTransaction(lobbiesRepository *LobbiesRepository,
	usersRepository *UsersRepository) *models.ResponseError {
	transaction := lobbiesRepository.transaction
	lobbiesRepository.transaction = nil
	usersRepository.transaction = nil
	err := transaction.Commit()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
