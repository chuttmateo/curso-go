package main

import (
	"fmt"
	"runtime"
	"sync"
)
var contador int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(1000)
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			x := contador
			x++
			contador = x
			mutex.Unlock()
			wg.Done()
		}()
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("total ", contador)
	// con mutex soluciono la race condition, ahora siempre el valor final será 1000
}
