package Repos

import (
	mysql "auth/internal/storage/mysql/Repos"
	postgres "auth/internal/storage/postgres/Repos"

	"database/sql"
)

func CreateUserRepo(driver string, db *sql.DB) UserRepository {
	switch driver {
	case "sqlserver":
		return mysql.New(db)
	case "postgres":
		return postgres.New(db)
	default:
		return nil
	}
}
