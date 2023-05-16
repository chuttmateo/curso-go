package primo

import "fmt"

func EsPrimo(x int) string {
	if x == 1 {
		return fmt.Sprintln(x, "NO es primo")
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return fmt.Sprintln(x, "NO es primo")
		}
	}
	return fmt.Sprintln(x, "es primo")
}
