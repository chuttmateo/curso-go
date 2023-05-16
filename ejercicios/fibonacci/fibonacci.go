package fibonacci

func SecuenciaFibonacci(x int) []int {
	var z []int
	z = append(z, 0, 1)
	for i := 0; i < (x - 2); i++ {
		z = append(z, z[len(z)-2]+z[len(z)-1])
	}
	return z
}
