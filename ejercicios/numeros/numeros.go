package numeros

import (
	"fmt"
	"strings"
)

// Implementa una función que determine si un número dado es primo.
func EsPrimo(x int) string {
	if x == 1 {
		return fmt.Sprintln(x, "NO es primo")
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return fmt.Sprintln(x, "NO es primo")
		}
	}
	return fmt.Sprintln(x, "es primo")
}

// Desarrolla un programa que solicite al usuario un número y luego calcule e imprima los primeros n números de la secuencia de Fibonacci.
func SecuenciaFibonacci(x int) []int {
	var z []int
	z = append(z, 0, 1)
	for i := 0; i < (x - 2); i++ {
		z = append(z, z[len(z)-2]+z[len(z)-1])
	}
	return z
}

// Crea una función que tome una cadena como parámetro y devuelva la cantidad de caracteres que tiene.
func Cantidad(str string) int {
	contador := 0
	for range str {
		contador++
	}
	return contador
}

// Implementa un programa que encuentre el número más grande en un arreglo de números enteros.
func MayorDelArreglo(x []int) int {
	var mayor int
	for _, v := range x {
		if v > mayor {
			mayor = v
		}
	}
	return mayor
}

// Desarrolla un programa que tome una cadena de palabras separadas por espacios
// y las imprima en orden inverso.
func ImprimirEnOrdenInverso(str string) string {
	var sb strings.Builder
	sliceString := strings.Split(str, " ")
	for i := len(sliceString) - 1; i >= 0; i-- {
		sb.WriteString(sliceString[i])
		sb.WriteString(" ")
	}
	return sb.String()
}

// Implementa una función que tome un arreglo de enteros como parámetro
// y devuelva la suma de todos los números pares en el arreglo.
func ParesDelArreglo(arr []int) int {
	var suma int
	for _, v := range arr {
		if v%2 == 0 {
			suma += v
		}
	}
	return suma
}

// Crea una función que tome un arreglo de cadenas como parámetro y devuelva la cadena más larga.
func CadenaMasLarga(strs []string) string {
	var cadenaMasLarga string
	for _, v := range strs {
		if len(v) > len(cadenaMasLarga) {
			cadenaMasLarga = v
		}
	}
	return cadenaMasLarga
}

// Desarrolla un programa que genere e imprima los primeros n números primos.
func NNumerosPrimos(n int) []int {
	var numerosPrimos []int
	for i := 2; len(numerosPrimos) < n; i++ {
		v := esPrimoBool(i)
		if v {
			numerosPrimos = append(numerosPrimos, i)
		}
	}
	return numerosPrimos
}
func esPrimoBool(x int) bool {
	if x == 1 {
		return false
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Escribe una función que tome un número entero positivo como parámetro y devuelva el factorial de ese número.
func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

// Implementa una función que tome una cadena como parámetro
// y devuelva true si es un palíndromo y false en caso contrario.
func EsPalindromo(str string) bool {
	strs := strings.Split(str, "")
	for i := 1; i < len(strs); i++ {
		if strs[i-1] != strs[len(strs)-i] {
			return false
		}
	}
	return true
}
