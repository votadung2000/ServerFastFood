package components

import "golang.org/x/crypto/bcrypt"

func CheckHash(hashedPassword, passWord string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passWord))
}
