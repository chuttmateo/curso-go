package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Persona struct {
	Name     string
	LastName string
}

func main() {

	p1 := Persona{
		Name:     "Mateo",
		LastName: "Chutt",
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
	os.Stdout.Write(b)
}
