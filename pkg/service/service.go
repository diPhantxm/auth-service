package service

import (
	"auth/internal/model"
	"auth/internal/tokenController"
	"context"

	"auth/internal/storage"
	//dbDriver "auth/internal/storage/mysql"
	dbDriver "auth/internal/storage/postgres"
)

// AuthService describes the service.
type AuthService interface {
	Auth(ctx context.Context, login string, password string) (jwt string, err error)
	ChangePassword(ctx context.Context, login string, curr string, new string) (success bool, err error)
	SignUp(ctx context.Context, user model.User) (id string, err error)
	Delete(ctx context.Context, id string) (err error)
}

type basicAuthService struct {
	storage         storage.Storage
	TokenController tokenController.TokenController
}

func (b *basicAuthService) Auth(ctx context.Context, login string, password string) (token string, err error) {
	user, err := b.storage.Users().Auth(login, password)

	if err != nil {
		token = ""
		return token, nil
	}

	return b.TokenController.Create(*user)
}
func (b *basicAuthService) ChangePassword(ctx context.Context, login string, curr string, new string) (success bool, err error) {
	success = b.storage.Users().ChangePassword(login, curr, new)

	return success, err
}
func (b *basicAuthService) SignUp(ctx context.Context, user model.User) (id string, err error) {
	return b.storage.Users().SignUp(user)
}
func (b *basicAuthService) Delete(ctx context.Context, id string) (err error) {
	return b.storage.Users().Delete(id)
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService(dbConn string) AuthService {
	return &basicAuthService{
		storage: dbDriver.New(dbConn),
	}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(dbConn string, tokenController tokenController.TokenController, middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService(dbConn)

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
