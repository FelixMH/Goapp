package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()

	fmt.Println("Consultando información..., set de datos")
	defer cerrarSetDatos()
}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos...")
}

func cerrarSetDatos() {
	fmt.Println("cerrar el set de datos...")
}
