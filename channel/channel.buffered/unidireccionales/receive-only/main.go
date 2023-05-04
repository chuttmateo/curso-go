package main

import "fmt"

func main() {
	receiveonly := make(<-chan int, 2)
	fmt.Printf("receiveonly: %T\n", receiveonly)
}