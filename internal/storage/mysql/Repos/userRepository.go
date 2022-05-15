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

func New(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Auth(login string, pwd string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.QueryRow(
		`SELECT * FROM [auth-service].[dbo].[users] WHERE login=@p1`,
		login,
	).Scan(&user.ID, &user.Login, &user.Password); err != nil {
		return nil, err
	}

	if ok, err := crypto.TestPassword(user.Password, pwd); ok {
		return user, nil
	} else {
		return nil, err
	}
}

func (r UserRepository) ChangePassword(login string, curr, new string) bool {
	if _, err := r.Auth(login, curr); err != nil {
		return false
	}

	newH, err := crypto.Hash(new)
	if err != nil {
		return false
	}

	if err := r.db.QueryRow(
		`UPDATE [auth-service].[dbo].[users] SET password=@p1 WHERE login=@p2`,
		newH,
		login,
	); err.Err() != nil {
		return false
	}

	return true
}

func (r UserRepository) SignUp(user model.User) (string, error) {
	password, err := crypto.Hash(user.Password)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()

	if err := r.db.QueryRow(
		`INSERT INTO [auth-service].[dbo].[users] OUTPUT INSERTED.ID VALUES (@p1, @p2, @p3)`,
		id,
		user.Login,
		password,
	).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r UserRepository) Delete(id string) error {
	if err := r.db.QueryRow(
		`DELETE FROM [auth-service].[dbo].[users] WHERE ID=@p1`,
		id,
	).Err(); err != nil {
		return err
	}

	return nil
}
