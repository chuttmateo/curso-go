package main

import "fmt"

type Persona struct {
	Nombre string
}

type humano interface {
	hablar()
}

func (p *Persona) hablar() {
	fmt.Println("hablando")
}

func main() {
	p1 := Persona{
		Nombre: "Mateo",
	}
	diAlgo(&p1) // &p1, de tipo *Persona, sí implementa la interface `humano`. El método `hablar()` tiene como receptor tipo `*Persona`. `&p1` es del tipo `Humano` ya que implementa implícitamente la interfaz.
	//diAlgo(p1)// p1, de tipo Persona, no implementa la interface `humano`. No hay un método `hablar()` que tenga un receptor de tipo `Persona`, por ende `p1` no es de tipo humano
}
func diAlgo(h humano)  {
	h.hablar()
}