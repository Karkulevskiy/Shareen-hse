package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/karkulevskiy/shareen/src/internal/domain"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"github.com/lib/pq"
)

const (
	postgres = "postgres"
)

type Postgres struct {
	db *sql.DB
}

// MustInitDB creates postgres connection
func MustInit(connectionString string) *Postgres {
	db, err := sql.Open(postgres, connectionString)
	if err != nil {
		panic("failed to initialize Database: " + err.Error())
	}

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) SaveUserUsers(login string) error {
	const op = "storage.postgres.SaveUserUsers"

	stmt, err := p.db.Prepare("INSERT INTO users (login) VALUES ($1)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(login)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok && postgresErr.Constraint != "" {
			return fmt.Errorf("%s: %w", op, storage.ErrUserAlreadyExists)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	totalRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if totalRows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrUserAlreadyExists)
	}

	return nil
}

func (p *Postgres) SaveUser(login string, passHash []byte) error {
	const op = "storage.postgres.SaveUser"

	err := p.SaveUserUsers(login)

	if err != nil {
		return err
	}

	stmt, err := p.db.Prepare("INSERT INTO users_secrets (login, pass_hash) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query(login, passHash)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok && postgresErr.Constraint != "" {
			return fmt.Errorf("%s: %w", op, storage.ErrUserAlreadyExists)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var userID int64

	for rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if rows.Err() != nil {
		return fmt.Errorf("%s: %w", op, rows.Err())
	}

	return nil
}

func (p *Postgres) User(login string) (*domain.User, error) {
	const op = "storage.Postgres.User"

	stmt, err := p.db.Prepare("SELECT * FROM users_secrets WHERE login = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRow(login)

	var user domain.User

	err = row.Scan(&user.ID, &user.Login, &user.PassHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &domain.User{
		ID:       user.ID,
		Login:    user.Login,
		PassHash: user.PassHash,
	}, nil
}

func (p *Postgres) CreateLobby(lobbyURL string) (string, error) {
	const op = "postgres.CreateLobby"

	stmt, err := p.db.Prepare("INSERT INTO lobbies (lobby_url) VALUES ($1)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(lobbyURL)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok && postgresErr.Constraint != "" {
			return "", fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if totalRows, err := res.RowsAffected(); err != nil || totalRows == 0 {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return lobbyURL, nil
}

func (p *Postgres) Lobby(lobbyURL string) (*domain.Lobby, error) {
	const op = "postgres.Lobby"

	stmt, err := p.db.Prepare("SELECT * FROM lobbies WHERE lobby_url = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows := stmt.QueryRow(lobbyURL)

	var lobbyID int64
	var lobbyURL_ string
	var videoURL sql.NullString
	var pause bool
	var timing sql.NullString

	err = rows.Scan(&lobbyID, &lobbyURL_, &videoURL, &pause, &timing)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	lobby := &domain.Lobby{
		LobbyURL: lobbyURL_,
		VideoURL: videoURL.String,
		Pause:    pause,
		Timing:   timing.String,
	}

	chat, err := p.Chat(lobbyID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	lobby.Chat = chat

	return lobby, nil
}

func (p *Postgres) Chat(lobbyID int64) ([]domain.Message, error) {
	const op = "storage.postgres.Chat"

	stmt, err := p.db.Prepare("SELECT user_login, time, message FROM chats WHERE lobby_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query(lobbyID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var chat []domain.Message

	var login, message string
	var time_ sql.NullTime

	for rows.Next() {
		err = rows.Scan(&login, &time_, &message)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		chat = append(chat, domain.Message{
			Login:   login,
			Message: message,
			Time:    time_.Time,
		})
	}

	return chat, nil
}

func (p *Postgres) InsertVideo(lobbyURL, videoURL string) error {
	const op = "storage.postgres.InsertVideo"

	stmt, err := p.db.Prepare("UPDATE lobbies SET video_url = $1 WHERE lobby_url = $2")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(videoURL, lobbyURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	totalRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if totalRows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
	}

	return nil
}

func (p *Postgres) SaveMessage(lobbyURL, login, message string) error {
	const op = "storage.postgres.SaveMessage"

	lobbyID, err := p.LobbyID(lobbyURL)
	if err != nil {
		return err
	}

	stmt, err := p.db.Prepare("INSERT INTO chats (lobby_id, user_login, time, message) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(lobbyID, login, time.Now(), message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (p *Postgres) LobbyID(lobbyURL string) (int64, error) {
	const op = "storage.postgres.LobbyID"

	stmt, err := p.db.Prepare("SELECT id FROM lobbies WHERE lobby_url = $1")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	rows := stmt.QueryRow(lobbyURL)

	var lobbyID int64

	err = rows.Scan(&lobbyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrLobbyNotFound)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return lobbyID, nil
}
