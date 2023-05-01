package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "hola"
	sb , err :=bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))
	
}


