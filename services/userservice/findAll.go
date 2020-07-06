package userservice

import (
	"context"
	"time"

	"github.com/jdpadillaac/go-api-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindAll => return all users registered
func FindAll() ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.User

	condition := bson.M{}
	cursor, err := UserColllection.Find(ctx, condition)
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return results, false
		}
		user.Password = ""
		results = append(results, &user)
	}

	return results, true
}
