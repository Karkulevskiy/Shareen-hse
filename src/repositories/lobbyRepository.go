package repositories

import (
	"database/sql"
	"fmt"
	"shareen/src/models"
)

type LobbyRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewLobbyRepository(dbHandler *sql.DB) *LobbyRepository {
	return &LobbyRepository{
		dbHandler: dbHandler,
	}
}

func (lr *LobbyRepository) GetLobby(lobbyID string) (*models.Lobby, error) {
	query := "SELECT * FROM lobbies WHERE id = $1"
	rows, err := lr.dbHandler.Query(query, lobbyID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Error occured selecting lobby by id: %w", err.Error())
	}
	var id, lobbyUrl, videoUrl, createdAt string
	userList := make([]*models.User, 0)
	for rows.Next() {
		err = rows.Scan(&id, &lobbyUrl, &videoUrl, &createdAt, &userList)
		//TODO Check if we can scan entire slice without cycle in order to append
		if err != nil {
			return nil, fmt.Errorf("Error occured selecting lobby by id: %w", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("Any error occure during selecting by id: %w", err.Error())
	}
	fmt.Println(userList)
	return &models.Lobby{
		ID:        id,
		LobbyURL:  lobbyID,
		VideoURL:  videoUrl,
		CreatedAt: createdAt,
		UserList:  userList,
	}, nil
}

func (lr *LobbyRepository) CreateLobby(lobby *models.Lobby) (*models.Lobby, error) {
	query := "INSERT INTO lobbies (lobby_url, video_url, created_at, user_list) VALUES ($1, $2, $3, $4) RETURNING id"
	rows, err := lr.dbHandler.Query(query, lobby.LobbyURL, lobby.VideoURL, lobby.CreatedAt, lobby.UserList)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Error occured while inserting into lobby: %w", err.Error())
	}
	var lobbyId string
	for rows.Next() {
		err := rows.Scan(&lobbyId)
		if err != nil {
			return nil, fmt.Errorf("Error occured while inserting into lobby: %w", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("Any errors occured while inserting into lobby: %w", err.Error())
	}
	return &models.Lobby{
		ID:        lobbyId,
		LobbyURL:  lobby.LobbyURL,
		VideoURL:  lobby.VideoURL,
		CreatedAt: lobby.CreatedAt,
		UserList:  lobby.UserList,
	}, nil
}

func (lr *LobbyRepository) DeleteLobby(lobbyID string) error {
	query := "DELETE FROM lobbies WHERE id = $1"
	res, err := lr.dbHandler.Exec(query, lobbyID)
	if err != nil {
		return fmt.Errorf("Error occured while deleting lobby: %w", err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error occured while deleting lobby: %w", err.Error())
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Lobby with id: %s. NOT FOUND", lobbyID)
	}
	return nil
}

func (lr *LobbyRepository) GetAllLobbies() ([]*models.Lobby, error) {
	query := "SELECT * FROM LOBBIES"
	rows, err := lr.dbHandler.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Error occured while getting all lobbies: %w", err.Error())
	}
	//TODO Проверить, что при сканировании всех лобби будут подтягиваться User'ы
	// Наверное надо посмотреть сложные запросы в постгресе
	lobbies := make([]*models.Lobby, 0)
	var id, lobbyUrl, videoUrl, createdAt string
	for rows.Next() {
		err = rows.Scan()
	}
}

func (lr *LobbyRepository) UpdateLobby() {

}
