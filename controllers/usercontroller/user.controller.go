package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jdpadillaac/go-api-mongodb/models"
	"github.com/jdpadillaac/go-api-mongodb/services/userservice"
)

// UserRegister => controlador de la ruta de registro de usuario
func UserRegister(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en la validacion de datos "+err.Error(), http.StatusBadRequest)
		return
	}

	var userFounded bool = userservice.ValidateIfUserExistByEmail(user.Email)

	if userFounded == true {
		http.Error(w, "El usuario con este email ya está registrado en la base de datos", http.StatusBadRequest)
		return
	}

	id, _, err := userservice.SaveUser(user)
	if err != nil {
		http.Error(w, "Error en registro en la base de datos "+err.Error(), 500)
		return
	}

	userFOunded, _ := userservice.FindByID(id)
	userFOunded.Password = ""

	userResponse, err := json.Marshal(userFOunded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userResponse)

}
