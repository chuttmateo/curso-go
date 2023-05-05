package main

import (
	"fmt"
	"sync"
)
func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)
	nums := []int{1,23,45,34,5,12,34,12,3,13,2,23,4,121,23,34,56,42,34,54,67}


	go enviar(par, impar, nums...)

	go recibir(par, impar, fanin)

	for v := range fanin { 
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

func enviar(p, i chan<- int, n ...int)  {
	for _, v := range n {
		if v % 2 == 0{
			p <- v
		}else {
			i <- v
		}
	}
	close(p)
	close(i)
}

func recibir(p, i <-chan int,o chan<- int)  {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range p {
			o <- v
		}
	}()

	go func ()  {
		defer wg.Done()
		for v := range i {
			o <- v
		}
	}()
	wg.Wait()
	close(o)
}
//código de Eduar