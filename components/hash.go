package components

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytesHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytesHash), err
}
