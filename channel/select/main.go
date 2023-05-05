package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				par <- i
			} else {
				impar <- i
			}
		}
		close(salir)
	}()

	for {
		select {
		case v, ok := <-par:
			fmt.Println("Desde par", v, ok)
		case v, ok := <-impar:
			fmt.Println("Desde impar", v, ok)
		case v, ok := <-salir:
			if !ok {
				fmt.Println("Desde salir", v, ok)
				close(par)
				close(impar)
				return
			}
			fmt.Println(v, ok)
		}
	}

}
