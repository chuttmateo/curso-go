package main

import "fmt"

func main() {
	var c chan int = make(chan int)

	go enviar(c)

	recibir(c)
}

func enviar(x chan<- int) {
	x <- 10
}
func recibir(x <-chan int) {
	fmt.Println(<-x)
}