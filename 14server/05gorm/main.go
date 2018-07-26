package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Productos struct {
	gorm.Model
	Nombre       string
	codigoBarras string
	Price        uint
}

func main() {
	db, err := gorm.Open("mysql", "developer:secret@/Goapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")
}
