package tokenController

import "auth/internal/model"

type TokenController interface {
	Create(user model.User) (string, error)
	Verify(token string) (bool, error)
}
