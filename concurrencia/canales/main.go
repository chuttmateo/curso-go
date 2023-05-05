package main

import (
	"fmt"
	"sync"
)

func main() {
	ca := make(chan int)
	var wg sync.WaitGroup

	
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				for i := 0; i < 10; i++ {
					ca <- i
				}
				wg.Done()
			}()
		}
		fmt.Println("Última linea de la primer goroutine")
	}()

	go func() {
		fmt.Println("Antes del wg.Wait()")
		wg.Wait()
		fmt.Println("Pasé del wg.Wait()")
		close(ca) // cierro el canal una vez que todas las gorutinas hayan terminado
	}()

	for v := range ca {
		fmt.Println(v)
	}
	
}
