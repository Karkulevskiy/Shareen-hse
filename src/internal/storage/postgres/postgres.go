package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct{}

// MustInitDB - initializes DB
func MustInitDB(connectionString string) *Storage {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to init db: " + err.Error())
	}

	prepareDB(db)

	return &Storage{}
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
