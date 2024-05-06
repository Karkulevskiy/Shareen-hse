package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/karkulevskiy/shareen/src/internal/domain"
	"github.com/karkulevskiy/shareen/src/internal/storage"
	"github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

// MustInitDB - initializes DB
func MustInitDB(connectionString string) *Postgres {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to init db: " + err.Error())
	}

	prepareDB(db)

	return &Postgres{
		db: db,
	}
}

// prepareDB - describes prepared statement for DB initializing
func prepareDB(db *sql.DB) {
	const (
		lobbiesStmt = `
			CREATE TABLE IF NOT EXISTS lobbies
			(
				id SERIAL PRIMARY KEY,
				lobby_url varchar(255) UNIQUE NOT NULL,
				video_url varchar(255)
			);`
		index     = `CREATE INDEX IF NOT EXISTS lobby_url_idx ON lobbies (lobby_url);`
		usersStmt = `
			CREATE TABLE IF NOT EXISTS users
			(
				id SERIAL PRIMARY KEY,
				name VARCHAR(20) NOT NULL
			);`
		lobbiesUsersStmt = `
		CREATE TABLE IF NOT EXISTS lobbies_users
(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES users ON DELETE CASCADE, --Проверить поведение БД, при удалении пользователя и лобби
    lobby_url VARCHAR(255) REFERENCES lobbies (lobby_url) ON DELETE CASCADE,
    UNIQUE(user_id, lobby_url)
);`
		chatsStmt = `
		CREATE TABLE IF NOT EXISTS chats 
		(
			id SERIAL PRIMARY KEY,
			lobby_url VARCHAR(255) REFERENCES lobbies (lobby_url) ON DELETE CASCADE
		);`
	)

	for _, query := range []string{lobbiesStmt, index, usersStmt, lobbiesUsersStmt, chatsStmt} {
		execStmt(db, query)
	}
}

// execStmt executes prepared statement
func execStmt(db *sql.DB, query string) {
	const op = "storage.postgres.execStmt"

	stmt, err := db.Prepare(query)

	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	_, err = stmt.Exec()

	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
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

func (p *Postgres) CreateLobby(lobbyURL string) error {

}

func (p *Postgres) JoinLobby() {}
