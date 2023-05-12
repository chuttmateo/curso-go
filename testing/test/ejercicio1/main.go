package main

func main() {

}
func suma(x ...int) int {
	acc := 0
	for _, v := range x {
		acc += v
	}
	return acc
}
