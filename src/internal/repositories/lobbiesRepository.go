package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"shareen/src/internal/models"
)

type LobbiesRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

// Constructor for creating repository layer
func NewLobbiesRepository(dbHandler *sql.DB) *LobbiesRepository {
	return &LobbiesRepository{
		dbHandler: dbHandler,
	}
}

func (lr *LobbiesRepository) GetLobby(lobbyID string) (*models.Lobby, *models.ResponseError) {
	query := `SELECT lobbies.id, lobby_url, video_url, created_at, changed_at, users.id, name
				FROM lobbies
				LEFT JOIN lobbies_users ON lobbies.id = lobbies_users.lobby_id
				LEFT JOIN users ON users.id = lobbies_users.user_id
				WHERE lobbies.id = $1`
	rows, err := lr.dbHandler.Query(query, lobbyID)
	if err != nil {
		log.Println("error occured while getting lobby by id: ", err.Error())
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var lobbyId, lobbyUrl, createdAt string
	var videoURL, changedAt, userId, userName sql.NullString
	var lobbyUsers []*models.User
	for rows.Next() {
		err = rows.Scan(&lobbyId, &lobbyUrl, &videoURL, &createdAt, &changedAt, &userId, &userName)
		if err != nil {
			log.Println("error occured while getting (scanning) lobby by id: ", err.Error())
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		if userId.String != "" {
			lobbyUsers = append(lobbyUsers, &models.User{
				ID:   userId.String,
				Name: userName.String,
			})
		}
	}
	if rows.Err() != nil {
		log.Println("any error occured while getting lobby by id: ", err.Error())
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if len(lobbyUsers) == 0 {
		lobbyUsers = nil
	}
	return &models.Lobby{
		ID:        lobbyId,
		LobbyURL:  lobbyUrl,
		VideoURL:  videoURL.String,
		CreatedAt: createdAt,
		ChangedAt: changedAt.String,
		UserList:  lobbyUsers,
	}, nil
}

func (lr *LobbiesRepository) GetLobbyUsers(lobbyId string) ([]*models.User, *models.ResponseError) {
	query := `SELECT FROM lobbies_users, users 
			  WHERE users.id = lobbies_users.user_id
			  AND lobbies_users.lobby_id = $1`
	rows, err := lr.dbHandler.Query(query, lobbyId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var userId, userName string
	var usersList []*models.User
	for rows.Next() {
		err := rows.Scan(&userId, &userName)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		usersList = append(usersList, &models.User{
			ID:   userId,
			Name: userName,
		})
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return usersList, nil
}

func (lr *LobbiesRepository) CreateLobby(lobby *models.Lobby) (*models.Lobby, *models.ResponseError) {
	queryFirst := "INSERT INTO lobbies (lobby_url, video_url, created_at) VALUES ($1, $2, $3) RETURNING id"
	rows, err := lr.dbHandler.Query(queryFirst, lobby.LobbyURL, lobby.VideoURL, lobby.CreatedAt)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var lobbyId string
	for rows.Next() {
		err := rows.Scan(&lobbyId)
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
	return &models.Lobby{
		ID:        lobbyId,
		LobbyURL:  lobby.LobbyURL,
		VideoURL:  lobby.VideoURL,
		CreatedAt: lobby.CreatedAt,
		UserList:  nil,
	}, nil
}

func (lr *LobbiesRepository) DeleteLobby(lobbyID string) *models.ResponseError {
	query := "DELETE FROM lobbies WHERE id = $1"
	res, err := lr.dbHandler.Exec(query, lobbyID)
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
			Message: fmt.Sprintf("Not found lobby with id: {%s}", lobbyID),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (lr *LobbiesRepository) GetAllLobbies() ([]*models.Lobby, *models.ResponseError) {
	query := `SELECT lobbies.id as lobby_id, lobbies.lobby_url,
			 lobbies.video_url, lobbies.created_at, lobbies.changed_at,
    		users.id as user_id, users.name
			FROM lobbies
			LEFT JOIN lobbies_users ON lobbies.id = lobbies_users.lobby_id
			LEFT JOIN users ON lobbies_users.user_id = users.id;` //тут написать сложный запрос
	rows, err := lr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var lobbyId, lobbyURL, videoURL, createdAt string
	var changedAt, userID, userName sql.NullString
	lobbiesUsers := map[string]*models.Lobby{}
	for rows.Next() {
		err = rows.Scan(&lobbyId, &lobbyURL, &videoURL, &createdAt, &changedAt, &userID, &userName)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		lobby := &models.Lobby{
			ID:        lobbyId,
			LobbyURL:  lobbyURL,
			VideoURL:  videoURL,
			CreatedAt: createdAt,
			ChangedAt: changedAt.String,
		}
		if _, ok := lobbiesUsers[lobby.ID]; !ok {
			lobbiesUsers[lobby.ID] = lobby
		}
		if userID.String == "" {
			if len(lobbiesUsers[lobby.ID].UserList) == 0 {
				lobbiesUsers[lobby.ID].UserList = nil
			}
			continue
		}
		user := &models.User{
			ID:   userID.String,
			Name: userName.String,
		}
		lobbiesUsers[lobby.ID].UserList = append(lobbiesUsers[lobby.ID].UserList, user)
	}

	// Посмотреть что за null в сваггере, лоигку добавления в массив
	// Странно отображается время

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	lobbies := make([]*models.Lobby, 0)
	for _, lobby := range lobbiesUsers {
		lobbies = append(lobbies, lobby)
	}
	return lobbies, nil
}

func (lr *LobbiesRepository) DeleteAllLobbies() *models.ResponseError {
	query := "TRUNCATE TABLE lobbies"
	row := lr.dbHandler.QueryRow(query)
	if row.Err() != nil {
		return &models.ResponseError{
			Message: row.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (lr *LobbiesRepository) UpdateLobby(lobby *models.Lobby) *models.ResponseError {
	query := `
		UPDATE lobbies
		SET
		lobby_url = $1,
		video_url = $2,
		changed_at = $3
		WHERE id = $4
	`
	res, err := lr.dbHandler.Exec(query, lobby.LobbyURL, lobby.VideoURL, lobby.ChangedAt, lobby.ID)
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
			Message: fmt.Sprintf("lobby for updating with id: {%s} not found", lobby.ID),
			Status:  http.StatusNotFound,
		}
	}
	return nil
}
