package common

import "fmt"

const (
	CurrentUser = "current_user"
	JWT         = "jwt"
)

func RecoverConst() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}
