package Repos

import (
	"auth/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func (r UserRepository) GetByLogin(login string) (model.User, error) {
	user := model.User{}

	if err := r.db.QueryRow(
		`SELECT * FROM users WHERE login=$1`,
		login,
	).Scan(&user.ID, &user.Login, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepository) ChangePassword(login string, currentPassword, newPassword string) bool {
	user, err := r.GetByLogin(login)
	if err != nil {
		return false
	}

	if user.Password != currentPassword {
		return false
	}

	if err := r.db.QueryRow(
		`UPDATE users SET password=$1 WHERE login=$2`,
		newPassword,
		login,
	); err.Err() != nil {
		return false
	}

	return true
}

func (r UserRepository) SignUp(user model.User) (string, error) {
	id := uuid.New().String()

	if err := r.db.QueryRow(
		`INSERT INTO users (id, login, password) VALUES ($1, $2, $3)`,
		id,
		user.Login,
		user.Password,
	); err.Err() != nil {
		return "", err.Err()
	}

	return id, nil
}

func (r UserRepository) Delete(id string) error {
	if err := r.db.QueryRow(
		`DELETE FROM users WHERE id=$1`,
		id,
	); err.Err() != nil {
		return err.Err()
	}

	return nil
}

func New(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}
