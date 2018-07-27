package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Productos struct {
	gorm.Model
	Nombre       string
	codigoBarras string
	Price        uint
}

func main() {
	db, err := gorm.Open("mysql", "dev:secret@/Goapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")
}
