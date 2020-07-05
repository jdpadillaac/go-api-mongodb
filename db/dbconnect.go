package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN => Se exporta conexion a mongo para utilizar en toda la app
var MongoCN = conectarDB()

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/go-test-db?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false")

func conectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa a la db")
	return client
}

// TestConnection => verifica la conexion a la db mediante un pong
func TestConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
