package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cargaMaxima int
}
type automotor interface {
	llevarCarga()
}

func main() {
	c1 := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "rojo",
		},
		cargaMaxima: 30000,
	}
	v1 := vehiculo{
		puertas: 2,
		color:   "azul",
	}

	fmt.Println(c1)
	fmt.Println(v1)
	c1.llevarCarga()
	v1.llevarCarga()
	tocarBocina(c1)
	tocarBocina(v1)
}

func tocarBocina(a automotor) {
	switch a.(type) {
	case vehiculo:
		fmt.Printf("Estoy tocando bocina, soy un Vehiculo de color: %v\n", a.(vehiculo).color)
	case camion:
		fmt.Printf("Estoy tocando bocina, soy un Camión de color: %v\n", a.(camion).color)
	}
}

func (c camion) llevarCarga() {
	fmt.Println("Estoy llevando una carga de", c.cargaMaxima, "kilogramos")
}
func (v vehiculo) llevarCarga() {
	fmt.Println("Yo no llevo cargas, solo llevo personas")
}
