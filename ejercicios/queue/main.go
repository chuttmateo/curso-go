package main

import (
	"fmt"
)

// Crea una estructura de datos llamada Queue (cola) que tenga los métodos
// Enqueue (para agregar un elemento a la cola),
// Dequeue (para eliminar y devolver el elemento frontal de la cola) y
// IsEmpty (para verificar si la cola está vacía).

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(elemento interface{}) {
	*q = append(*q, elemento)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	elemento := (*q)[0]
	*q = (*q)[1:]
	return elemento
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
func (q *Queue) Size() int {
	return len(*q)
}

func main() {
	cola := NewQueue()
	fmt.Println("La cola esta vacia: ", cola.IsEmpty())
	cola.Enqueue("hola")
	cola.Enqueue("chau")
	fmt.Println("Los elementos que están en la cola: ", *cola)
	fmt.Println("La cola esta vacia: ", cola.IsEmpty())
	fmt.Println("Retorno de cola.Dequeue(): ", cola.Dequeue())
	fmt.Println("Los elementos que están en la cola: ", *cola)
	fmt.Println("Tamaño de la cola: ", cola.Size())
}
