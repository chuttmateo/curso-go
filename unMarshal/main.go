package main

import (
	"encoding/json"
	"fmt"
)

type pais struct {
	Nombre           string  `json:"nombre"`
	Tamano           string  `json:"tamano"`
	Poblacion        float64 `json:"poblacion"`
	MundialesGanados int     `json:"mundiales_ganados"`
}

func main() {
	paisesjson := `[
    {
      "nombre": "Brasil",
      "tamano": "8.515.767 km²",
      "poblacion": 213.3,
      "mundiales_ganados": 5
    },
    {
      "nombre": "Argentina",
      "tamano": "2.780.400 km²",
      "poblacion": 45.5,
      "mundiales_ganados": 3
    },
    {
      "nombre": "Alemania",
      "tamano": "357.386 km²",
      "poblacion": 83.2,
      "mundiales_ganados": 4
    },
    {
      "nombre": "Italia",
      "tamano": "301.340 km²",
      "poblacion": 60.4,
      "mundiales_ganados": 4
    },
    {
      "nombre": "Francia",
      "tamano": "643.801 km²",
      "poblacion": 66.99,
      "mundiales_ganados": 2
    }
  ]`
	var paises []pais

	err := json.Unmarshal([]byte(paisesjson), &paises)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", paises)
	for _, v := range paises {
		fmt.Println(v)
	}

}
