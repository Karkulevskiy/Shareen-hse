package repositories

import (
	"database/sql"
	"fmt"
	"shareen/src/models"
)

type UsersRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (ur *UsersRepository) GetUser(userId string) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	rows, err := ur.dbHandler.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error occured during getting user: %s", err.Error())
	}
	defer rows.Close()
	//TODO Проверить, что будет если lobbyid == nil
	var id, lobbyid, name string
	for rows.Next() {
		err = rows.Scan(&id, &lobbyid, &name)
		if err != nil {
			return nil, fmt.Errorf("error occured during getting user: %s", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error occured during getting user: %s", err.Error())
	}
	return &models.User{
		ID:      id,
		LobbyID: lobbyid,
		Name:    name,
	}, nil
}

func (ur *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	query := "INSERT INTO users (id, lobby_id, name) VALUES ($1, $2, $3) RETURNING id"
	rows, err := ur.dbHandler.Query(query, user.LobbyID, user.Name)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating user: %s", err.Error())
	}
	defer rows.Close()
	var userId string
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return nil, fmt.Errorf("error occured while creating user: %s", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("any errors occured while creating user: %s", err.Error())
	}
	return &models.User{
		ID:      userId,
		LobbyID: user.LobbyID,
		Name:    user.Name,
	}, nil
}

func (ur *UsersRepository) GetAllUsers() ([]*models.User, error) {
	query := "SELECT * FROM users"
	rows, err := ur.dbHandler.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting all users: %s", err.Error())
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	var id, lobbyId, name string
	for rows.Next() {
		err = rows.Scan(&id, &lobbyId, &name)
		if err != nil {
			return nil, fmt.Errorf("error occured while getting all users: %s", err.Error())
		}
		users = append(users, &models.User{
			ID:      id,
			LobbyID: lobbyId,
			Name:    name,
		})
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("any error occured while getting all users: %s", err.Error())
	}
	return users, nil
}

func (ur *UsersRepository) DeleteUser() {

}
