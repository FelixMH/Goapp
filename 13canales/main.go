package main

import "fmt"

func main() {
	bodegaOrigen := []string{
		"php",
		"javascript",
		"html",
		"css",
		"java",
		"bd",
		"git",
	}

	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopia := make(chan string)

	// se deja en la bodega de destino
	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	// se carga la fotocopiadora

	go func() {
		for {
			// se entrega el contenido de el canal a variable libro
			libro := <-fotocopiadora

			fotocopia <- libro

			// Agregar el libro a la bodegaDestino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				// Cerrar el canal
				close(fotocopia)
			}
		}
	}()

	fmt.Println("** Listado de libros fotocopiados **")
	for {
		if libro, ok := <-fotocopia; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
}
