package numeros

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Desarrolla un programa que tome una lista de palabras y las ordene alfabéticamente.
func Ordenar(x []string) []string {
	sort.Strings(x)
	return x
}

// Escribe una función que tome un arreglo de enteros como parámetro
// y devuelva la suma de los números pares y la suma de los números impares por separado.
func SumaParesEImpares(x []int) (int, int) {
	var (
		pares   int
		impares int
	)
	for _, v := range x {
		if v%2 == 0 {
			pares = pares + v
		} else {
			impares = impares + v
		}
	}
	return pares, impares
}

// Implementa una función que tome un número entero positivo como parámetro
// y devuelva true si es un número perfecto y false en caso contrario.
// Un número perfecto es aquel cuya suma de sus divisores propios
// (excluyendo al propio número) es igual al número.
func EsPerfecto(x uint) bool {
	z := int(x)
	var divisoresPropios int
	for i := 1; i < z; i++ {
		if z%i == 0 {
			divisoresPropios += i
		}
	}

	return divisoresPropios == z
}

// Crea una función que tome dos slices de enteros como parámetros
// y devuelva un nuevo slice que contenga los elementos comunes entre ambos slices.
func ElementosComunes(x, z []int) []int {
	var (
		menor       []int = x
		mayor       []int = z
		comunes     []int
		iteraciones int
	)

	if len(x) > len(z) {
		menor = z
		mayor = x
	}
	//al poner el slice menor en el ciclo exterior las iteraciones disminuyen
	for _, vx := range menor {
		iteraciones++
		for _, vz := range mayor {
			iteraciones++
			if vx == vz {
				comunes = append(comunes, vz)
			}
		}
	}
	fmt.Println("Iteraciones en la func ElementosComunes", iteraciones)
	return comunes
}

// ahora usando map
func ElementosComunesMap(x, z []int) []int {
	var (
		menor       []int = x
		mayor       []int = z
		comunes     []int
		iteraciones int
	)

	if len(x) > len(z) {
		menor = z
		mayor = x
	}

	elementos := make(map[int]bool)
	for _, v := range menor {
		iteraciones++
		elementos[v] = true
	}

	for _, v := range mayor {
		iteraciones++
		if elementos[v] {
			comunes = append(comunes, v)
		}
	}
	fmt.Println("Iteraciones en la func ElementosComunesMap ", iteraciones)
	return comunes
}

// Escribe una función que tome una matriz cuadrada de números enteros como parámetro
// y devuelva true si es una matriz simétrica y false en caso contrario.
// Una matriz es simétrica si es igual a su traspuesta.
func EsSimetrica(m [][]int) bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[j][i] != m[i][j] {
				return false
			}
		}
	}
	return true
}

// Implementa una función que tome una cadena como parámetro
// y devuelva la frecuencia de cada carácter en la cadena.
func PalabrasFrecuentes(s string) map[string]int {
	sliceString := strings.Fields(s) // strings.Fields(s) es mejor que strings.Split(s, " ")

	var retorno map[string]int = make(map[string]int)

	for _, v := range sliceString {
		retorno[v]++
	}
	return retorno
}

// Desarrolla un programa que lea un archivo de texto y encuentre la palabra más frecuente en el texto.
// Si hay varias palabras con la misma frecuencia máxima, el programa debe imprimir todas ellas.
func LeerArchivo(s string) string {
	osFile, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer osFile.Close()

	scanner := bufio.NewScanner(osFile)
	var sb strings.Builder

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	p := PalabrasFrecuentes(sb.String())
	frecuenciaMaxima := 0
	var palabrasFrecuentes []string = []string{}
	for k, v := range p {
		if v == frecuenciaMaxima {
			palabrasFrecuentes = append(palabrasFrecuentes, k)
		}
		if v > frecuenciaMaxima {
			frecuenciaMaxima = v
			palabrasFrecuentes = []string{k}
		}
	}
	if len(palabrasFrecuentes) > 1 {
		return fmt.Sprintf("Estas son las palabras más frecuentes: %v, apareciendo %v cada una", palabrasFrecuentes, frecuenciaMaxima)
	}
	return fmt.Sprintf("La palabra '%v' es la más frecuente, apareciendo %v veces", palabrasFrecuentes[0], frecuenciaMaxima)

}

// Escribe una función que tome una lista de palabras
// y devuelva un mapa que muestre la longitud de cada palabra.
func LongitudDePalabras(s ...string) map[string]int {
	retorno := make(map[string]int, len(s)) //lo inicializo con una longitud igual a la longitud del arreglo entrante por parametro para mejorar el rendimiento
	for _, v := range s {
		retorno[v] = len(v)
	}
	return retorno
}

// Implementa una función que tome una matriz bidimensional de enteros como parámetro
// y devuelva la suma de todos los elementos.
func SumaDeMatriz(matriz [][]int) int {
	if len(matriz) == 0 || len(matriz[0]) == 0 {
		return 0
	}
	var suma int
	for _, v := range matriz {
		for _, num := range v {
			suma += num
		}
	}
	return suma
}

// Implementa una función que tome una cadena
// y devuelva true si la cadena contiene solo caracteres alfabéticos (letras) y false en caso contrario.
func SonSoloLetras(str string) bool {
	for _, v := range str {
		if v > 64 && v < 91 || v > 96 && v < 123 || v == 32 {
			continue
		}
		return false
	}
	return true
}

//Implementa una función llamada EliminarDuplicados que tome una lista de enteros 
//y elimine los elementos duplicados, devolviendo una nueva lista sin duplicados.
func EliminarDuplicados(x []int) []int {
	retorno := []int{}
	m := make(map[int]bool)
	for _, v := range x {
		if m[v] {
			continue
		}
		m[v]=true
		retorno = append(retorno, v)
	}
	return retorno
}

//Crea una función llamada ReemplazarPalabras que tome una cadena y dos palabras como parámetros, 
//y devuelva una nueva cadena donde todas las apariciones de la primera palabra 
//sean reemplazadas por la segunda palabra.
func ReemplazarPalabras(cad string, p string, r string) string {
	sliceString := strings.Fields(cad)
	var sb strings.Builder
	for _, v := range sliceString {
		if p == v {
			sb.WriteString(r+" ")
		}else{
			sb.WriteString(v+" ")
		}
	}	
	return sb.String()
}