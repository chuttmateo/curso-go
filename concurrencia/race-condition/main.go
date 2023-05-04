package main

import (
	"fmt"
	"runtime"
	"sync"
)
var contador int

func main() {
	
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(1000)
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		go func() {
			x := contador
			runtime.Gosched()
			x++
			contador = x
			wg.Done()
		}()
		//fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("total ", contador)
	// este código presenta inconsistencias, el valor total no siempre es 100, cuando debería serlo. 
	// acá se presenta un erro llamado race condition
	//algo curioso: es que a mas grande el número de goroutines mas % de error hay. otra cosa es que con el Println de la linea 25 el % de error disminuye; si lo quito aumeta el % de error
}
