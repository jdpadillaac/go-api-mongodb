package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-api-mongodb/services/userservice"
)

// GetUserByID => obtener un solo usuario meduiante un oid en los parametros
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var idUser string = vars["id"]
	if len(idUser) == 0 {
		http.Error(w, "El id enviado no es un iod valido", http.StatusBadRequest)
		return
	}

	userFounded, founded := userservice.FindByID(idUser)
	if founded == false {
		http.Error(w, "No se encontró un usuario registrado con el id especificado", http.StatusBadRequest)
		return
	}

	userResponse, err := json.Marshal(userFounded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userResponse)

}
