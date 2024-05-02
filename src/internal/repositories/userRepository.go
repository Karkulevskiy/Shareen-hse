package repositories

import (
	"database/sql"
	"fmt"

	"github.com/shareen/src/internal/domain/models"
	"github.com/shareen/src/internal/storage"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository is a constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// TODO: доделать авторизацию пользователей
// ВРЕМЕННОЙ РЕШЕНИЕ
// CreateUser creates user and returning userID
func (u *UserRepository) CreateUser(name string) (int64, error) {
	const op = "repositories.userRepository.CreateUser"
	const query = "INSERT INTO users (name) VALUES ($1)"

	row := u.db.QueryRow(query, name)

	var userID int64

	err := row.Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}

// GetUser returns user by userID
func (u *UserRepository) GetUser(userID int64) (*models.User, error) {
	const op = "repositories.userRepository.GetUser"
	const query = "SELECT name FROM users WHERE id = $1"

	row := u.db.QueryRow(query, userID)

	var name string

	err := row.Scan(&name)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &models.User{
		ID:   userID,
		Name: name,
	}, nil
}

// DeleteUser deletes user by userID
func (u *UserRepository) DeleteUser(userID int64) error {
	const op = "repositories.userRepository.DeleteUser"
	const query = "DELETE FROM users WHERE id = $1"

	res, err := u.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return nil
}
