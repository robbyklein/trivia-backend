package helpers

import "golang.org/x/crypto/bcrypt"

func ComparePasswordToHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
