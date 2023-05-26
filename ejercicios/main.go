package main

import (
	"fmt"
	"time"

	"github.com/mateobr0/curso.go/ejercicios/numeros"
)

func main() {
	fmt.Println("no es primo", numeros.EsPrimo(1))
	fmt.Println("es primo", numeros.EsPrimo(2))
	fmt.Println("es primo", numeros.EsPrimo(3))
	fmt.Println("no es primo", numeros.EsPrimo(4))
	fmt.Println("Secuencia Fibonacci --- ", numeros.SecuenciaFibonacci(10))
	fmt.Println("Mayor del arreglo --- ", numeros.MayorDelArreglo([]int{2, 5, 100, 3, 80}))
	fmt.Println("Imprimir En Orden Inverso --- ", numeros.ImprimirEnOrdenInverso("Hola soy Mateo"))
	fmt.Println("Pares del arreglo --- ", numeros.ParesDelArreglo([]int{2, 3, 4, 6, 1, 1, 1}))
	fmt.Println("Cadena mas larga --- ", numeros.CadenaMasLarga([]string{"aaa", "b", "cccc"}))
	fmt.Println("Primeros 10 números primos --- ", numeros.NNumerosPrimos(10))
	fmt.Println("Factorial de 10 es --- ", numeros.Factorial(10))
	fmt.Println("¿Neuquen es palindromo? --- ", numeros.EsPalindromo("neuquen"))
	fmt.Println("Ordenar lista de strings --- ", numeros.Ordenar([]string{"mateo", "antonella", "lucas"}))
	pares, impares := numeros.SumaParesEImpares([]int{4, 4, 2, 3})
	fmt.Printf("Suma de números pares %v e impares %v\n", pares, impares)
	fmt.Println("Números perfectos", numeros.EsPerfecto(6))
	fmt.Println("Números perfectos", numeros.EsPerfecto(8128))
	t1 := time.Now()
	fmt.Println(numeros.ElementosComunes([]int{1, 2, 0, 4, 5, 6, 3, 2, 32, 2, 324, 234, 234, 234, 234, 2, 34, 234, 234, 23, 4, 234, 23, 4, 234, 234, 2, 34}, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(time.Since(t1))
	t2 := time.Now()
	fmt.Println(numeros.ElementosComunesMap([]int{1, 2, 0, 4, 5, 6, 3, 2, 32, 2, 324, 234, 234, 234, 234, 2, 34, 234, 234, 23, 4, 234, 23, 4, 234, 234, 2, 34}, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(time.Since(t2))
	//en la funcion con el uso de map se ve una pequeña diferencia en tiempo respecto a la que no utiliza map
	matriz := [][]int{
		{1,1,1},
		{1,1,1},
		{1,1,1},
	}
	fmt.Println(numeros.EsSimetrica(matriz))
	
}
