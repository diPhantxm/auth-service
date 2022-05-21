package Repos

import (
	"auth/internal/model"
	"database/sql"
	"math/rand"
	"testing"

	repoDriver "auth/internal/storage/mysql/Repos"

	_ "github.com/denisenkom/go-mssqldb"
)

var driver = "sqlserver"
var connectionString = "auth-service:auth-service@192.168.0.179:3306/MSSQLSERVER?database=auth-service"

func TestRepository(t *testing.T) {
	testUsers := []model.User{
		{Login: "diPhantxm", Password: "123456"},
		{Login: "test", Password: "test"},
		{Login: "123", Password: "123"},
		{Login: "-_-", Password: "-_-"},
		{Login: "OneMoreTest", Password: ""},
		{Login: "t", Password: "t"},
		{Login: "_", Password: "_"},
	}

	dbConnection := connectDB(t)
	repo := repoDriver.New(dbConnection)

	for _, user := range testUsers {
		create(&user, repo, t)
		read(&user, repo, t)
		updatePassword(&user, repo, t)
		delete(&user, repo, t)
	}
}

func create(user *model.User, repo UserRepository, t *testing.T) {
	var err error

	user.ID, err = repo.SignUp(*user)
	if err != nil {
		t.Errorf("Test failed on Create. Error: %s\n", err.Error())
	}
}

func read(user *model.User, repo UserRepository, t *testing.T) {
	gotUser, err := repo.GetByLogin(user.Login)
	if err != nil {
		t.Errorf("Test failed on Read. Error: %s\n", err.Error())
	}

	if gotUser.Login != user.Login || gotUser.Password != user.Password {
		t.Errorf("Test failed on Read. Got wrong user\n")
	}
}

func updatePassword(user *model.User, repo UserRepository, t *testing.T) {
	newPassword := generateRandomPassword()
	success := repo.ChangePassword(user.Login, user.Password, newPassword)
	if !success {
		t.Errorf("Test failed on Update\n")
	}

	user.Password = newPassword
}

func delete(user *model.User, repo UserRepository, t *testing.T) {
	err := repo.Delete(user.ID)
	if err != nil {
		t.Errorf("Test failed on Delete. Error: %s\n", err.Error())
	}
}

func connectDB(t *testing.T) *sql.DB {
	db, err := sql.Open(driver, connectionString)

	if err != nil {
		t.Fatalf("Test failed with no db connection. Error: %s\n", err.Error())
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Test failed with no db connection. Error: %s\n", err.Error())
	}

	return db
}

func generateRandomPassword() string {
	const size = 10
	password := make([]rune, size)

	for i := range password {
		password[i] = rune('a' + rand.Intn(26)) // English alphabet size
	}

	return string(password)
}
