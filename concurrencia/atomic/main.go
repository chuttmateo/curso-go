package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&contador,1)
			fmt.Printf("atomic.LoadInt64(&contador): %v\n", atomic.LoadInt64(&contador))
		}()
	}
	wg.Wait()
	fmt.Println("total ", contador)
	//esta es otra manera de evitar el race condition
}
