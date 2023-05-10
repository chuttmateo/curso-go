package main

import "fmt"

type errorPersonalizado struct {
	str string
}


func (e errorPersonalizado) Error() string {
	return e.str
}

func main() {
	er := errorPersonalizado{"huston tenemos un problema"}
	foo(er)
	
}
func foo(e error) {
	fmt.Println(e.(errorPersonalizado).str)//hay que hacer una assertion para poder acceder a los atributos internos de "errorPersonalizado"
}