package routers

import (
	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-api-mongodb/controllers/authcontroller"
	"github.com/jdpadillaac/go-api-mongodb/middlewares"
)

// AuthRoutes => Rutas de validacion y autenticacion
func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", middlewares.CheckDB(authcontroller.Login)).Methods("POST")
}
