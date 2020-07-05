package authservices

import (
	"github.com/jdpadillaac/go-api-mongodb/models"
	"github.com/jdpadillaac/go-api-mongodb/services/userservice"
	"golang.org/x/crypto/bcrypt"
)

// Login => Proceso de validaciond e usuario
func Login(email string, pass string) (models.User, bool) {

	var user models.User

	var founded bool = userservice.ValidateIfUserExistByEmail(email)
	if founded == false {
		return user, false
	}

	userRegistered, exist := userservice.FindByEmail(email)
	if exist == false {
		return user, false
	}

	passwordBytes := []byte(pass)
	passwordDB := []byte(userRegistered.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return userRegistered, true

}
