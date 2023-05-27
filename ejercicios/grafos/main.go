package main

import "fmt"

// Crea una estructura de datos para representar un grafo dirigido.
// La estructura debe tener métodos para agregar vértices, agregar aristas y realizar un recorrido en profundidad (DFS) desde un vértice dado.

type Grafo struct {
	vertices map[string][]string
}

func NewGrafo() *Grafo {
	return &Grafo{
		vertices: make(map[string][]string),
	}
}
func (g *Grafo) addVertice(vertice string) {
	if _, existe := g.vertices[vertice]; !existe {
		g.vertices[vertice] = []string{}
	}
}
func (g *Grafo) addArista(desde, hasta string) {
	if _, existe := g.vertices[hasta]; existe {
		g.vertices[desde] = append(g.vertices[desde], hasta)
	}
}
func (g *Grafo) recorrer(desde string) {
	fmt.Println(desde)
	for _, v := range g.vertices[desde] {
		g.recorrer(v)
	}
}

func main() {
	g := NewGrafo()
	g.addVertice("a")
	g.addVertice("b")
	g.addVertice("B")
	g.addVertice("c")

	g.addArista("a", "b")
	g.addArista("a", "B")
	g.addArista("a", "j")
	g.addArista("b", "c")

	g.recorrer("a")
}
