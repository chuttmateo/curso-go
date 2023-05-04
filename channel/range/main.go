package main

import "fmt"

func main() {
	var c chan int = make(chan int)

	go func ()  {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)// hay que cerrar el canal porque, de no ser así, range intentará recibir datos del canal (una vez mas), eso provocaría un deadlock
	}()

	for v := range c {
		fmt.Printf("v: %v\n", v)
	}
}
