package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}
type porApellido []usuario

func (s porApellido) Len() int {
	return len(s)
}
func (s porApellido) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s porApellido) Less(i, j int) bool {
	return s[i].Apellido < s[j].Apellido
}

type porEdad []usuario

func (s porEdad) Len() int {
	return len(s)
}
func (s porEdad) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s porEdad) Less(i, j int) bool {
	return s[i].Edad < s[j].Edad
}

type StringSlice []string
func (s StringSlice) Len() int {
	return len(s)
}
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Pepe diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println(usuarios)

	// tu código va aquí
	sort.Sort(porEdad(usuarios))
	
	fmt.Println("\t\tPor Edad")
	for _, v := range usuarios{
		fmt.Printf("Nombre: %v\n", v.Nombre)
		fmt.Printf("Apellido: %v\n", v.Apellido)
		fmt.Printf("Edad: %v\n", v.Edad)
		sort.Strings(v.Dichos)
		for _, dicho := range v.Dichos{
			fmt.Printf("\t%v\n",dicho)
		} 
	}
	sort.Sort(porApellido(usuarios))
	fmt.Println("\t\tPor Apellido")
	for _, v := range usuarios{
		fmt.Printf("Nombre: %v\n", v.Nombre)
		fmt.Printf("Apellido: %v\n", v.Apellido)
		fmt.Printf("Edad: %v\n", v.Edad)
		if sort.IsSorted(StringSlice(v.Dichos)) {
			fmt.Println("Los dichos están ordenados")
		}
		for _, dicho := range v.Dichos{
			fmt.Printf("\t%v\n",dicho)
		} 
	}

}
