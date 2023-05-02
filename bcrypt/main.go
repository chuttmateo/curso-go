package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "hola"
	passhasheada, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(passhasheada))

	err = bcrypt.CompareHashAndPassword(passhasheada, []byte("hola"))
	if err != nil {
		fmt.Println("La clave y el hash code no coinciden")
		return
	}
	fmt.Println("la clave es correcta")
}