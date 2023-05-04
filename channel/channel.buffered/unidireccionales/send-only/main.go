package main

import "fmt"

func main() {
	sendonly := make(chan<- int, 1)
	sendonly <- 50
	fmt.Printf("sendonly: %v\n", sendonly)
}