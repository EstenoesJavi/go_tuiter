package main

import (
	"log"

	"github.com/EstenoesJavi/go_tuiter/bd"
	"github.com/EstenoesJavi/go_tuiter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
