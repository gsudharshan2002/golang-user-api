package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	
	clientOptions := options.Client().ApplyURI("mongodb+srv://sudharshan:<addpass>@mymango.hdvc0tg.mongodb.net/userdb?retryWrites=true&w=majority&appName=MyMango&tls=true")


	
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client
	log.Println("✅ Connected to MongoDB")
}


func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("userdb").Collection(collectionName)
}
