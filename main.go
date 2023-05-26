package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("main")
	x := strings.Split("hola a todos", " ")
	fmt.Println(x)
	y := strings.Join(x, " ")
	fmt.Println(y)

	var paises []Pais = []Pais{{"sdf", 23}, {"sdf", 23}, {"sdf", 23}}
	fmt.Println(paises)
	paises = append(paises, Pais{"Argentina", 50_000_000})
	fmt.Println(paises[len(paises)-1])
	uuid := uuid.New()
	fmt.Println(uuid)
}

type Pais struct {
	Nombre    string
	Poblacion int
}
