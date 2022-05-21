package Repos

import (
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

func (r UserRepository) GetByLogin(login string) (model.User, error) {
	user := model.User{}

	if err := r.db.QueryRow(
		`SELECT * FROM [auth-service].[dbo].[users] WHERE login=@p1`,
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
		`UPDATE [auth-service].[dbo].[users] SET password=@p1 WHERE login=@p2`,
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
		`INSERT INTO [auth-service].[dbo].[users] OUTPUT INSERTED.ID VALUES (@p1, @p2, @p3)`,
		id,
		user.Login,
		user.Password,
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
