package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"shareen/src/internal/models"
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
	var id, name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
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
		ID:   id,
		Name: name,
	}, nil
}

func (ur *UsersRepository) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	rows, err := ur.dbHandler.Query(query, user.Name)
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
		ID:   userId,
		Name: user.Name,
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
	var id, name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		users = append(users, &models.User{
			ID:   id,
			Name: name,
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
	name = $1
	WHERE id = $2`
	res, err := ur.dbHandler.Exec(query, user.Name, user.ID)
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

func (ur *UsersRepository) JoinUserInLobby(userID, lobbyID string) *models.ResponseError {
	query := `UPDATE lobbies_users SET user_id = $1 WHERE lobby_id = $2`
	rows, err := ur.dbHandler.Exec(query, userID, lobbyID)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: fmt.Sprintf("lobby with id: {%s} not found", lobbyID),
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
