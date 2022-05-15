package Repos

import "auth/internal/model"

type UserRepository interface {
	Auth(login string, pwd string) (*model.User, error)
	ChangePassword(login string, currPwd, newPwd string) bool
	SignUp(user model.User) (string, error)
	Delete(id string) error
}
