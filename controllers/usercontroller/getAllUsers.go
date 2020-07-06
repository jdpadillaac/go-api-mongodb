package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jdpadillaac/go-api-mongodb/services/userservice"
)

// GetAllUsers => function to get all enable users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userList, founded := userservice.FindAll()
	if founded == false {
		http.Error(w, "Error en la validacion de datos ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(userList)

}
