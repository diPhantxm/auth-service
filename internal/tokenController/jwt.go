package tokenController

import (
	"auth/internal/model"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var key = "KEY"

type JWTController struct {
}

func NewJWTController() JWTController {
	return JWTController{}
}

func (c JWTController) Create(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"login": user.Login,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (c JWTController) Verify(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return key, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if time.Unix(claims["exp"].(int64), 0).After(time.Now()) {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, fmt.Errorf("Token is not valid")
	}
}
