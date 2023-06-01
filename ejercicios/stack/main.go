package main

import "fmt"

// Crea una estructura de datos llamada Stack (pila) que tenga los métodos Push (para agregar un elemento a la pila),
// Pop (para eliminar y devolver el elemento superior de la pila)
// y IsEmpty (para verificar si la pila está vacía).

type stack []interface{}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(element interface{}) {
	*s = append(*s, element)
}
func (s *stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}

	elemento := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elemento
}
func (s *stack) IsEmpty() bool {
	// retorno := true
	// if len(*s) > 0 {
	// 	retorno = false
	// }
	// return retorno
	return len(*s) == 0
}

func main() {
	s := NewStack()
	fmt.Println("La pila está vacia: ", s.IsEmpty())
	fmt.Println("Intento hacer pop a una pila vacia: ", s.Pop())
	fmt.Println("Hago push de un elemento a la pila")
	s.Push("hola")
	fmt.Println("La pila está vacia: ", s.IsEmpty())
	fmt.Println("Intento hacer pop a una pila con un elemento: ", s.Pop())
}
