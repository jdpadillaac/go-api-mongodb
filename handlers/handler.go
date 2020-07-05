package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-api-mongodb/controllers/usercontroller"
	"github.com/jdpadillaac/go-api-mongodb/middlewares"
	"github.com/jdpadillaac/go-api-mongodb/routers"
	"github.com/rs/cors"
)

// Handlers => En la funcion que maneja las peticiones
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/new", middlewares.CheckDB(usercontroller.UserRegister)).Methods("POST")
	router.HandleFunc("/users", middlewares.CheckDB(usercontroller.UserRegister)).Methods("GET")

	routers.AuthRoutes(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4800"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
