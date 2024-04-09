package repositories

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// TODO Посмотреть документацию gin, swagger_gint

import (
	"database/sql"
	"fmt"
	"log"
	"shareen/src/models"
)

type LobbiesRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewLobbiesRepository(dbHandler *sql.DB) *LobbiesRepository {
	return &LobbiesRepository{
		dbHandler: dbHandler,
	}
}

func (lr *LobbiesRepository) GetLobby(lobbyID string) (*models.Lobby, error) {
	query := "SELECT * FROM lobbies WHERE id = $1"
	rows, err := lr.dbHandler.Query(query, lobbyID)
	if err != nil {
		return nil, fmt.Errorf("error occured selecting lobby by id: %s", err.Error())
	}
	defer rows.Close()
	var id, lobbyUrl, videoUrl, createdAt string
	userList := make([]*models.User, 0)
	for rows.Next() {
		err = rows.Scan(&id, &lobbyUrl, &videoUrl, &createdAt, &userList)
		//TODO Check if we can scan entire slice without cycle in order to append
		if err != nil {
			return nil, fmt.Errorf("error occured selecting lobby by id: %s", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("any error occure during selecting by id: %s", err.Error())
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

func (lr *LobbiesRepository) CreateLobby(lobby *models.Lobby) (*models.Lobby, error) {
	query := "INSERT INTO lobbies (lobby_url, video_url, created_at) VALUES ($1, $2, $3) RETURNING id"
	rows, err := lr.dbHandler.Query(query, lobby.LobbyURL, lobby.VideoURL, lobby.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error occured while inserting into lobby: %s", err.Error())
	}
	defer rows.Close()
	var lobbyId string
	for rows.Next() {
		err := rows.Scan(&lobbyId)
		if err != nil {
			return nil, fmt.Errorf("error occured while inserting into lobby: %s", err.Error())
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("any errors occured while inserting into lobby: %s", rows.Err())
	}
	return &models.Lobby{
		ID:        lobbyId,
		LobbyURL:  lobby.LobbyURL,
		VideoURL:  lobby.VideoURL,
		CreatedAt: lobby.CreatedAt,
		UserList:  lobby.UserList,
	}, nil
}

func (lr *LobbiesRepository) DeleteLobby(lobbyID string) error {
	query := "DELETE FROM lobbies WHERE id = $1"
	res, err := lr.dbHandler.Exec(query, lobbyID)
	if err != nil {
		return fmt.Errorf("error occured while deleting lobby: %s", err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error occured while deleting lobby: %s", err.Error())
	}
	if rowsAffected == 0 {
		return fmt.Errorf("lobby with id: %s. NOT FOUND", lobbyID)
	}
	return nil
}

func (lr *LobbiesRepository) GetAllLobbies() ([]*models.Lobby, error) {
	query := "SELECT * FROM lobbies l LEFT JOIN users u on l.id = u.lobby_id"
	rows, err := lr.dbHandler.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting all lobbies: %s", err.Error())
	}
	defer rows.Close()
	var lobbyId, lobbyURL, videoURL, createdAt string
	var userID, userName, userLobbyID sql.NullString
	lobbies_users := map[*models.Lobby][]*models.User{}
	for rows.Next() {
		err = rows.Scan(&lobbyId, &lobbyURL, &videoURL, &createdAt, &userID, &userName, &userLobbyID)
		if err != nil {
			log.Println(err.Error())
			return nil, fmt.Errorf("error occured while scanning lobbies")
		}
		lobby := &models.Lobby{
			ID:        lobbyId,
			LobbyURL:  lobbyURL,
			VideoURL:  videoURL,
			CreatedAt: createdAt,
		}
		user := &models.User{
			ID:      userID.String,
			LobbyID: userLobbyID.String,
			Name:    userName.String,
		}
		lobbies_users[lobby] = append(lobbies_users[lobby], user)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("any error occured while scannig lobbies")
	}
	lobbies := make([]*models.Lobby, len(lobbies_users))
	for lKey, uVal := range lobbies_users {
		lKey.UserList = uVal
		lobbies = append(lobbies, lKey)
	}
	return lobbies, nil
}

func (lr *LobbiesRepository) UpdateLobby() {

}
