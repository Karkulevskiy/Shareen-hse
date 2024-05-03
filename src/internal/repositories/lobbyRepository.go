package repositories

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/shareen/src/internal/domain/models"
	"github.com/shareen/src/internal/storage"
)

// NewLobbyRepository is a constructor
func NewLobbyRepository(db *sql.DB) *LobbyRepository {
	return &LobbyRepository{
		db: db,
	}
}

type LobbyRepository struct {
	db *sql.DB
}

// CreateLobby created new lobby and returning lobbdy
func (l *LobbyRepository) CreateLobby(lobbyURL string) (string, error) {
	const op = "repositories.lobbyRepository.CreateLobby"
	const query = "INSERT INTO lobbies (lobby_url) VALUES ($1)"

	res, err := l.db.Exec(query, lobbyURL)

	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok && postgresErr.Constraint != "" {
			return "", fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
	}

	if totalRows, err := res.RowsAffected(); err != nil || totalRows == 0 {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return lobbyURL, nil
}

func (l *LobbyRepository) GetLobby(lobbyURL string) (*models.Lobby, error) {
	const op = "repositories.lobbyRepository.GetLobby"
	const query = `SELECT lobbies.id, video_url, users.id, name
	 			FROM lobbies
				LEFT JOIN lobbies_users ON lobbies.lobby_url = lobbies_users.lobby_url
				LEFT JOIN users ON users.id = lobbies_users.user_id
				WHERE lobbies.lobby_url = $1`

	rows, err := l.db.Query(query, lobbyURL)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var lobbyID, userID sql.NullInt64
	var userName sql.NullString
	var videoURL sql.NullString
	var userList []models.User

	for rows.Next() {
		err = rows.Scan(&lobbyID, &videoURL, &userID, &userName)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if userID.Int64 == 0 {
			continue
		}

		userList = append(userList, models.User{
			ID:   userID.Int64,
			Name: userName.String,
		})
	}

	if lobbyID.Int64 == 0 {
		return nil, fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
	}

	return &models.Lobby{
		ID:        lobbyID.Int64,
		VideoURL:  videoURL.String,
		LobbdyURL: lobbyURL,
		Users:     userList,
	}, nil
}

func (l *LobbyRepository) DeleteLobby(lobbyURL string) error {
	const op = "repositories.lobbyRepository.DeleteLobby"
	const query = "DELETE FROM lobbies WHERE lobby_url = $1"

	res, err := l.db.Exec(query, lobbyURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
	}

	return nil
}

func (l *LobbyRepository) UpdateLobbyVideoURL(lobbyURL string, videoURL string) error {
	const op = "repositories.lobbyRepository.UpdateLobbyVideoURL"
	const query = "UPDATE lobbies SET video_url = $1 WHERE lobby_url = $2"

	res, err := l.db.Exec(query, videoURL, lobbyURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
	}

	return nil
}
