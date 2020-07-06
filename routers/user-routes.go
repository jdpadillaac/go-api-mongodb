package routers

import (
	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-api-mongodb/controllers/authcontroller"
	"github.com/jdpadillaac/go-api-mongodb/middlewares"
)

// UserRoutes => Rutas de validacion y autenticacion
func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", middlewares.CheckDB(middlewares.ValidateJWT(authcontroller.Login))).Methods("GET")
}
