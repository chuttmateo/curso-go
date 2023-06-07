package main

import "fmt"

func main() {
	fmt.Println(sumaDePares(suma, []int{1, 2, 3, 4, 5, 6}...))
}

func suma(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sumaDePares(f func(...int) int, nums ...int) int {
	var y []int
	for _, v := range nums {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
