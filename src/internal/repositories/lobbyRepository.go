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

	row := l.db.QueryRow(query, lobbyURL)

	var lobbyID string

	err := row.Scan(&lobbyID)

	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok && postgresErr.Constraint != "" {
			return "", fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
	}

	return lobbyID, nil
}

func (l *LobbyRepository) GetLobby(lobbyURL string) (*models.Lobby, error) {
	const op = "repositories.lobbyRepository.GetLobby"
	const query = `SELECT lobbies.id, video_url, users.id, name
	 			FROM lobbies
				LEFT JOIN lobbies_users ON lobbies.id = lobbies_users.lobby_id
				LEFT JOIN users ON users.id = lobbies_users.user_id
				WHERE lobby_url = $1`

	rows, err := l.db.Query(query, lobbyURL)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var lobbyID, userID int64
	var videoURL, userName string

	var userList []models.User

	for rows.Next() {
		err = rows.Scan(&lobbyID, &videoURL, &userID, &userName)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if userID == 0 {
			continue
		}

		userList = append(userList, models.User{
			ID:   userID,
			Name: userName,
		})
	}

	return &models.Lobby{
		ID:        lobbyID,
		VideoURL:  videoURL,
		LobbdyURL: lobbyURL,
		Users:     userList,
	}, nil
}

func (l *LobbyRepository) DeleteLobby(lobbyID string) error {
	const op = "repositories.lobbyRepository.DeleteLobby"
	const query = "DELETE FROM lobbies WHERE id = $1"

	res, err := l.db.Exec(query, lobbyID)
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

func (l *LobbyRepository) UpdateLobbyVideoURL(lobbyID string, videoURL string) error {
	const op = "repositories.lobbyRepository.UpdateLobbyVideoURL"
	const query = "UPDATE lobbies SET video_url = $1 WHERE id = $2"

	res, err := l.db.Exec(query, videoURL, lobbyID)
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

func (l *LobbyRepository) JoinUserToLobby(lobbyID string, userID string) error {
	const op = "repositories.lobbyRepository.JoinUserToLobby"
	const query = "INSERT INTO lobbies_users (lobby_id, user_id) VALUES ($1, $2)"

	res, err := l.db.Exec(query, lobbyID, userID)

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
