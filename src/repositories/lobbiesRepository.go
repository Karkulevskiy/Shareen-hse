package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"shareen/src/models"
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
	query := "SELECT * FROM lobbies WHERE id = $1"
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
	var videoURL, changedAt, userLobbyId, userId, userName sql.NullString
	userList := make([]*models.User, 0)
	for rows.Next() {
		err = rows.Scan(&lobbyId, &lobbyUrl, &videoURL, &createdAt, &changedAt, &userId, &userLobbyId, &userName)
		if err != nil {
			log.Println("error occured while getting (scanning) lobby by id: ", err.Error())
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		if userId.String != "" {
			user := &models.User{
				ID:      userId.String,
				LobbyID: userLobbyId.String,
				Name:    userName.String,
			}
			userList = append(userList, user)
		}
	}
	if rows.Err() != nil {
		log.Println("any error occured while getting lobby by id: ", err.Error())
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if lobbyId == "" {
		return nil, &models.ResponseError{
			Message: fmt.Sprintf("lobby not found with id : {%s}", lobbyID),
			Status:  http.StatusNotFound,
		}
	}
	if len(userList) == 0 {
		userList = nil
	}
	return &models.Lobby{
		ID:        lobbyId,
		LobbyURL:  lobbyUrl,
		VideoURL:  videoURL.String,
		CreatedAt: createdAt,
		ChangedAt: changedAt.String,
		UserList:  userList,
	}, nil
}

func (lr *LobbiesRepository) CreateLobby(lobby *models.Lobby) (*models.Lobby, *models.ResponseError) {
	query := "INSERT INTO lobbies (lobby_url, video_url, created_at) VALUES ($1, $2, $3) RETURNING id"
	rows, err := lr.dbHandler.Query(query, lobby.LobbyURL, lobby.VideoURL, lobby.CreatedAt)
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
	query := "SELECT * FROM lobbies" //тут написать сложный запрос
	rows, err := lr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var lobbyId, lobbyURL, videoURL, createdAt string
	var changedAt, userID, userName, userLobbyID sql.NullString
	lobbies_users := map[*models.Lobby][]*models.User{}
	for rows.Next() {
		err = rows.Scan(&lobbyId, &lobbyURL, &videoURL, &createdAt, &changedAt, &userID, &userName, &userLobbyID)
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
			ChangedAt: createdAt,
		}
		user := &models.User{
			ID:      userID.String,
			LobbyID: userLobbyID.String,
			Name:    userName.String,
		}
		if lobbies_users[lobby] == nil {
			lobbies_users[lobby] = nil
		} else {
			lobbies_users[lobby] = append(lobbies_users[lobby], user)
		}
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
	for lKey, uVal := range lobbies_users {
		lKey.UserList = uVal
		lobbies = append(lobbies, lKey)
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
		created_at = $3,
		changed_at = $4
		WHERE id = $5
	`
	res, err := lr.dbHandler.Exec(query, lobby.LobbyURL, lobby.VideoURL, lobby.CreatedAt, lobby.ChangedAt, lobby.ID)
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

func (lr *LobbiesRepository) GetLobbyUsers(lobbyId string) ([]*models.User, *models.ResponseError) {
	query := "SELECT * FROM users WHERE lobby_id = $1"
	rows, err := lr.dbHandler.Query(query, lobbyId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	var userId, lobbyUserId, name string
	for rows.Next() {
		err = rows.Scan(&userId, &lobbyUserId, &name)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		user := &models.User{
			ID:      userId,
			LobbyID: lobbyUserId,
			Name:    name,
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: rows.Err().Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return users, nil
}
