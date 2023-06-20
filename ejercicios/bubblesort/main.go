package main

import (
	"fmt"
)

func main() {
	x := []int{2, 3, 4, 14, 8, 11, 10, 13, 7, 9, 5, 6, 15, 12, 1}
	i := bubbleSort(x)
	fmt.Println(x, i)
}

func bubbleSort(nums []int) int {
	size := len(nums)
	fmt.Println(size)
	ieraciones := 0
	for i := 0; i < size; i++ {
		ieraciones++
		for j := 0; j < size-i-1; j++ {
			ieraciones++
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return ieraciones
}

//con el -i reduzco de 225 a 120 las iteracciones para un slice con una logitud de 15 elementos
