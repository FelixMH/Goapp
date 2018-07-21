package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
		}
	}()
	fmt.Println("Llamando g.")
	g(4)
	fmt.Println("Retorna normalmente desde g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Aaaaaaaahhhh!!!")
		panic("El número no puede ser mayor a 3")
	}
	defer fmt.Println("defer en la función g.")
	fmt.Println("Imprimiendo en g", i)
}
