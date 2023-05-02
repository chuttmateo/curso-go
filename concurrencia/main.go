package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num Goroutine:", runtime.NumGoroutine())
	fmt.Println("ARCH:",runtime.GOARCH)
	fmt.Println("OS:", runtime.GOOS)

	wg.Add(1)
	fmt.Println("Num Goroutine:", runtime.NumGoroutine())
	go foo() // con "go" indico que se ejecute la función en una goroutine
	fmt.Println("Num Goroutine:", runtime.NumGoroutine())
	bar()

	wg.Wait()
}

func foo() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 700; i++ {
		fmt.Println("bar:", i)
	}
}