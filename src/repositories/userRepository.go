package repositories

import (
	"database/sql"
	"fmt"
	"shareen/src/models"
)

type UserRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
	}
}

func (ur *UserRepository) GetUser(userId string) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	rows, err := ur.dbHandler.Query(query, userId)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Error occured during getting user: %w", err.Error())
	}
	//TODO Проверить, что будет если lobbyid == nil
	var id, lobbyid, name string
	for rows.Next() {
		err = rows.Scan(&id, &lobbyid, &name)
		if err != nil {
			return nil, fmt.Errorf("Error occured during getting user: %w", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("Error occured during getting user: %w", err.Error())
	}
	return &models.User{
		ID:      id,
		LobbyID: lobbyid,
		Name:    name,
	}, nil
}

func (ur *UserRepository) CreateUser() {

}

func (ur *UserRepository) GetAllUsers() {

}

func (ur *UserRepository) DeleteUser() {

}
