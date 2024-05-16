package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() (client *mongo.Client, ctx context.Context) {

	mongoURL := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		log.Fatal(err) // log.Fatal will exit the program
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	return
}
