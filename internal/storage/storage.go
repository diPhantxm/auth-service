package storage

import (
	"database/sql"

	"auth/internal/storage/Repos"

	log "github.com/go-kit/kit/log"
)

type Storage interface {
	Connect(connStr string)
	Disconnect()
	Users() Repos.UserRepository
}

type AStorage struct {
	UsersRepo Repos.UserRepository
	DB        *sql.DB
	Logger    log.Logger
	Driver    string
}

func (s *AStorage) Connect(connStr string) {
	var err error
	s.DB, err = sql.Open(s.Driver, connStr)

	if err != nil {
		s.Logger.Log("err", err)
	}

	if err = s.DB.Ping(); err != nil {
		s.Logger.Log("err", err)
	}
}

func (s *AStorage) Disconnect() {
	err := s.DB.Close()
	if err != nil {
		s.Logger.Log("err", err)
	}
}

func (s *AStorage) Users() Repos.UserRepository {
	return s.UsersRepo
}
