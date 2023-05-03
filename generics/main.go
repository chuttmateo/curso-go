package main

import "fmt"

type intOrFloat interface {
	int | float64
}

type dinero float64 // no es buena práctica crear nuevos tipos si no le vamos a agregar métodos o algún uso en particular

func main() {
	divirPorDos(10)
	divirPorDos(10.50)
	hacerPrint("hola")
	d := multiplicarDinero[dinero](5.50)
	fmt.Println(d)
}

func divirPorDos[T intOrFloat](a T) {
	fmt.Println(a/2)
}
// utilizo ~ para indicar que la función pueda aceptar cualquier tipo que tenga como tipo subyasente float64, en este caso "dinero"
func multiplicarDinero[T ~float64](d T) T {
	return d * 5
}

func hacerPrint[T any](a T){
	fmt.Println(a)
}