package postgres

import (
	"auth/internal/storage"
	"auth/internal/storage/postgres/Repos"

	_ "github.com/lib/pq"
)

type Storage struct {
	storage.AStorage
}

func New(connStr string) *Storage {
	s := &Storage{
		//Logger: log.New(),
	}

	s.Driver = "postgres"

	s.Connect(connStr)
	s.UsersRepo = Repos.New(s.DB)

	return s
}

func (s Storage) CreateRepos() {
	s.UsersRepo = Repos.New(s.DB)
}
