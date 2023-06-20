package numeros

import (
	"math/rand"
	"strings"
)

// Calculadora de propinas: Escribe un programa que calcule el monto de la propina en función
// del total de la cuenta y un porcentaje de propina dado.
func CalcularPropinas(monto, porcentaje float32) float32 {
	return monto * (porcentaje / 100)
}

// Conversor de unidades: Crea un programa que convierta una cantidad en una unidad de medida a otra.
// Por ejemplo, puedes convertir kilómetros a millas, grados Celsius a Fahrenheit, etc.
func ConversorDeUnidades(cantidad float32, unidadSalida string) float32 {
	if unidadSalida == "Fahrenheit" {
		return (cantidad * 9 / 5) + 32
	}
	if unidadSalida == "Millas" {
		return cantidad / 1.609
	}
	return 0
}

// Contador de palabras: Desarrolla una función que cuente la cantidad de palabras en una cadena de texto.
// Puedes ignorar los caracteres especiales y contar solo las palabras separadas por espacios.
func ContadorDePalabras(str string) int {
	return len(strings.Fields(str))
}

// Generador de contraseñas: Escribe un programa que genere contraseñas aleatorias con diferentes niveles de seguridad.
// Puedes permitir que el usuario especifique la longitud.
func GeneradorDeContrasenas(longitud int) string {
	var sb strings.Builder
	for sb.Len() < longitud {
		num := rand.Intn(55) + 65
		if num <= 96 && num >= 91 {
			sb.Write([]byte{byte(num + 10)})
			continue
		}
		sb.Write([]byte{byte(num)})
	}
	return sb.String()
}

// Cálculo del factorial: Implementa una función que calcule el factorial de un número entero dado.
// Hazlo de manera iterativa.
func FactorialIterativo(num int) int {
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	return factorial
}
