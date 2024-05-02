package postgres

import "database/sql"

type Storage struct {
	db *sql.DB
}

func MustInitDB(connectionString string) *Storage {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to init db: " + err.Error())
	}

	prepareDB()
}

func prepareDB() {
	const firstPreparedStmt = ``
}
