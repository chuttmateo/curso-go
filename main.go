package main

import (
	"fmt"
)

func main() {
	c := []int{1, 2, 3, 4, 6, 7}
	foo(c...)
}
func foo(x ...int)  {
	for v, i := range x{
		fmt.Println(v, i)
	}
}