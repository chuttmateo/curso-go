package main

import "fmt"

func main() {
	ca := make(chan int, 1) //creo un canal buffered de tamaño 1 y lo asigno a una variable de nombre ca

	ca <- 100 // al canal en la variable ca le envio un int 100

	fmt.Printf("ca: %v\n", <-ca)
	// a diferencia del channel unbuffered este anda porque puede almacenar un dato, en el (bufér)
}