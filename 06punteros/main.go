package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	edad   int8
}

func (p *Persona) asignarEdad(e int8) {
	if p.edad > 0 {
		p.edad = e
		fmt.Println("Mi edad es:", p.edad)
	} else {
		fmt.Println("no se puede validar sobre negativos")
	}
}

func main() {
	var persona Persona
	persona.edad = 20
	persona.asignarEdad(persona.edad)
}
