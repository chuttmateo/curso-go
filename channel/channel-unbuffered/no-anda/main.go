package main

import "fmt"

func main() {
	ca := make(chan int)
	ca <- 100
	fmt.Println(<-ca)
	
}