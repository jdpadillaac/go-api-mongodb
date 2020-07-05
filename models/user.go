package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User => Estructura usuario
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Surenames string             `bson:"surenames" json:"surenames,omitempty"`
	Email     string             `bson:"email" json:"email"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Image     string             `bson:"image" json:"image,omitempty"`
}
