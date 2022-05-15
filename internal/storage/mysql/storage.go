package mysql

import (
	"auth/internal/storage"
	"auth/internal/storage/mysql/Repos"

	_ "github.com/denisenkom/go-mssqldb"
)

type Storage struct {
	storage.AStorage
}

func New(connStr string) *Storage {
	s := &Storage{}

	s.Driver = "sqlserver"

	s.Connect(connStr)
	s.UsersRepo = Repos.New(s.DB)

	return s
}

func (s Storage) CreateRepos() {
	s.UsersRepo = Repos.New(s.DB)
}
