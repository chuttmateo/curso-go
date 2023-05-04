package main

import "fmt"

func main() {
	ca := make(chan int)
	go func() {
		ca <- 100
	}()
	fmt.Printf("ca: %v\n", <-ca)
	//acá anda porque hay una goroutine que recibe el valor que envia otra goroutine; la que recibe es la main
}