package Repos

import (
	mysql "auth/pkg/storage/mysql/Repos"
	postgres "auth/pkg/storage/postgres/Repos"

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
