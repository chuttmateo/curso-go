package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Soy la función anónima")
	}()

	//Expresión función; ciudadano de primer orden; una función también es un type y por ende la podemos asignar a una variable
	f := func(s string) {
		fmt.Println("Hola,", s)
	}
	f("Mateo")
	//Función que retorna una función
	x := bar()
	fmt.Printf("%T\n", x)
	y := x()
	fmt.Println(y) 
		//fmt.Println(bar()())

}
func foo() {
	fmt.Println("Soy la función foo")
}

func bar() func() int {

	return func() int {
		return 1492
	}

}
