package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Se indica el numero de procesadores a utilizar en las gorutinas
	// con el plugin runtime de go
	runtime.GOMAXPROCS(2)
	// Sala de espera para gorutinas
	var wg sync.WaitGroup

	numbers := []uint32{
		125424,
		235,
		1241877,
		32135647,
		6544774,
		4174858,
	}
	// Agregar cantidad de gorutinas dependiendo de la cantidad de datos que haya en el slice
	wg.Add(len(numbers))

	fmt.Println("COmienza el proceso...")

	for _, v := range numbers {
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a, esPrimo(a))
		}(v)
	}
	wg.Wait()
	fmt.Println("Terminó el proceso.")
}

func esPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
