package password_hashes

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ConvertToHashPassword(password string) (string, error) {
	if len([]byte(password)) > 72 {
		return "", errors.New("only 72 bytes or less will be accepted")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), nil
}

func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
