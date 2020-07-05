package main

import (
	"log"

	"github.com/jdpadillaac/go-api-mongodb/db"
	"github.com/jdpadillaac/go-api-mongodb/handlers"
)

func main() {

	if db.TestConnection() == false {
		log.Fatal("Sin conexion")
		return
	}

	handlers.Handlers()
}
