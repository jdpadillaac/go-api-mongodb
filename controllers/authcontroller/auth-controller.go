package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/jdpadillaac/go-api-mongodb/models"
	"github.com/jdpadillaac/go-api-mongodb/services/authservices"
)

// Login => Ruta de autenticacion
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos", http.StatusBadRequest)
		return
	}

	userLogged, exist := authservices.Login(user.Email, user.Password)
	if exist == false {
		http.Error(w, "Usuario y/o contraseña invalidos", http.StatusUnauthorized)
		return
	}

	userLogged.Password = ""

	jwtKey, err := generarJwt(userLogged)
	if err != nil {
		http.Error(w, "Error en la generaciond e token de autenticación", http.StatusInternalServerError)
		return
	}

	resp := models.LoginReponse{
		Token: jwtKey,
		User:  userLogged,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)

}

func generarJwt(user models.User) (string, error) {

	myKey := []byte("go-mongo-apu-key-to-secret")
	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
