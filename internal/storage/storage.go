package storage

import (
	"database/sql"
	"fmt"

	"auth/internal/storage/Repos"

	log "github.com/go-kit/kit/log"
)

type IStorage interface {
	Disconnect()
	Users() Repos.UserRepository
}

type storage struct {
	UsersRepo Repos.UserRepository
	DB        *sql.DB
	Logger    log.Logger
}

func Connect(driver string, connStr string) (IStorage, error) {
	db, err := sql.Open(driver, connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	repo := Repos.CreateUserRepo(driver, db)
	if repo == nil {
		return nil, fmt.Errorf("Factory didn't manage to create User Repository\n")
	}

	return &storage{
		DB:        db,
		UsersRepo: repo,
	}, nil
}

func (s *storage) Disconnect() {
	err := s.DB.Close()
	if err != nil {
		s.Logger.Log("err", err)
	}
}

func (s *storage) Users() Repos.UserRepository {
	return s.UsersRepo
}
