package Repos

import (
	"auth/internal/crypto"
	"auth/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func (r UserRepository) Auth(login string, pwd string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.QueryRow(
		`SELECT * FROM users WHERE login=$1`,
		login,
	).Scan(&user.ID, &user.Login, &user.Password); err != nil {
		return nil, err
	}

	if ok, res := crypto.TestPassword(user.Password, pwd); ok {
		return user, nil
	} else {
		return nil, res
	}
}

func (r UserRepository) ChangePassword(login string, currPwd, newPwd string) bool {
	if _, err := r.Auth(login, currPwd); err != nil {
		return false
	}

	newPwdHashed, err := crypto.Hash(newPwd)
	if err != nil {
		return false
	}

	if err := r.db.QueryRow(
		`UPDATE users SET password=$1 WHERE login=$2`,
		newPwdHashed,
		login,
	); err.Err() != nil {
		return false
	}

	return true
}

func (r UserRepository) SignUp(user model.User) (string, error) {
	passwordHashed, err := crypto.Hash(user.Password)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()

	if err := r.db.QueryRow(
		`INSERT INTO users (id, login, password) VALUES ($1, $2, $3)`,
		id,
		user.Login,
		passwordHashed,
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
