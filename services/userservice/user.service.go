package userservice

import (
	"context"
	"time"

	"github.com/jdpadillaac/go-api-mongodb/db"
	"github.com/jdpadillaac/go-api-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// ValidateIfUserExistByEmail => buscar unsuario en la base de datos por email
func ValidateIfUserExistByEmail(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbs := db.MongoCN.Database("go-test-db")
	col := dbs.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		return false
	}
	return true
}

// SaveUser => guarda un nuevo usuario en la base de datos
func SaveUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbs := db.MongoCN.Database("go-test-db")
	col := dbs.Collection("users")

	u.Password, _ = encriptPassword(u.Password)
	u.CreatedAt = time.Now()

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID := result.InsertedID.(primitive.ObjectID).Hex()
	return ObjID, true, nil
}

// FindByID => Funcion para obtener usuario por id
func FindByID(id string) models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbs := db.MongoCN.Database("go-test-db")
	col := dbs.Collection("users")

	userID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": userID}
	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result
	}
	return result

}

func encriptPassword(password string) (string, error) {
	var cost int = 5
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err

}
