package crypto

import "golang.org/x/crypto/bcrypt"

func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)

	return string(bytes), err
}

func TestPassword(h, s string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(h), []byte(s)); err != nil {
		return false, err
	}

	return true, nil
}
