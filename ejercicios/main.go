package main

import (
	"fmt"

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
}
