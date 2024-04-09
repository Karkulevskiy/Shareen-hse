package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
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

func (ur *UsersRepository) GetUser(userId string) (*models.User, *models.ResponseError) {
	query := "SELECT * FROM users WHERE id = $1"
	rows, err := ur.dbHandler.Query(query, userId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	//TODO Проверить, что будет если lobbyid == nil
	var id, lobbyid, name string
	for rows.Next() {
		err = rows.Scan(&id, &lobbyid, &name)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.User{
		ID:      id,
		LobbyID: lobbyid,
		Name:    name,
	}, nil
}

func (ur *UsersRepository) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := "INSERT INTO users (id, lobby_id, name) VALUES ($1, $2, $3) RETURNING id"
	rows, err := ur.dbHandler.Query(query, user.LobbyID, user.Name)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var userId string
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.User{
		ID:      userId,
		LobbyID: user.LobbyID,
		Name:    user.Name,
	}, nil
}

func (ur *UsersRepository) GetAllUsers() ([]*models.User, *models.ResponseError) {
	query := "SELECT * FROM users"
	rows, err := ur.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	var id, lobbyId, name string
	for rows.Next() {
		err = rows.Scan(&id, &lobbyId, &name)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		users = append(users, &models.User{
			ID:      id,
			LobbyID: lobbyId,
			Name:    name,
		})
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return users, nil
}

func (ur *UsersRepository) DeleteUser(userId string) *models.ResponseError {
	query := "DELETE FROM users WHERE id = $1"
	res, err := ur.dbHandler.Exec(query, userId)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: fmt.Sprintf("user with id: {%s} not found ", userId),
			Status:  http.StatusNotFound,
		}
	}
	return nil
}
func (ur *UsersRepository) UpdateUser(user *models.User) *models.ResponseError {
	query := `UPDATE users
	SET
	name = $1,
	lobby_id = $2
	WHERE id = $3`
	res, err := ur.dbHandler.Exec(query, user.LobbyID, user.Name, user.ID)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: fmt.Sprintf("user with id: {%s} for updating not found", user.ID),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
