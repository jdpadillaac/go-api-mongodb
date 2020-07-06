package routers

import (
	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-api-mongodb/controllers/usercontroller"
	"github.com/jdpadillaac/go-api-mongodb/middlewares"
)

// UserRoutes => Rutas de validacion y autenticacion
func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetUserByID))).Methods("GET")
	r.HandleFunc("/users", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetAllUsers))).Methods("GET")
}
