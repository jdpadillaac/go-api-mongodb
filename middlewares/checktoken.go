package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/jdpadillaac/go-api-mongodb/models"
	"github.com/jdpadillaac/go-api-mongodb/services/userservice"
)

// UserLogged => Usuario que ha hecho en la operación
var UserLogged models.User

// ValidateJWT func
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := validateToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en validacion de autorización"+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)

	}
}

func validateToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("go-mongo-apu-key-to-secret")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	cleanToken := strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(cleanToken, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		user, founded := userservice.FindByEmail(claims.Email)
		if founded == true {
			UserLogged = user
			log.Println(UserLogged)
		}
		return claims, founded, user.ID.Hex(), nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err

}
