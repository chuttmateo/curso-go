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
	fmt.Println("es primo", numeros.EsPrimo(5))
	fmt.Println("no es primo", numeros.EsPrimo(6))
	fmt.Println("es primo", numeros.EsPrimo(7))
	fmt.Println("no es primo", numeros.EsPrimo(8))
	fmt.Println("no es primo", numeros.EsPrimo(9))
	fmt.Println("no es primo", numeros.EsPrimo(10))
	fmt.Println("es primo", numeros.EsPrimo(11))
	fmt.Println("no es primo", numeros.EsPrimo(12))
	fmt.Println("es primo", numeros.EsPrimo(13))
	fmt.Println("no es primo", numeros.EsPrimo(14))
	fmt.Println("no es primo", numeros.EsPrimo(15))
	fmt.Println("no es primo", numeros.EsPrimo(16))
	fmt.Println("es primo", numeros.EsPrimo(17))
	fmt.Println("no es primo", numeros.EsPrimo(18))
	fmt.Println("es primo", numeros.EsPrimo(19))
	fmt.Println("no es primo", numeros.EsPrimo(20))

	fmt.Println("Secuencia Fibonacci --- ", numeros.SecuenciaFibonacci(10))
	fmt.Println("Mayor del arreglo --- ", numeros.MayorDelArreglo([]int{2, 5, 100, 3, 80}))
	fmt.Println("Imprimir En Orden Inverso --- ", numeros.ImprimirEnOrdenInverso("Hola soy Mateo"))
	fmt.Println("Pares del arreglo --- ", numeros.ParesDelArreglo([]int{2, 3, 4, 6, 1, 1, 1}))
	fmt.Println("Cadena mas larga --- ", numeros.CadenaMasLarga([]string{"aaa", "b", "cccc"}))
}
