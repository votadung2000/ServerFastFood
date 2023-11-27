package components

import "golang.org/x/crypto/bcrypt"

func CheckHashPassword(hashedPassword, passWord string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passWord))
}
