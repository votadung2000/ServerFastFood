package common

import "fmt"

func RecoverConst() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}
