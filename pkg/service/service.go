package service

import (
	"auth/internal/crypto"
	"auth/internal/model"
	"auth/internal/tokenController"
	"context"
	"fmt"
	"log"

	"auth/internal/storage"
)

// AuthService describes the service.
type AuthService interface {
	Auth(ctx context.Context, login string, password string) (jwt string, err error)
	ChangePassword(ctx context.Context, login string, curr string, new string) (success bool, err error)
	SignUp(ctx context.Context, user model.User) (id string, err error)
	Delete(ctx context.Context, id string) (err error)
}

type basicAuthService struct {
	storage         storage.IStorage
	TokenController tokenController.TokenController
}

func (b *basicAuthService) Auth(ctx context.Context, login string, password string) (token string, err error) {
	user, err := b.storage.Users().GetByLogin(login)
	if err != nil {
		return "", err
	}

	match, err := crypto.TestPassword(user.Password, password)
	if err != nil {
		return "", err
	}

	if !match {
		return "", fmt.Errorf("Password is incorrect")
	}

	return b.TokenController.Create(user)
}
func (b *basicAuthService) ChangePassword(ctx context.Context, login string, curr string, new string) (success bool, err error) {
	success = b.storage.Users().ChangePassword(login, curr, new)

	return success, err
}
func (b *basicAuthService) SignUp(ctx context.Context, user model.User) (id string, err error) {
	hashedPassword, err := crypto.Hash(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword

	return b.storage.Users().SignUp(user)
}
func (b *basicAuthService) Delete(ctx context.Context, id string) (err error) {
	return b.storage.Users().Delete(id)
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService(driver string, connectionString string) AuthService {
	storage, err := storage.Connect(driver, connectionString)
	if err != nil {
		log.Fatalf("Storage didn't connect to database. Error: %s\n", err.Error())
	}

	basicSvc := &basicAuthService{
		storage: storage,
	}

	return basicSvc
}

// New returns a AuthService with all of the expected middleware wired in.
func New(driver string, connectionString string, tokenController tokenController.TokenController, middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService(driver, connectionString)

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
