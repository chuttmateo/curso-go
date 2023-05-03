package main

import (
	"fmt"

	customtypes "github.com/mateobr0/curso.go/generics/ejercicio2/customtypes"
	fakedb "github.com/mateobr0/curso.go/generics/ejercicio2/fakeDB"
)


func main() {
	persons := fakedb.Read[[]customtypes.Person]("persons")
	for i, v := range persons {
		fmt.Printf("Persona num: %v de nombre: %v y edad %v\n", i, v.Name, v.Age)
	}
	employees := fakedb.Read[[]customtypes.Employee]("employees")
	for i, v := range employees {
		fmt.Printf("Empleado num: %v de nombre: %v y puesto %v\n", i, v.Name, v.Position)
	}
}