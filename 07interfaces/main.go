package main

import "github.com/FelixMH/goapp/07interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	p.Nombre = "Manchas"
	// adoptarPerro(p)
	// adoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

// funcion que permite recibir los datos de la interface.
func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}

// func adoptarPerro(p animales.Perro) {
// 	p.Comunicarse()
// }
//
// func adoptarGato(g animales.Gato) {
// 	g.Comunicarse()
// }
